package openapiclient

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/StackVista/stackstate-receiver-go-client/generated/receiver_api"
	"golang.org/x/oauth2"
)

// NewOpenAPIClient constructs a Receiver client and its authenticated context.
func NewOpenAPIClient(parent context.Context, opts ConnectionOptions) (*receiver_api.APIClient, context.Context, error) {
	if parent == nil {
		return nil, nil, errors.New("parent context is required")
	}
	if _, err := parseEndpoint(opts.ReceiverURL, false); err != nil {
		return nil, nil, fmt.Errorf("invalid receiver URL: %w", err)
	}
	if opts.RequestTimeout <= 0 {
		return nil, nil, errors.New("request timeout must be positive")
	}
	if opts.APIKey != "" && opts.ServiceAccountToken != nil {
		return nil, nil, errors.New("exactly one receiver authentication source is required")
	}
	source := dynamicTokenSource{tokenFunc: opts.ServiceAccountToken, tokenType: "ServiceBearer"}
	if opts.APIKey != "" {
		source = dynamicTokenSource{tokenFunc: func() string { return opts.APIKey }, tokenType: "ApiKey"}
	}
	if _, err := source.Token(); err != nil {
		return nil, nil, err
	}
	transport, err := newTransport(opts)
	if err != nil {
		return nil, nil, err
	}
	cfg := receiver_api.NewConfiguration()
	cfg.HTTPClient = &http.Client{
		Timeout:       opts.RequestTimeout,
		Transport:     transport,
		CheckRedirect: func(_ *http.Request, _ []*http.Request) error { return http.ErrUseLastResponse },
	}
	cfg.UserAgent = opts.UserAgent
	cfg.Servers[0] = receiver_api.ServerConfiguration{URL: makeBaseURL(opts.ReceiverURL)}
	cfg.Debug = false
	authCtx := context.WithValue(parent, receiver_api.ContextOAuth2, source)
	return receiver_api.NewAPIClient(cfg), authCtx, nil
}

type dynamicTokenSource struct {
	tokenFunc func() string
	tokenType string
}

func (d dynamicTokenSource) Token() (*oauth2.Token, error) {
	if d.tokenFunc == nil {
		return nil, ErrMissingCredential
	}
	token := d.tokenFunc()
	if strings.TrimSpace(token) == "" {
		return nil, ErrMissingCredential
	}
	// Reject header delimiters here so transport errors cannot echo credentials.
	if strings.ContainsAny(token, "\r\n") {
		return nil, ErrMissingCredential
	}
	return &oauth2.Token{AccessToken: token, TokenType: d.tokenType}, nil
}
