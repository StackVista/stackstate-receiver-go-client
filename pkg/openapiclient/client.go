package openapiclient

import (
	"context"
	"crypto/tls"
	"errors"
	"fmt"
	"net"
	"net/http"
	"net/url"
	"strings"
	"time"

	"golang.org/x/oauth2"

	"github.com/StackVista/stackstate-receiver-go-client/generated/receiver_api"
	log "github.com/cihub/seelog"
)

// OpenAPIClient provides a client for connecting to the openapi generated portion of the receiver api
type OpenAPIClient interface {
	Connect() *receiver_api.APIClient
}

// NewOpenAPIClient constructs the OpenAPIClient client
func NewOpenAPIClient(ctx context.Context,
	isVerbose bool,
	userAgent string,
	url string,
	apiToken string,
	serviceAccountToken func() string,
	skipSSL bool,
	proxy *url.URL) (OpenAPIClient, context.Context) {
	baseURL := makeBaseURL(url)
	client, clientAuth := newAPIClient(isVerbose, userAgent, baseURL, apiToken, serviceAccountToken, skipSSL, proxy)

	withClient := ctx
	if clientAuth != nil {
		withClient = context.WithValue(
			ctx,
			receiver_api.ContextOAuth2,
			clientAuth,
		)
	}

	return openAPIClientImpl{
		client:      client,
		Context:     withClient,
		receiverURL: baseURL,
	}, withClient
}

func newAPIClient(
	isVerbose bool,
	userAgent string,
	receiverURL string,
	apiKey string,
	serviceAccountToken func() string,
	skipSSL bool,
	proxy *url.URL,
) (*receiver_api.APIClient, oauth2.TokenSource) {
	configuration := receiver_api.NewConfiguration()

	transport := &http.Transport{
		DialContext: (&net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 30 * time.Second,
		}).DialContext,
		MaxIdleConns:          100,
		IdleConnTimeout:       90 * time.Second,
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
		TLSClientConfig:       &tls.Config{InsecureSkipVerify: skipSSL},
	}

	if skipSSL {
		log.Warnf("Using univerified ssl connection")
	}

	if proxy != nil {
		log.Infof("configuring proxy through: %s", proxy.String())
		transport.Proxy = http.ProxyURL(proxy)
	}

	configuration.HTTPClient = &http.Client{Timeout: 30 * time.Second, Transport: transport}
	configuration.UserAgent = userAgent
	configuration.Servers[0] = receiver_api.ServerConfiguration{
		URL:         receiverURL,
		Description: "",
		Variables:   nil,
	}
	configuration.Debug = isVerbose

	client := receiver_api.NewAPIClient(configuration)

	if apiKey != "" {
		return client, oauth2.StaticTokenSource(&oauth2.Token{
			AccessToken: apiKey,
			TokenType:   "ApiKey",
		})
	}
	token := serviceAccountToken()
	if token != "" {
		return client, dynamicTokenSource{tokenFunc: serviceAccountToken, tokenType: "ServiceBearer"}
	}

	return client, nil
}

// dynamicTokenSource calls tokenFunc on every Token() invocation so that
// refreshed credentials (e.g. rotated Kubernetes service-account tokens) are
// picked up automatically.
type dynamicTokenSource struct {
	tokenFunc func() string
	tokenType string
}

func (d dynamicTokenSource) Token() (*oauth2.Token, error) {
	return &oauth2.Token{
		AccessToken: d.tokenFunc(),
		TokenType:   d.tokenType,
	}, nil
}

type openAPIClientImpl struct {
	client      *receiver_api.APIClient
	Context     context.Context
	receiverURL string
}

func (c openAPIClientImpl) Connect() *receiver_api.APIClient {
	// Placeholder in case we want to do something while connecting
	log.Infof("Connected to receiver: %s", c.receiverURL)

	return c.client
}

// Drop /stsAgent/ part from the url is it exists, because it is included in openapi
func makeBaseURL(url string) string {
	return strings.TrimSuffix(strings.Trim(url, "/"), "/stsAgent")
}

// NewOpenAPIClientWithOptions constructs a Receiver client and its authenticated context.
func NewOpenAPIClientWithOptions(parent context.Context, opts ConnectionOptions) (*receiver_api.APIClient, context.Context, error) {
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
	source := validatedTokenSource{tokenFunc: opts.ServiceAccountToken, tokenType: "ServiceBearer"}
	if opts.APIKey != "" {
		source = validatedTokenSource{tokenFunc: func() string { return opts.APIKey }, tokenType: "ApiKey"}
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

type validatedTokenSource struct {
	tokenFunc func() string
	tokenType string
}

func (d validatedTokenSource) Token() (*oauth2.Token, error) {
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
