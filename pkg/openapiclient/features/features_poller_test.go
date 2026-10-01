package features

import (
	"context"
	"math"
	"net/http"
	"sync"
	"sync/atomic"
	"testing"
	"testing/synctest"
	"time"

	"github.com/StackVista/stackstate-receiver-go-client/generated/receiver_api"
)

type pollTestAPI struct {
	ctx     context.Context
	calls   atomic.Int64
	execute func(context.Context, int) (map[string]any, *http.Response, error)
}

func (a *pollTestAPI) GetFeatures(ctx context.Context) receiver_api.ApiGetFeaturesRequest {
	a.ctx = ctx
	return receiver_api.ApiGetFeaturesRequest{ApiService: a}
}

func (a *pollTestAPI) GetFeaturesExecute(receiver_api.ApiGetFeaturesRequest) (map[string]any, *http.Response, error) {
	call := a.calls.Add(1)
	return a.execute(a.ctx, int(call))
}

func pollTestClient(t *testing.T, api *pollTestAPI) *Client {
	t.Helper()
	client, err := NewClient(api, QueryOptions{
		Timeout: 20 * time.Second, AttemptTimeout: 5 * time.Second, MaxAttempts: 3,
		InitialBackoff: time.Second, MaxBackoff: 2 * time.Second,
	})
	if err != nil {
		t.Fatal(err)
	}
	client.random = func() float64 { return 0.5 }
	return client
}

func startTestPoller(ctx context.Context, t *testing.T, client *Client, opts PollOptions) *Poller {
	t.Helper()
	poller, err := client.StartPolling(ctx, opts)
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(poller.Stop)
	return poller
}

func assertPollerClosed(t *testing.T, poller *Poller) {
	t.Helper()
	<-poller.Done()
	if _, ok := <-poller.Results(); ok {
		t.Fatal("Results was not closed before Done")
	}
}

func TestPollingOptions(t *testing.T) {
	client := pollTestClient(t, &pollTestAPI{})
	for _, opts := range []PollOptions{
		{}, {Interval: -time.Second}, {Interval: time.Second, Jitter: -0.1},
		{Interval: time.Second, Jitter: 1}, {Interval: time.Second, Jitter: math.NaN()},
		{Interval: time.Second, Jitter: math.Inf(1)}, {Interval: time.Second, Jitter: math.Inf(-1)},
		{Interval: time.Duration(math.MaxInt64), Jitter: 0.2},
	} {
		if poller, err := client.StartPolling(context.Background(), opts); err == nil {
			poller.Stop()
			t.Errorf("accepted invalid options: %+v", opts)
		}
	}
	if _, err := client.StartPolling(nil, PollOptions{Interval: time.Second}); err == nil {
		t.Error("accepted nil context")
	}
	opts := PollOptions{Interval: time.Duration(math.MaxInt64)}
	poller := startTestPoller(context.Background(), t, client, opts)
	poller.Stop()
	assertPollerClosed(t, poller)
	if got := pollDelay(opts, 0.5); got != opts.Interval {
		t.Errorf("zero jitter changed maximum interval: %v", got)
	}
}

func TestPollingDelayAndJitter(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		api := &pollTestAPI{execute: func(context.Context, int) (map[string]any, *http.Response, error) {
			return map[string]any{}, &http.Response{StatusCode: 200}, nil
		}}
		client := pollTestClient(t, api)
		draws := []float64{0, 0.5, 0.999}
		draw := 0
		client.random = func() float64 {
			value := draws[draw%len(draws)]
			draw++
			return value
		}
		opts := PollOptions{Interval: 10 * time.Second, Jitter: 0.2}
		poller := startTestPoller(context.Background(), t, client, opts)
		for i, random := range draws {
			synctest.Wait()
			delay := pollDelay(opts, random)
			time.Sleep(delay - time.Nanosecond)
			synctest.Wait()
			if api.calls.Load() != int64(i) {
				t.Fatalf("query before jittered interval: calls=%d, want %d", api.calls.Load(), i)
			}
			time.Sleep(time.Nanosecond)
			result := <-poller.Results()
			if result.Class != Valid || result.Attempts != 1 || !result.FinishedAt.Equal(time.Now()) {
				t.Fatalf("unexpected observation: %+v", result)
			}
		}
		poller.Stop()
		assertPollerClosed(t, poller)
	})
	for _, test := range []struct {
		random float64
		want   time.Duration
	}{{0, 8 * time.Second}, {0.5, 10 * time.Second}, {1, 12 * time.Second}} {
		if got := pollDelay(PollOptions{Interval: 10 * time.Second, Jitter: 0.2}, test.random); got != test.want {
			t.Errorf("delay = %v, want %v", got, test.want)
		}
	}
	if got := pollDelay(PollOptions{Interval: time.Nanosecond, Jitter: 0.9}, 0); got != time.Nanosecond {
		t.Errorf("minimum delay = %v", got)
	}
}

