package openapiclient

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"io"
	"strings"
	"testing"
)

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
