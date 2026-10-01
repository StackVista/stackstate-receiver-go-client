package features

import (
	"context"
	"errors"
	"math"
	"time"
)

// PollOptions controls the delay between completed observations.
type PollOptions struct {
	Interval time.Duration
	Jitter   float64
}

// Poller delivers each completed query until stopped or its context is canceled.
type Poller struct {
	results chan Result
	done    chan struct{}
	cancel  context.CancelFunc
}

// StartPolling waits one jittered interval before querying, including the first query.
func (c *Client) StartPolling(authCtx context.Context, opts PollOptions) (*Poller, error) {
	if authCtx == nil {
		return nil, errors.New("polling context is required")
	}
	if opts.Interval <= 0 || math.IsNaN(opts.Jitter) || opts.Jitter < 0 || opts.Jitter >= 1 ||
		(opts.Jitter > 0 && float64(opts.Interval)*(1+opts.Jitter) >= float64(math.MaxInt64)) {
		return nil, errors.New("poll interval must be positive and representable with jitter in [0, 1)")
	}
	ctx, cancel := context.WithCancel(authCtx)
	poller := &Poller{results: make(chan Result), done: make(chan struct{}), cancel: cancel}
	go func() {
		defer close(poller.done)
		defer close(poller.results)
		defer cancel()
		for {
			if ctx.Err() != nil {
				return
			}
			timer := time.NewTimer(pollDelay(opts, c.random()))
			select {
			case <-ctx.Done():
				timer.Stop()
				return
			case <-timer.C:
			}
			if ctx.Err() != nil {
				return
			}
			result := c.FetchFeatures(ctx)
			select {
			case <-ctx.Done():
				return
			case poller.results <- result:
			}
		}
	}()
	return poller, nil
}

func pollDelay(opts PollOptions, random float64) time.Duration {
	if opts.Jitter == 0 {
		return opts.Interval
	}
	delay := time.Duration(float64(opts.Interval) * (1 + opts.Jitter*(2*random-1)))
	if delay < time.Nanosecond {
		return time.Nanosecond
	}
	return delay
}

// Results is unbuffered; slow consumers delay the next poll.
func (p *Poller) Results() <-chan Result { return p.results }

// Stop cancels requests, waits and delivery without waiting for shutdown.
func (p *Poller) Stop() { p.cancel() }

// Done closes after the worker has closed Results.
func (p *Poller) Done() <-chan struct{} { return p.done }
