package openapiclient

import (
	"crypto/tls"
	"crypto/x509"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
)

const maxFeatureResponseBytes = 1 << 20

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
	return boundedTransport{transport}, nil
}

type boundedTransport struct{ base http.RoundTripper }

func (t boundedTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	response, err := t.base.RoundTrip(req)
	if response != nil && response.Body != nil && strings.HasSuffix(req.URL.Path, "/stsAgent/features") {
		response.Body = &boundedBody{ReadCloser: response.Body, remaining: maxFeatureResponseBytes}
	}
	return response, err
}

type boundedBody struct {
	io.ReadCloser
	remaining int
}

func (b *boundedBody) Read(p []byte) (int, error) {
	if len(p) == 0 {
		return 0, nil
	}
	if len(p) > b.remaining+1 {
		p = p[:b.remaining+1]
	}
	n, err := b.ReadCloser.Read(p)
	if n > b.remaining {
		n = b.remaining
		b.remaining = 0
		return n, ErrResponseTooLarge
	}
	b.remaining -= n
	return n, err
}
