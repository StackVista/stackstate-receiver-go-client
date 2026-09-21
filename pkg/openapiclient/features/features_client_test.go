package features_test

import (
	"context"
	"crypto/x509"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/http/httptest"
	"net/url"
	"reflect"
	"strings"
	"sync/atomic"
	"testing"
	"testing/synctest"
	"time"

	"github.com/StackVista/stackstate-receiver-go-client/generated/receiver_api"
	"github.com/StackVista/stackstate-receiver-go-client/pkg/openapiclient"
	"github.com/StackVista/stackstate-receiver-go-client/pkg/openapiclient/features"
)

const (
	queryTestAPIKey = "synthetic-query-api-key"
	queryTestBody   = "synthetic-private-response-body"
	queryTestError  = "synthetic-private-transport-error"
)

func queryTestOptions() features.QueryOptions {
	return features.QueryOptions{
		Timeout:             10 * time.Second,
		AttemptTimeout:      time.Second,
		MaxAttempts:         3,
		InitialBackoff:      time.Nanosecond,
		MaxBackoff:          time.Nanosecond,
		BooleanCapabilities: []string{"otel-logs"},
	}
}

func newQueryTestClient(t *testing.T, api receiver_api.FeaturesAPI, opts features.QueryOptions) *features.Client {
	t.Helper()
	client, err := features.NewClient(api, opts)
	if err != nil {
		t.Fatalf("NewClient: %v", err)
	}
	return client
}

func newHTTPQueryTestClient(t *testing.T, handler http.HandlerFunc, opts features.QueryOptions) (*features.Client, context.Context) {
	t.Helper()
	server := httptest.NewServer(handler)
	t.Cleanup(server.Close)
	api, authCtx, err := openapiclient.NewOpenAPIClient(context.Background(), openapiclient.ConnectionOptions{
		ReceiverURL:    server.URL + "/deployment/stsAgent/",
		APIKey:         queryTestAPIKey,
		RequestTimeout: 5 * time.Second,
	})
	if err != nil {
		t.Fatalf("NewOpenAPIClient: %v", err)
	}
	return newQueryTestClient(t, api.FeaturesAPI, opts), authCtx
}

func checkQueryResult(t *testing.T, result features.Result, class features.Class, status, attempts int, started time.Time) {
	t.Helper()
	if result.Class != class || result.StatusCode != status || result.Attempts != attempts {
		t.Errorf("result class/status/attempts = %v/%d/%d; want %v/%d/%d",
			result.Class, result.StatusCode, result.Attempts, class, status, attempts)
	}
	if result.FinishedAt.IsZero() || result.FinishedAt.Before(started) || result.FinishedAt.After(time.Now()) {
		t.Error("FinishedAt must be the query completion time")
	}
	if class != features.Valid && result.Features != nil {
		t.Error("failed query returned feature data")
	}
	encoded, err := json.Marshal(result)
	if err != nil {
		t.Fatalf("marshal Result: %v", err)
	}
	for _, rendering := range []string{string(encoded), fmt.Sprint(result), fmt.Sprintf("%+v", result), fmt.Sprintf("%#v", result)} {
		for _, secret := range []string{queryTestAPIKey, queryTestBody, queryTestError} {
			if strings.Contains(rendering, secret) {
				t.Error("Result exposed synthetic sensitive data")
			}
		}
	}
}

