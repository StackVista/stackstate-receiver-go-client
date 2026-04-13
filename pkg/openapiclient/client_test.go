package openapiclient

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMakeBaseUrl(t *testing.T) {
	assert.Equal(t, makeBaseURL("https://bla/"), "https://bla")
	assert.Equal(t, makeBaseURL("https://bla"), "https://bla")
	assert.Equal(t, makeBaseURL("https://bla/stsAgent/"), "https://bla")
	assert.Equal(t, makeBaseURL("https://bla/stsAgent"), "https://bla")
}

func TestServiceAccountTokenSourceReturnsRefreshedToken(t *testing.T) {
	currentToken := "initial-token"
	tokenFunc := func() string { return currentToken }

	_, tokenSource := newAPIClient(false, "test-agent", "https://receiver", "", tokenFunc, true, nil)
	require.NotNil(t, tokenSource, "tokenSource should not be nil when a service account token is provided")

	// First call should return the initial token
	tok, err := tokenSource.Token()
	require.NoError(t, err)
	assert.Equal(t, "initial-token", tok.AccessToken)
	assert.Equal(t, "ServiceBearer", tok.TokenType)

	// Simulate token rotation (e.g. PeriodicTokenFileReader picked up a new token)
	currentToken = "refreshed-token"

	// Subsequent call must return the refreshed token, not the stale one
	tok, err = tokenSource.Token()
	require.NoError(t, err)
	assert.Equal(t, "refreshed-token", tok.AccessToken,
		"token source must return the latest token, not a cached value from initialization")
}
