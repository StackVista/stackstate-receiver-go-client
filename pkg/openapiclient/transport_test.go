package openapiclient

import (
	"context"
	"io"
	"net"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestClientClosesIdleFeatureConnections(t *testing.T) {
	closed := make(chan struct{}, 1)
	server := httptest.NewUnstartedServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_, _ = io.WriteString(w, `{}`)
	}))
	server.Config.ConnState = func(_ net.Conn, state http.ConnState) {
		if state == http.StateClosed {
			select {
			case closed <- struct{}{}:
			default:
			}
		}
	}
	server.Start()
	defer server.Close()
	api, ctx, err := NewOpenAPIClientWithOptions(context.Background(), testOptions(server.URL))
	require.NoError(t, err)
	_, response, err := api.FeaturesAPI.GetFeatures(ctx).Execute()
	require.NoError(t, err)
	require.NoError(t, response.Body.Close())
	api.GetConfig().HTTPClient.CloseIdleConnections()
	select {
	case <-closed:
	case <-time.After(time.Second):
		t.Fatal("client did not close its idle feature connection")
	}
}

func TestBoundedBody(t *testing.T) {
	for _, size := range []int{maxFeatureResponseBytes - 1, maxFeatureResponseBytes, maxFeatureResponseBytes + 1} {
		body := &boundedBody{ReadCloser: io.NopCloser(strings.NewReader(strings.Repeat("x", size))), remaining: maxFeatureResponseBytes}
		content, err := io.ReadAll(body)
		if size > maxFeatureResponseBytes {
			require.ErrorIs(t, err, ErrResponseTooLarge)
		} else {
			require.NoError(t, err)
		}
		assert.LessOrEqual(t, len(content), maxFeatureResponseBytes)
	}
}
