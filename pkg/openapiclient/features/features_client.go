package features

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"errors"
	"io"
	"math/rand/v2"
	"net"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/StackVista/stackstate-receiver-go-client/generated/receiver_api"
	"github.com/StackVista/stackstate-receiver-go-client/pkg/openapiclient"
)

// QueryOptions bounds a complete feature query, including attempts and retry waits.
type QueryOptions struct {
	Timeout, AttemptTimeout    time.Duration
	MaxAttempts                int
	InitialBackoff, MaxBackoff time.Duration
}

// Client queries the generated features API with bounded retries.
type Client struct {
	api    receiver_api.FeaturesAPI
	opts   QueryOptions
	now    func() time.Time
	random func() float64
}

// NewClient validates query bounds and constructs a feature client.
func NewClient(api receiver_api.FeaturesAPI, opts QueryOptions) (*Client, error) {
	if api == nil {
		return nil, errors.New("features API is required")
	}
	if opts.Timeout <= 0 || opts.AttemptTimeout <= 0 || opts.AttemptTimeout > opts.Timeout || opts.MaxAttempts < 1 || opts.InitialBackoff <= 0 || opts.MaxBackoff < opts.InitialBackoff {
		return nil, errors.New("invalid feature query timeout, attempt count or backoff bounds")
	}
	return &Client{api: api, opts: opts, now: time.Now, random: rand.Float64}, nil
}

// FetchFeatures returns one observation after the bounded query completes.
func (c *Client) FetchFeatures(authCtx context.Context) Result {
	queryCtx, cancel := context.WithTimeout(authCtx, c.opts.Timeout)
	defer cancel()
	result := Result{}
	backoff := c.opts.InitialBackoff
	for {
		if queryCtx.Err() != nil {
			result.Class = contextClass(authCtx, queryCtx.Err())
			break
		}
		attemptCtx, stop := context.WithTimeout(queryCtx, c.opts.AttemptTimeout)
		values, response, err := c.api.GetFeaturesExecute(c.api.GetFeatures(attemptCtx))
		result.Attempts++
		result.StatusCode = 0
		retryAfter := time.Duration(0)
		if response != nil {
			result.StatusCode = response.StatusCode
			if response.StatusCode == http.StatusTooManyRequests || response.StatusCode == http.StatusServiceUnavailable {
				retryAfter = parseRetryAfter(response.Header.Get("Retry-After"), c.now(), c.opts.Timeout)
			}
			if response.Body != nil {
				response.Body.Close()
			}
		}
		result.Class = classify(authCtx, attemptCtx, values, response, err)
		stop()
		if result.Class == Valid {
			result.Features = values
		}
		if (result.Class != Transient && result.Class != Timeout) || result.Attempts >= c.opts.MaxAttempts {
			break
		}
		if queryCtx.Err() != nil {
			result.Class = contextClass(authCtx, queryCtx.Err())
			break
		}
		delay := time.Duration(c.random() * float64(backoff))
		if retryAfter > delay {
			delay = retryAfter
		}
		deadline, _ := queryCtx.Deadline()
		if delay >= time.Until(deadline) {
			break
		}
		timer := time.NewTimer(delay)
		select {
		case <-queryCtx.Done():
			timer.Stop()
			result.Class = contextClass(authCtx, queryCtx.Err())
			result.FinishedAt = c.now()
			return result
		case <-timer.C:
		}
		if backoff > c.opts.MaxBackoff/2 {
			backoff = c.opts.MaxBackoff
		} else {
			backoff *= 2
		}
	}
	result.FinishedAt = c.now()
	return result
}

func contextClass(parent context.Context, err error) Class {
	if parent.Err() != nil {
		return Canceled
	}
	if errors.Is(err, context.DeadlineExceeded) {
		return Timeout
	}
	return Canceled
}

func classify(parent, attempt context.Context, values map[string]any, response *http.Response, err error) Class {
	if parent.Err() != nil {
		return Canceled
	}
	if response != nil {
		switch status := response.StatusCode; {
		case status == 401 || status == 403:
			return Authentication
		case status == 404:
			return Unsupported
		case status == 408 || status == 429 || status >= 500 && status <= 599:
			return Transient
		case status != 200:
			return Rejected
		}
	}
	if errors.Is(err, openapiclient.ErrMissingCredential) {
		return Authentication
	}
	var verification *tls.CertificateVerificationError
	var unknown x509.UnknownAuthorityError
	var hostname x509.HostnameError
	var invalid x509.CertificateInvalidError
	if errors.As(err, &verification) || errors.As(err, &unknown) || errors.As(err, &hostname) || errors.As(err, &invalid) {
		return Configuration
	}
	if attempt.Err() != nil || errors.Is(err, context.DeadlineExceeded) {
		return Timeout
	}
	if errors.Is(err, context.Canceled) {
		return Canceled
	}
	var network net.Error
	if errors.As(err, &network) && network.Timeout() {
		return Timeout
	}
	var operation *net.OpError
	if errors.As(err, &operation) || errors.Is(err, io.EOF) || errors.Is(err, io.ErrUnexpectedEOF) || (errors.As(err, &network) && network.Temporary()) {
		return Transient
	}
	if response == nil {
		if err != nil {
			return Transient
		}
		return Rejected
	}
	if err != nil || values == nil {
		return Malformed
	}
	if capability, present := values["otel-logs"]; present {
		if _, ok := capability.(bool); !ok {
			return Malformed
		}
	}
	return Valid
}

func parseRetryAfter(value string, now time.Time, limit time.Duration) time.Duration {
	value = strings.TrimSpace(value)
	if value == "" {
		return 0
	}
	if seconds, err := strconv.ParseUint(value, 10, 64); err == nil {
		if seconds > uint64(limit/time.Second) {
			return limit
		}
		return time.Duration(seconds) * time.Second
	} else if errors.Is(err, strconv.ErrRange) && strings.Trim(value, "0123456789") == "" {
		return limit
	}
	if deadline, err := http.ParseTime(value); err == nil {
		delay := deadline.Sub(now)
		if delay > limit {
			return limit
		}
		if delay > 0 {
			return delay
		}
	}
	return 0
}
