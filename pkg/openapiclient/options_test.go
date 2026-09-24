package openapiclient

import (
	"context"
	"encoding/pem"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"sync/atomic"
	"testing"
	"time"

	"github.com/StackVista/stackstate-receiver-go-client/generated/receiver_api"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func testOptions(endpoint string) ConnectionOptions {
	return ConnectionOptions{ReceiverURL: endpoint, APIKey: "synthetic-key", RequestTimeout: time.Second}
}

func TestConnectionURLAndContext(t *testing.T) {
	for _, suffix := range []string{"", "/", "/stsAgent", "/stsAgent/", "/stsAgent///", "/receiver/stsAgent/"} {
		t.Run(suffix, func(t *testing.T) {
			expectedPath := "/stsAgent/features"
			if strings.HasPrefix(suffix, "/receiver") {
				expectedPath = "/receiver/stsAgent/features"
			}
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, expectedPath, r.URL.Path)
				assert.Equal(t, "ApiKey synthetic-key", r.Header.Get("Authorization"))
				assert.Equal(t, "test-agent", r.UserAgent())
				w.Header().Set("Content-Type", "application/json")
				fmt.Fprint(w, `{"otel-logs":true}`)
			}))
			defer server.Close()
			type contextKey struct{}
			parent, cancel := context.WithCancel(context.WithValue(context.Background(), contextKey{}, "preserved"))
			defer cancel()
			opts := testOptions(server.URL + suffix)
			opts.UserAgent = "test-agent"
			api, ctx, err := NewOpenAPIClientWithOptions(parent, opts)
			require.NoError(t, err)
			assert.Equal(t, "preserved", ctx.Value(contextKey{}))
			result, response, err := api.FeaturesAPI.GetFeatures(ctx).Execute()
			require.NoError(t, err)
			response.Body.Close()
			assert.Equal(t, true, result["otel-logs"])
			cancel()
			_, _, err = api.FeaturesAPI.GetFeatures(ctx).Execute()
			require.ErrorIs(t, err, context.Canceled)
		})
	}
}

func TestConnectionValidationIsCredentialSafe(t *testing.T) {
	cases := map[string]func(*ConnectionOptions){
		"relative":         func(o *ConnectionOptions) { o.ReceiverURL = "/receiver" },
		"scheme":           func(o *ConnectionOptions) { o.ReceiverURL = "ftp://receiver" },
		"missing host":     func(o *ConnectionOptions) { o.ReceiverURL = "http:///path" },
		"userinfo":         func(o *ConnectionOptions) { o.ReceiverURL = "https://user:synthetic-secret@receiver" },
		"query":            func(o *ConnectionOptions) { o.ReceiverURL += "?key=synthetic-secret" },
		"empty query":      func(o *ConnectionOptions) { o.ReceiverURL += "?" },
		"fragment":         func(o *ConnectionOptions) { o.ReceiverURL += "#synthetic-secret" },
		"empty fragment":   func(o *ConnectionOptions) { o.ReceiverURL += "#" },
		"malformed URL":    func(o *ConnectionOptions) { o.ReceiverURL = "https://synthetic-secret%" },
		"timeout":          func(o *ConnectionOptions) { o.RequestTimeout = 0 },
		"no auth":          func(o *ConnectionOptions) { o.APIKey = "" },
		"blank auth":       func(o *ConnectionOptions) { o.APIKey = " \t" },
		"invalid header":   func(o *ConnectionOptions) { o.APIKey = "synthetic-secret\r\n" },
		"empty token":      func(o *ConnectionOptions) { o.APIKey = ""; o.ServiceAccountToken = func() string { return "" } },
		"two auth sources": func(o *ConnectionOptions) { o.ServiceAccountToken = func() string { return "synthetic-secret" } },
		"proxy":            func(o *ConnectionOptions) { o.ProxyURL = "socks5://synthetic-secret" },
	}
	for name, modify := range cases {
		t.Run(name, func(t *testing.T) {
			opts := testOptions("https://receiver/stsAgent")
			modify(&opts)
			client, ctx, err := NewOpenAPIClientWithOptions(context.Background(), opts)
			require.Error(t, err)
			assert.Nil(t, client)
			assert.Nil(t, ctx)
			assert.NotContains(t, err.Error(), "synthetic-secret")
		})
	}
	_, _, err := NewOpenAPIClientWithOptions(nil, testOptions("https://receiver"))
	require.Error(t, err)
}

