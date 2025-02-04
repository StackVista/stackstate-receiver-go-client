package client

import (
	"context"
	"crypto/tls"
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
	Connect() (*receiver_api.APIClient, error)
}

// NewOpenAPIClient constructs the OpenAPIClient client
func NewOpenAPIClient(ctx context.Context,
	isVerbose bool,
	userAgent string,
	url string,
	apiToken string,
	k8sServiceAccountToken string,
	skipSSL bool,
	proxy *url.URL) (OpenAPIClient, context.Context) {
	baseURL := makeBaseURL(url)
	client, clientAuth := newAPIClient(isVerbose, userAgent, baseURL, apiToken, k8sServiceAccountToken, skipSSL, proxy)

	withClient := context.WithValue(
		ctx,
		receiver_api.ContextAPIKeys,
		clientAuth,
	)

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
	skipSSL bool,
	proxy *url.URL,
) (*receiver_api.APIClient, map[string]receiver_api.APIKey) {
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

	auth := make(map[string]receiver_api.APIKey)
	if apiKey != "" {
		auth["ApiKey"] = receiver_api.APIKey{
			Key:    apiKey,
			Prefix: "",
		}
	}
	if k8sServiceAccountToken != "" {
		auth["ServiceBearer"] = receiver_api.APIKey{
			Key:    k8sServiceAccountToken,
			Prefix: "",
		}
	}

	return client, auth
}

type openAPIClientImpl struct {
	client      *receiver_api.APIClient
	Context     context.Context
	receiverURL string
}

func (c openAPIClientImpl) Connect() (*receiver_api.APIClient, error) {
	// Placeholder in case we want to do something while connecting
	log.Infof("Connected to receiver: %s", c.receiverURL)

	return c.client, nil
}

// Drop /stsAgent/ part from the url is it exists, because it is included in openapi
func makeBaseURL(url string) string {
	return strings.TrimSuffix(strings.Trim(url, "/"), "/stsAgent")
}