func TestFetchFeaturesHTTPStatusAndShape(t *testing.T) {
	tests := []struct {
		name        string
		status      int
		contentType string
		body        string
		class       features.Class
		want        map[string]any
	}{
		{"empty_object", 200, "application/json", `{}`, features.Valid, map[string]any{}},
		{"enabled", 200, "application/json", `{"otel-logs":true}`, features.Valid, map[string]any{"otel-logs": true}},
		{"disabled", 200, "application/json", `{"otel-logs":false}`, features.Valid, map[string]any{"otel-logs": false}},
		{"legacy_numeric_and_unknown_values", 200, "application/json", `{"rbac":true,"capacity":42,"label":"capability","nested":{"a":[1,null]}}`, features.Valid,
			map[string]any{"rbac": true, "capacity": float64(42), "label": "capability", "nested": map[string]any{"a": []any{float64(1), nil}}}},
		{"json_charset", 200, "application/json; charset=utf-8", `{"otel-logs":true}`, features.Valid, map[string]any{"otel-logs": true}},
		{"empty_body", 200, "application/json", "", features.Malformed, nil},
		{"whitespace", 200, "application/json", " \n\t", features.Malformed, nil},
		{"null", 200, "application/json", `null`, features.Malformed, nil},
		{"array", 200, "application/json", `[]`, features.Malformed, nil},
		{"string", 200, "application/json", `"value"`, features.Malformed, nil},
		{"number", 200, "application/json", `42`, features.Malformed, nil},
		{"boolean", 200, "application/json", `true`, features.Malformed, nil},
		{"truncated", 200, "application/json", `{"otel-logs":`, features.Malformed, nil},
		{"trailing_json", 200, "application/json", `{} {}`, features.Malformed, nil},
		{"trailing_garbage", 200, "application/json", `{} garbage`, features.Malformed, nil},
		{"flag_string", 200, "application/json", `{"otel-logs":"true"}`, features.Malformed, nil},
		{"flag_number", 200, "application/json", `{"otel-logs":1}`, features.Malformed, nil},
		{"flag_null", 200, "application/json", `{"otel-logs":null}`, features.Malformed, nil},
		{"flag_array", 200, "application/json", `{"otel-logs":[]}`, features.Malformed, nil},
		{"flag_object", 200, "application/json", `{"otel-logs":{}}`, features.Malformed, nil},
		{"html_success", 200, "text/html", "<html>" + queryTestBody + "</html>", features.Malformed, nil},
		{"sensitive_decode_error", 200, "application/json", `{"` + queryTestAPIKey + `":` + queryTestBody, features.Malformed, nil},
	}
	for _, status := range []int{401, 403, 404, 408, 429, 500, 502, 503, 504, 599, 400, 405, 409, 413, 418, 301, 302, 303, 304, 307, 308, 201, 202, 204, 206} {
		class := features.Rejected
		switch {
		case status == 401 || status == 403:
			class = features.Authentication
		case status == 404:
			class = features.Unsupported
		case status == 408 || status == 429 || status >= 500:
			class = features.Transient
		}
		tests = append(tests, struct {
			name        string
			status      int
			contentType string
			body        string
			class       features.Class
			want        map[string]any
		}{fmt.Sprintf("status_%d_html", status), status, "text/html", "<html>" + queryTestBody + queryTestAPIKey + "</html>", class, nil})
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var calls atomic.Int32
			client, ctx := newHTTPQueryTestClient(t, func(w http.ResponseWriter, r *http.Request) {
				calls.Add(1)
				if r.Method != http.MethodGet || r.URL.Path != "/deployment/stsAgent/features" || r.URL.RawQuery != "" {
					t.Error("unexpected feature request method or URL")
				}
				if r.Header.Get("Authorization") != "ApiKey "+queryTestAPIKey {
					t.Error("attempt lost authenticated context")
				}
				w.Header().Set("Content-Type", tt.contentType)
				w.Header().Set("Location", "/must-not-follow")
				w.WriteHeader(tt.status)
				_, _ = io.WriteString(w, tt.body)
			}, queryTestOptions())
			started := time.Now()
			result := client.FetchFeatures(ctx)
			attempts := 1
			if tt.class == features.Transient {
				attempts = 3
			}
			checkQueryResult(t, result, tt.class, tt.status, attempts, started)
			if int(calls.Load()) != attempts {
				t.Errorf("HTTP calls = %d; want %d", calls.Load(), attempts)
			}
			if !reflect.DeepEqual(result.Features, tt.want) {
				t.Error("decoded feature map differs from expected values")
			}
		})
	}
}

