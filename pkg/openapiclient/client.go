package openapiclient

import (
	"context"
	"crypto/tls"
	"golang.org/x/oauth2"
	"net"
	"net/http"
	"net/url"
	"strings"
	"time"

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
	k8sServiceAccountToken string,
	serviceToken string,
	skipSSL bool,
	proxy *url.URL) (OpenAPIClient, context.Context) {
	baseURL := makeBaseURL(url)
	client, clientAuth := newAPIClient(isVerbose, userAgent, baseURL, apiToken, k8sServiceAccountToken, serviceToken, skipSSL, proxy)

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
	k8sServiceAccountToken string,
	serviceToken string,
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
	if k8sServiceAccountToken != "" {
		return client, oauth2.StaticTokenSource(&oauth2.Token{
			AccessToken: k8sServiceAccountToken,
			TokenType:   "ServiceBearer",
		})
	}

	if serviceToken != "" {
		return client, oauth2.StaticTokenSource(&oauth2.Token{
			AccessToken: serviceToken,
			TokenType:   "ServiceToken",
		})
	}
	return client, nil
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
