package client

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMakeBaseUrl(t *testing.T) {
	assert.Equal(t, makeBaseURL("https://bla/"), "https://bla")
	assert.Equal(t, makeBaseURL("https://bla"), "https://bla")
	assert.Equal(t, makeBaseURL("https://bla/stsAgent/"), "https://bla")
	assert.Equal(t, makeBaseURL("https://bla/stsAgent"), "https://bla")
}