func TestFetchFeaturesHTTPResponseLimit(t *testing.T) {
	const limit = 1 << 20
	for _, chunked := range []bool{false, true} {
		for _, size := range []int{limit - 1, limit, limit + 1} {
			t.Run(fmt.Sprintf("chunked_%t_bytes_%d", chunked, size), func(t *testing.T) {
				body := `{"padding":"` + strings.Repeat("x", size-len(`{"padding":""}`)) + `"}`
				opts := queryTestOptions()
				opts.MaxAttempts = 1
				client, ctx := newHTTPQueryTestClient(t, func(w http.ResponseWriter, _ *http.Request) {
					w.Header().Set("Content-Type", "application/json")
					if chunked {
						w.(http.Flusher).Flush()
					} else {
						w.Header().Set("Content-Length", fmt.Sprint(len(body)))
					}
					_, _ = io.WriteString(w, body)
				}, opts)
				started := time.Now()
				result := client.FetchFeatures(ctx)
				class := features.Valid
				if size > limit {
					class = features.Malformed
				}
				checkQueryResult(t, result, class, http.StatusOK, 1, started)
				if class == features.Valid && result.Features["padding"] != strings.Repeat("x", size-len(`{"padding":""}`)) {
					t.Error("in-bound response was truncated")
				}
			})
		}
	}
	for _, status := range []int{401, 404, 503} {
		t.Run(fmt.Sprintf("oversized_status_%d", status), func(t *testing.T) {
			opts := queryTestOptions()
			opts.MaxAttempts = 1
			client, ctx := newHTTPQueryTestClient(t, func(w http.ResponseWriter, _ *http.Request) {
				w.Header().Set("Content-Type", "text/html")
				w.WriteHeader(status)
				_, _ = io.WriteString(w, strings.Repeat(queryTestBody, limit/len(queryTestBody)+1))
			}, opts)
			class := map[int]features.Class{401: features.Authentication, 404: features.Unsupported, 503: features.Transient}[status]
			started := time.Now()
			checkQueryResult(t, client.FetchFeatures(ctx), class, status, 1, started)
		})
	}
}

func TestFetchFeaturesConsumerCapabilities(t *testing.T) {
	for _, consumer := range []struct {
		name string
		keys []string
	}{
		{"logs", []string{"otel-logs"}},
		{"rbac", []string{"k8s-rbac"}},
		{"both", []string{"otel-logs", "k8s-rbac"}},
		{"object_only", nil},
	} {
		for _, tt := range []struct {
			name string
			body string
		}{
			{"empty", `{}`},
			{"absent", `{"capacity":42}`},
			{"enabled", `{"otel-logs":true,"k8s-rbac":true,"capacity":42}`},
			{"disabled", `{"otel-logs":false,"k8s-rbac":false}`},
			{"logs_string", `{"otel-logs":"true","k8s-rbac":true}`},
			{"rbac_string", `{"otel-logs":true,"k8s-rbac":"true"}`},
			{"logs_null", `{"otel-logs":null,"k8s-rbac":true}`},
			{"rbac_null", `{"otel-logs":true,"k8s-rbac":null}`},
			{"null", `null`},
			{"scalar", `true`},
			{"array", `[]`},
			{"invalid", `{`},
			{"oversized", `{"padding":"` + strings.Repeat("x", 1<<20) + `"}`},
		} {
			t.Run(consumer.name+"/"+tt.name, func(t *testing.T) {
				opts := queryTestOptions()
				opts.BooleanCapabilities = consumer.keys
				client, ctx := newHTTPQueryTestClient(t, func(w http.ResponseWriter, _ *http.Request) {
					w.Header().Set("Content-Type", "application/json")
					_, _ = io.WriteString(w, tt.body)
				}, opts)
				wantClass := features.Valid
				switch tt.name {
				case "null", "scalar", "array", "invalid", "oversized":
					wantClass = features.Malformed
				case "logs_string", "logs_null":
					if consumer.name == "logs" || consumer.name == "both" {
						wantClass = features.Malformed
					}
				case "rbac_string", "rbac_null":
					if consumer.name == "rbac" || consumer.name == "both" {
						wantClass = features.Malformed
					}
				}
				started := time.Now()
				result := client.FetchFeatures(ctx)
				checkQueryResult(t, result, wantClass, 200, 1, started)
				if wantClass == features.Valid {
					var want map[string]any
					if err := json.Unmarshal([]byte(tt.body), &want); err != nil {
						t.Fatal(err)
					}
					if !reflect.DeepEqual(result.Features, want) {
						t.Error("query changed unrelated capability values")
					}
				}
			})
		}
	}
}

func TestQueryOptionsCopiesCapabilities(t *testing.T) {
	opts := queryTestOptions()
	client, ctx := newHTTPQueryTestClient(t, func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_, _ = io.WriteString(w, `{"otel-logs":"true","k8s-rbac":true}`)
	}, opts)
	opts.BooleanCapabilities[0] = "k8s-rbac"
	started := time.Now()
	checkQueryResult(t, client.FetchFeatures(ctx), features.Malformed, 200, 1, started)
}

