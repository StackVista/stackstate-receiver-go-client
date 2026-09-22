package openapiclient

import (
	"crypto/tls"
	"crypto/x509"
	"errors"
	"fmt"
	"net/http"
)

func newTransport(opts ConnectionOptions) (http.RoundTripper, error) {
	transport := http.DefaultTransport.(*http.Transport).Clone()
	transport.Proxy = nil
	transport.TLSClientConfig = &tls.Config{InsecureSkipVerify: opts.InsecureSkipVerify} // Explicit operator configuration.
	if len(opts.CABundlePEM) != 0 {
		roots, err := x509.SystemCertPool()
		if err != nil {
			return nil, errors.New("cannot load system certificate trust")
		}
		if !roots.AppendCertsFromPEM(opts.CABundlePEM) {
			return nil, errors.New("CA bundle contains no valid certificates")
		}
		transport.TLSClientConfig.RootCAs = roots
	}
	if opts.ProxyURL != "" {
		proxy, err := parseEndpoint(opts.ProxyURL, true)
		if err != nil {
			return nil, fmt.Errorf("invalid proxy URL: %w", err)
		}
		transport.Proxy = http.ProxyURL(proxy)
	}
	return transport, nil
}