func TestRotatingTokenUsedByFeaturesAndRBAC(t *testing.T) {
	var token atomic.Value
	token.Store("first")
	var requests atomic.Int32
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requests.Add(1)
		assert.Equal(t, "ServiceBearer "+token.Load().(string), r.Header.Get("Authorization"))
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `{}`)
	}))
	defer server.Close()
	opts := testOptions(server.URL)
	opts.APIKey = ""
	opts.ServiceAccountToken = func() string { return token.Load().(string) }
	api, ctx, err := NewOpenAPIClientWithOptions(context.Background(), opts)
	require.NoError(t, err)
	payload := receiver_api.RBACSnapshotRequestAsRBACRequest(&receiver_api.RBACSnapshotRequest{})
	for _, value := range []string{"first", "rotated"} {
		token.Store(value)
		_, response, err := api.FeaturesAPI.GetFeatures(ctx).Execute()
		require.NoError(t, err)
		response.Body.Close()
		response, err = api.ReceiverRbacInstanceAPI.IngestInstanceRBAC(ctx).RBACRequest(payload).Execute()
		require.NoError(t, err)
		response.Body.Close()
	}
	token.Store("")
	_, _, err = api.FeaturesAPI.GetFeatures(ctx).Execute()
	require.ErrorIs(t, err, ErrMissingCredential)
	_, err = api.ReceiverRbacInstanceAPI.IngestInstanceRBAC(ctx).RBACRequest(payload).Execute()
	require.ErrorIs(t, err, ErrMissingCredential)
	assert.EqualValues(t, 4, requests.Load())
}

func TestTLSOptions(t *testing.T) {
	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `{}`)
	}))
	defer server.Close()
	ca := pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: server.Certificate().Raw})
	if runtime.GOOS == "linux" {
		caFile := filepath.Join(t.TempDir(), "receiver-ca.pem")
		require.NoError(t, os.WriteFile(caFile, ca, 0600))
		command := exec.Command(os.Args[0], "-test.run=^TestTLSOptionsSystemTrustHelper$")
		command.Env = append(os.Environ(), "SSL_CERT_FILE="+caFile, "RECEIVER_TEST_URL="+server.URL)
		output, err := command.CombinedOutput()
		require.NoErrorf(t, err, "system trust helper failed: %s", output)
	}
	for _, tc := range []struct {
		name     string
		skip     bool
		succeeds bool
	}{
		{name: "untrusted"}, {name: "explicit skip", skip: true, succeeds: true},
	} {
		t.Run(tc.name, func(t *testing.T) {
			opts := testOptions(server.URL)
			opts.InsecureSkipVerify = tc.skip
			api, ctx, err := NewOpenAPIClientWithOptions(context.Background(), opts)
			require.NoError(t, err)
			_, response, err := api.FeaturesAPI.GetFeatures(ctx).Execute()
			if response != nil {
				response.Body.Close()
			}
			if tc.succeeds {
				require.NoError(t, err)
			} else {
				require.Error(t, err)
			}
		})
	}
}

func TestTLSOptionsSystemTrustHelper(t *testing.T) {
	if runtime.GOOS != "linux" {
		t.Skip("SSL_CERT_FILE system trust test is supported on Linux only")
	}
	endpoint := os.Getenv("RECEIVER_TEST_URL")
	if endpoint == "" {
		t.Skip("helper process only")
	}
	api, ctx, err := NewOpenAPIClientWithOptions(context.Background(), testOptions(endpoint))
	require.NoError(t, err)
	_, response, err := api.FeaturesAPI.GetFeatures(ctx).Execute()
	require.NoError(t, err)
	response.Body.Close()
}