func TestFetchFeaturesHTTPTransientRecovery(t *testing.T) {
	for _, status := range []int{408, 429, 500, 503} {
		t.Run(fmt.Sprint(status), func(t *testing.T) {
			var calls atomic.Int32
			client, ctx := newHTTPQueryTestClient(t, func(w http.ResponseWriter, _ *http.Request) {
				w.Header().Set("Content-Type", "application/json")
				if calls.Add(1) < 3 {
					w.WriteHeader(status)
					_, _ = io.WriteString(w, queryTestBody)
					return
				}
				_, _ = io.WriteString(w, `{"otel-logs":true}`)
			}, queryTestOptions())
			started := time.Now()
			result := client.FetchFeatures(ctx)
			checkQueryResult(t, result, features.Valid, 200, 3, started)
			if calls.Load() != 3 || result.Features["otel-logs"] != true {
				t.Error("transient recovery did not return the final response")
			}
		})
	}
}

type queryTestAPI struct {
	ctx     context.Context
	calls   int
	execute func(context.Context, int) (map[string]any, *http.Response, error)
}

func (api *queryTestAPI) GetFeatures(ctx context.Context) receiver_api.ApiGetFeaturesRequest {
	api.ctx = ctx
	return receiver_api.ApiGetFeaturesRequest{ApiService: api}
}

func (api *queryTestAPI) GetFeaturesExecute(_ receiver_api.ApiGetFeaturesRequest) (map[string]any, *http.Response, error) {
	api.calls++
	return api.execute(api.ctx, api.calls)
}

func queryTestResponse(status int) *http.Response {
	return &http.Response{
		StatusCode: status,
		Header:     make(http.Header),
		Body:       io.NopCloser(strings.NewReader(queryTestBody)),
	}
}

func TestFetchFeaturesNilResponseErrors(t *testing.T) {
	tests := []struct {
		name     string
		err      error
		class    features.Class
		attempts int
	}{
		{"transport", errors.New(queryTestError + queryTestAPIKey), features.Transient, 3},
		{"credential_url", &url.Error{Op: "Get", URL: "https://synthetic.invalid/?api_key=" + queryTestAPIKey, Err: errors.New(queryTestError)}, features.Transient, 3},
		{"connection_reset", &net.OpError{Op: "read", Net: "tcp", Err: errors.New(queryTestError)}, features.Transient, 3},
		{"deadline", context.DeadlineExceeded, features.Timeout, 3},
		{"canceled", context.Canceled, features.Canceled, 1},
		{"untrusted_certificate", &url.Error{Op: "Get", URL: "https://synthetic.invalid/" + queryTestAPIKey, Err: x509.UnknownAuthorityError{}}, features.Configuration, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			synctest.Test(t, func(t *testing.T) {
				api := &queryTestAPI{execute: func(context.Context, int) (map[string]any, *http.Response, error) {
					return nil, nil, tt.err
				}}
				client := newQueryTestClient(t, api, queryTestOptions())
				started := time.Now()
				checkQueryResult(t, client.FetchFeatures(context.Background()), tt.class, 0, tt.attempts, started)
				if api.calls != tt.attempts {
					t.Errorf("API calls = %d; want %d", api.calls, tt.attempts)
				}
			})
		})
	}
	t.Run("nil_response_without_error", func(t *testing.T) {
		api := &queryTestAPI{execute: func(context.Context, int) (map[string]any, *http.Response, error) {
			return map[string]any{"otel-logs": true}, nil, nil
		}}
		opts := queryTestOptions()
		opts.MaxAttempts = 1
		started := time.Now()
		result := newQueryTestClient(t, api, opts).FetchFeatures(context.Background())
		if result.Class == features.Valid || result.Features != nil {
			t.Error("missing HTTP response must not produce a valid observation")
		}
		checkQueryResult(t, result, result.Class, 0, 1, started)
	})
}

type queryTestBodyCloser struct {
	io.Reader
	closed int
}

func (body *queryTestBodyCloser) Close() error {
	body.closed++
	return nil
}

func TestFetchFeaturesClosesResponses(t *testing.T) {
	for _, status := range []int{200, 401, 404, 503} {
		t.Run(fmt.Sprint(status), func(t *testing.T) {
			var bodies []*queryTestBodyCloser
			api := &queryTestAPI{execute: func(context.Context, int) (map[string]any, *http.Response, error) {
				body := &queryTestBodyCloser{Reader: strings.NewReader(queryTestBody)}
				bodies = append(bodies, body)
				response := queryTestResponse(status)
				response.Body = body
				if status == 200 {
					return map[string]any{}, response, nil
				}
				return nil, response, errors.New(queryTestError)
			}}
			newQueryTestClient(t, api, queryTestOptions()).FetchFeatures(context.Background())
			for i, body := range bodies {
				if body.closed != 1 {
					t.Errorf("attempt %d response closed %d times; want once", i+1, body.closed)
				}
			}
		})
	}
}

