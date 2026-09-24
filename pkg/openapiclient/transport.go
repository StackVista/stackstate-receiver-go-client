package openapiclient

import (
	"crypto/tls"
	"fmt"
	"net/http"
)

func newTransport(opts ConnectionOptions) (http.RoundTripper, error) {
	transport := http.DefaultTransport.(*http.Transport).Clone()
	transport.Proxy = nil
	transport.TLSClientConfig = &tls.Config{InsecureSkipVerify: opts.InsecureSkipVerify} // Explicit operator configuration.
	if opts.ProxyURL != "" {
		proxy, err := parseEndpoint(opts.ProxyURL, true)
		if err != nil {
			return nil, fmt.Errorf("invalid proxy URL: %w", err)
		}
		transport.Proxy = http.ProxyURL(proxy)
	}
	return transport, nil
}