func TestProxyAndRedirectBoundaries(t *testing.T) {
	var proxyCalls atomic.Int32
	proxy := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		proxyCalls.Add(1)
		assert.Equal(t, "http://unresolvable.invalid/stsAgent/features", r.URL.String())
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `{}`)
	}))
	defer proxy.Close()
	opts := testOptions("http://unresolvable.invalid")
	opts.ProxyURL = proxy.URL
	api, ctx, err := NewOpenAPIClientWithOptions(context.Background(), opts)
	require.NoError(t, err)
	_, response, err := api.FeaturesAPI.GetFeatures(ctx).Execute()
	require.NoError(t, err)
	response.Body.Close()
	assert.EqualValues(t, 1, proxyCalls.Load())
	t.Setenv("HTTP_PROXY", proxy.URL)
	t.Setenv("HTTPS_PROXY", proxy.URL)
	transport, err := newTransport(testOptions("http://receiver"))
	require.NoError(t, err)
	assert.Nil(t, transport.(*http.Transport).Proxy)
	var redirected atomic.Int32
	target := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) { redirected.Add(1); fmt.Fprint(w, `{}`) }))
	defer target.Close()
	redirect := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) { http.Redirect(w, r, target.URL, 302) }))
	defer redirect.Close()
	api, ctx, err = NewOpenAPIClientWithOptions(context.Background(), testOptions(redirect.URL))
	require.NoError(t, err)
	_, response, err = api.FeaturesAPI.GetFeatures(ctx).Execute()
	require.Error(t, err)
	require.NotNil(t, response)
	response.Body.Close()
	assert.Equal(t, 302, response.StatusCode)
	assert.Zero(t, redirected.Load())
}

func TestEndpointPorts(t *testing.T) {
	for _, proxy := range []bool{false, true} {
		for _, port := range []string{"0", "65536", "9999999999999999999999", "", "-1", "named"} {
			t.Run(fmt.Sprintf("proxy_%t_port_%s", proxy, port), func(t *testing.T) {
				opts := testOptions("https://receiver/stsAgent")
				invalid := "https://receiver:" + port
				if proxy {
					opts.ProxyURL = invalid
				} else {
					opts.ReceiverURL = invalid + "/stsAgent"
				}
				_, _, err := NewOpenAPIClientWithOptions(context.Background(), opts)
				require.Error(t, err)
			})
		}
	}
	for _, endpoint := range []string{"https://receiver:443/stsAgent", "http://[::1]:8080/stsAgent", "http://[::1]/stsAgent"} {
		_, _, err := NewOpenAPIClientWithOptions(context.Background(), testOptions(endpoint))
		require.NoError(t, err)
	}
}

func TestOptionsRequestTimeout(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(_ http.ResponseWriter, r *http.Request) {
		<-r.Context().Done()
	}))
	defer server.Close()
	opts := testOptions(server.URL)
	opts.RequestTimeout = 20 * time.Millisecond
	api, ctx, err := NewOpenAPIClientWithOptions(context.Background(), opts)
	require.NoError(t, err)
	_, _, err = api.FeaturesAPI.GetFeatures(ctx).Execute()
	require.ErrorIs(t, err, context.DeadlineExceeded)
}

func TestLegacyConstructorAuthentication(t *testing.T) {
	for _, apiKey := range []string{"synthetic-key", ""} {
		t.Run(apiKey, func(t *testing.T) {
			var callbackCalls int
			token := ""
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				expected := ""
				if apiKey != "" {
					expected = "ApiKey " + apiKey
				}
				assert.Equal(t, expected, r.Header.Get("Authorization"))
				w.Header().Set("Content-Type", "application/json")
				fmt.Fprint(w, `{}`)
			}))
			defer server.Close()
			client, ctx := NewOpenAPIClient(context.Background(), false, "legacy", server.URL, apiKey, func() string {
				callbackCalls++
				return token
			}, false, nil)
			token = "later-token"
			_, response, err := client.Connect().FeaturesAPI.GetFeatures(ctx).Execute()
			require.NoError(t, err)
			response.Body.Close()
			if apiKey != "" {
				assert.Zero(t, callbackCalls)
			} else {
				assert.Equal(t, 1, callbackCalls)
			}
		})
	}
}

func TestLegacyTokenSourceAcceptsEmptyRotation(t *testing.T) {
	token := "initial-token"
	_, source := newAPIClient(false, "legacy", "https://receiver", "", func() string { return token }, false, nil)
	require.NotNil(t, source)
	token = ""
	credential, err := source.Token()
	require.NoError(t, err)
	assert.Empty(t, credential.AccessToken)
}