func TestFetchFeaturesRetryAfter(t *testing.T) {
	for _, status := range []int{429, 503} {
		for _, format := range []string{"seconds", "http_date", "beyond_budget", "overflow", "invalid", "past_date", "negative"} {
			t.Run(fmt.Sprintf("%d_%s", status, format), func(t *testing.T) {
				synctest.Test(t, func(t *testing.T) {
					started := time.Now()
					retryAfter := "2"
					switch format {
					case "http_date":
						retryAfter = started.Add(2 * time.Second).UTC().Format(http.TimeFormat)
					case "beyond_budget":
						retryAfter = "60"
					case "overflow":
						retryAfter = "184467440737095516160"
					case "invalid":
						retryAfter = "not-a-delay"
					case "past_date":
						retryAfter = started.Add(-time.Hour).UTC().Format(http.TimeFormat)
					case "negative":
						retryAfter = "-1"
					}
					var attemptsAt []time.Time
					api := &queryTestAPI{execute: func(_ context.Context, attempt int) (map[string]any, *http.Response, error) {
						attemptsAt = append(attemptsAt, time.Now())
						if attempt == 1 {
							response := queryTestResponse(status)
							response.Header.Set("Retry-After", retryAfter)
							return nil, response, errors.New(queryTestError)
						}
						return map[string]any{}, queryTestResponse(200), nil
					}}
					result := newQueryTestClient(t, api, queryTestOptions()).FetchFeatures(context.Background())
					if format == "beyond_budget" || format == "overflow" {
						if api.calls != 1 {
							t.Error("retried earlier than Retry-After allowed")
						}
						if result.Class != features.Timeout && result.Class != features.Transient {
							t.Error("unaffordable retry must end as timeout or exhausted transient")
						}
						checkQueryResult(t, result, result.Class, status, 1, started)
					} else {
						checkQueryResult(t, result, features.Valid, 200, 2, started)
						if len(attemptsAt) != 2 {
							t.Fatalf("attempts = %d; want 2", len(attemptsAt))
						}
						if (format == "seconds" || format == "http_date") && attemptsAt[1].Sub(started) < 2*time.Second {
							t.Error("retried before Retry-After")
						}
						if format != "seconds" && format != "http_date" && attemptsAt[1].Sub(started) > time.Nanosecond {
							t.Error("invalid or expired Retry-After did not fall back to configured backoff")
						}
					}
					if time.Since(started) > queryTestOptions().Timeout {
						t.Error("Retry-After exceeded the query budget")
					}
				})
			})
		}
	}
}

func TestFetchFeaturesAttemptAndQueryDeadlines(t *testing.T) {
	for _, wholeQuery := range []bool{false, true} {
		t.Run(fmt.Sprintf("whole_query_%t", wholeQuery), func(t *testing.T) {
			synctest.Test(t, func(t *testing.T) {
				opts := queryTestOptions()
				if wholeQuery {
					opts.Timeout = 2500 * time.Millisecond
				}
				type contextKey struct{}
				parent := context.WithValue(context.Background(), contextKey{}, "synthetic-context-value")
				var attempts []context.Context
				api := &queryTestAPI{execute: func(ctx context.Context, _ int) (map[string]any, *http.Response, error) {
					attempts = append(attempts, ctx)
					if ctx.Value(contextKey{}) != "synthetic-context-value" {
						t.Error("attempt discarded parent context values")
					}
					deadline, ok := ctx.Deadline()
					if !ok || time.Until(deadline) > opts.AttemptTimeout {
						t.Error("attempt has no bounded deadline")
					}
					<-ctx.Done()
					return nil, nil, ctx.Err()
				}}
				started := time.Now()
				result := newQueryTestClient(t, api, opts).FetchFeatures(parent)
				checkQueryResult(t, result, features.Timeout, 0, 3, started)
				for _, ctx := range attempts {
					if ctx.Err() == nil {
						t.Error("completed attempt context was not canceled")
					}
				}
				wantElapsed := 3 * opts.AttemptTimeout
				if wholeQuery {
					wantElapsed = opts.Timeout
				}
				if elapsed := time.Since(started); elapsed < wantElapsed || elapsed > wantElapsed+2*opts.MaxBackoff {
					t.Errorf("query elapsed = %v; want %v plus bounded backoff", elapsed, wantElapsed)
				}
				if parent.Err() != nil {
					t.Error("query canceled the parent context")
				}
			})
		})
	}
}