func TestPollingPreservesOutcomesAndBackpressure(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		statuses := []int{200, 404, 401, 503, 200, 400, 200}
		classes := []Class{Valid, Unsupported, Authentication, Transient, Malformed, Rejected, Valid}
		type authKey struct{}
		api := &pollTestAPI{execute: func(ctx context.Context, call int) (map[string]any, *http.Response, error) {
			if ctx.Value(authKey{}) != "synthetic-auth-context" {
				t.Error("query lost authenticated context")
			}
			if call > len(statuses) {
				t.Error("unexpected overlapping or extra query")
				return nil, nil, nil
			}
			values := map[string]any{"capacity": float64(42)}
			if call == 5 {
				values = nil
			}
			return values, &http.Response{StatusCode: statuses[call-1]}, nil
		}}
		client := pollTestClient(t, api)
		client.opts.MaxAttempts = 1
		poller := startTestPoller(context.WithValue(context.Background(), authKey{}, "synthetic-auth-context"), t, client,
			PollOptions{Interval: time.Second})
		synctest.Wait()
		for i, class := range classes {
			time.Sleep(time.Second)
			synctest.Wait()
			finishedAt := time.Now()
			time.Sleep(time.Minute)
			synctest.Wait()
			if api.calls.Load() != int64(i+1) {
				t.Fatalf("queried while output blocked: calls=%d", api.calls.Load())
			}
			result := <-poller.Results()
			if result.Class != class || result.StatusCode != statuses[i] || !result.FinishedAt.Equal(finishedAt) {
				t.Fatalf("lost or changed outcome: %+v", result)
			}
			if class == Valid && result.Features["capacity"] != float64(42) {
				t.Error("lost unrelated feature value")
			}
			synctest.Wait()
		}
		poller.Stop()
		assertPollerClosed(t, poller)
	})
}

func TestPollingRetriesAreOneObservation(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		api := &pollTestAPI{execute: func(_ context.Context, call int) (map[string]any, *http.Response, error) {
			if call < 3 {
				return nil, &http.Response{StatusCode: 503}, nil
			}
			return map[string]any{}, &http.Response{StatusCode: 200}, nil
		}}
		poller := startTestPoller(context.Background(), t, pollTestClient(t, api), PollOptions{Interval: time.Second})
		result := <-poller.Results()
		if result.Class != Valid || result.Attempts != 3 || api.calls.Load() != 3 {
			t.Fatalf("attempts emitted as separate observations: %+v", result)
		}
		poller.Stop()
		assertPollerClosed(t, poller)
	})
}

func TestPollingCancellation(t *testing.T) {
	for _, phase := range []string{"interval", "http", "backoff", "output"} {
		for _, parentCancel := range []bool{false, true} {
			name := phase + "/stop"
			if parentCancel {
				name = phase + "/parent"
			}
			t.Run(name, func(t *testing.T) {
				synctest.Test(t, func(t *testing.T) {
					requestCanceled := false
					api := &pollTestAPI{execute: func(ctx context.Context, _ int) (map[string]any, *http.Response, error) {
						switch phase {
						case "http":
							<-ctx.Done()
							requestCanceled = true
							return nil, nil, ctx.Err()
						case "backoff":
							return nil, &http.Response{StatusCode: 503}, nil
						default:
							return map[string]any{}, &http.Response{StatusCode: 200}, nil
						}
					}}
					ctx, cancel := context.WithCancel(context.Background())
					defer cancel()
					poller := startTestPoller(ctx, t, pollTestClient(t, api), PollOptions{Interval: time.Second})
					synctest.Wait()
					if phase != "interval" {
						time.Sleep(time.Second)
						synctest.Wait()
					}
					if phase == "http" {
						time.Sleep(2 * time.Second)
						synctest.Wait()
					}
					before := time.Now()
					if parentCancel {
						cancel()
					} else {
						var stops sync.WaitGroup
						for range 10 {
							stops.Go(poller.Stop)
						}
						stops.Wait()
					}
					assertPollerClosed(t, poller)
					poller.Stop()
					if !time.Now().Equal(before) {
						t.Error("shutdown waited for a timer")
					}
					if phase == "interval" && api.calls.Load() != 0 || phase != "interval" && api.calls.Load() != 1 {
						t.Errorf("unexpected calls during cancellation: %d", api.calls.Load())
					}
					if phase == "http" && !requestCanceled {
						t.Error("request did not observe cancellation")
					}
				})
			})
		}
	}
}

func TestPollingStopDoesNotWaitForRequestReturn(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		release := make(chan struct{})
		api := &pollTestAPI{execute: func(ctx context.Context, _ int) (map[string]any, *http.Response, error) {
			<-ctx.Done()
			<-release
			return nil, nil, ctx.Err()
		}}
		poller := startTestPoller(context.Background(), t, pollTestClient(t, api), PollOptions{Interval: time.Second})
		synctest.Wait()
		time.Sleep(time.Second)
		synctest.Wait()
		poller.Stop()
		synctest.Wait()
		select {
		case <-poller.Done():
			t.Fatal("Done closed before request returned")
		default:
		}
		close(release)
		assertPollerClosed(t, poller)
	})
}
