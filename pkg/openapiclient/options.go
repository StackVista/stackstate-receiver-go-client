package openapiclient

import (
	"errors"
	"net/url"
	"strconv"
	"strings"
	"time"
)

// ConnectionOptions configures authentication and the owned Receiver transport.
type ConnectionOptions struct {
	ReceiverURL         string
	UserAgent           string
	APIKey              string
	ServiceAccountToken func() string
	ProxyURL            string
	CABundlePEM         []byte
	InsecureSkipVerify  bool
	RequestTimeout      time.Duration
}

// ErrMissingCredential indicates that no usable authentication credential is available.
var ErrMissingCredential = errors.New("receiver credential is empty")

func parseEndpoint(raw string, allowUserinfo bool) (*url.URL, error) {
	u, err := url.Parse(raw)
	if err != nil || u == nil || (u.Scheme != "http" && u.Scheme != "https") || u.Hostname() == "" || u.Opaque != "" {
		return nil, errors.New("endpoint must be an absolute HTTP(S) URL")
	}
	if (!allowUserinfo && u.User != nil) || u.RawQuery != "" || u.ForceQuery || u.Fragment != "" || strings.Contains(raw, "#") {
		return nil, errors.New("endpoint contains unsupported userinfo, query or fragment")
	}
	if port := u.Port(); port != "" {
		number, err := strconv.Atoi(port)
		if err != nil || number < 1 || number > 65535 {
			return nil, errors.New("endpoint port must be between 1 and 65535")
		}
	} else if strings.HasSuffix(u.Host, ":") {
		return nil, errors.New("endpoint port is empty")
	}
	return u, nil
}