func TestFetchFeaturesCancellation(t *testing.T) {
	for _, phase := range []string{"before_query", "during_attempt", "during_backoff", "during_retry_after"} {
		t.Run(phase, func(t *testing.T) {
			synctest.Test(t, func(t *testing.T) {
				ctx, cancel := context.WithCancel(context.Background())
				defer cancel()
				opts := queryTestOptions()
				opts.InitialBackoff = time.Second
				opts.MaxBackoff = time.Second
				api := &queryTestAPI{execute: func(ctx context.Context, _ int) (map[string]any, *http.Response, error) {
					if phase == "during_attempt" {
						<-ctx.Done()
						return nil, nil, ctx.Err()
					}
					response := queryTestResponse(503)
					if phase == "during_retry_after" {
						response.Header.Set("Retry-After", "5")
					}
					return nil, response, errors.New(queryTestError)
				}}
				if phase == "before_query" {
					cancel()
				}
				client := newQueryTestClient(t, api, opts)
				started := time.Now()
				results := make(chan features.Result, 1)
				go func() { results <- client.FetchFeatures(ctx) }()
				synctest.Wait()
				cancel()
				synctest.Wait()
				select {
				case result := <-results:
					attempts, status := 1, 503
					if phase == "before_query" {
						attempts, status = 0, 0
					} else if phase == "during_attempt" {
						status = 0
					}
					checkQueryResult(t, result, features.Canceled, status, attempts, started)
					if api.calls != attempts {
						t.Errorf("API calls = %d; want %d", api.calls, attempts)
					}
				default:
					t.Fatal("cancellation did not interrupt query")
				}
				if !time.Now().Equal(started) {
					t.Error("cancellation waited for a timeout or retry delay")
				}
			})
		})
	}
}

func TestFetchFeaturesCanceledHTTP(t *testing.T) {
	entered := make(chan struct{})
	requestCanceled := make(chan struct{})
	client, authCtx := newHTTPQueryTestClient(t, func(_ http.ResponseWriter, r *http.Request) {
		close(entered)
		<-r.Context().Done()
		close(requestCanceled)
	}, queryTestOptions())
	ctx, cancel := context.WithCancel(authCtx)
	defer cancel()
	started := time.Now()
	results := make(chan features.Result, 1)
	go func() { results <- client.FetchFeatures(ctx) }()
	select {
	case <-entered:
	case <-time.After(5 * time.Second):
		t.Fatal("HTTP request did not reach server")
	}
	cancel()
	select {
	case result := <-results:
		checkQueryResult(t, result, features.Canceled, 0, 1, started)
	case <-time.After(5 * time.Second):
		t.Fatal("HTTP query did not stop after cancellation")
	}
	select {
	case <-requestCanceled:
	case <-time.After(5 * time.Second):
		t.Fatal("server request was not canceled")
	}
}

func TestFetchFeaturesInterruptedSuccessfulResponse(t *testing.T) {
	var calls atomic.Int32
	client, ctx := newHTTPQueryTestClient(t, func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if calls.Add(1) == 1 {
			w.Header().Set("Content-Length", "1000")
			_, _ = io.WriteString(w, `{"otel-logs":`)
			return
		}
		_, _ = io.WriteString(w, `{"otel-logs":true}`)
	}, queryTestOptions())
	started := time.Now()
	result := client.FetchFeatures(ctx)
	checkQueryResult(t, result, features.Valid, 200, 2, started)
}

func TestQueryOptionsValidation(t *testing.T) {
	api := &queryTestAPI{}
	for _, modify := range []func(*features.QueryOptions){
		func(o *features.QueryOptions) { o.Timeout = 0 },
		func(o *features.QueryOptions) { o.AttemptTimeout = 0 },
		func(o *features.QueryOptions) { o.AttemptTimeout = o.Timeout + time.Second },
		func(o *features.QueryOptions) { o.MaxAttempts = 0 },
		func(o *features.QueryOptions) { o.InitialBackoff = 0 },
		func(o *features.QueryOptions) { o.MaxBackoff = 0 },
	} {
		opts := queryTestOptions()
		modify(&opts)
		if _, err := features.NewClient(api, opts); err == nil {
			t.Error("accepted invalid query bounds")
		}
	}
	if _, err := features.NewClient(nil, queryTestOptions()); err == nil {
		t.Error("accepted nil features API")
	}
}
