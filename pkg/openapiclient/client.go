package openapiclient

import (
	"context"
	"crypto/tls"
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
