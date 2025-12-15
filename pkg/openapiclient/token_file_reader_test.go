package openapiclient

import (
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestPeriodicTokenFileReader(t *testing.T) {
	// Create a temporary file for the token
	tmpfile, err := os.CreateTemp("", "token")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpfile.Name()) // clean up

	// Write initial token
	initialToken := "initial-token"
	if _, err := tmpfile.WriteString(initialToken); err != nil {
		t.Fatal(err)
	}
	tmpfile.Close()

	// Create a new reader with a short interval
	reader, err := NewPeriodicTokenFileReader(tmpfile.Name(), 50*time.Millisecond)
	assert.NoError(t, err)

	// Allow some time for the initial read
	time.Sleep(25 * time.Millisecond)
	assert.Equal(t, initialToken, reader.GetToken())

	// Update the token in the file
	updatedToken := "updated-token"
	err = os.WriteFile(tmpfile.Name(), []byte(updatedToken), 0644)
	assert.NoError(t, err)

	// Wait for the reader to pick up the change
	time.Sleep(100 * time.Millisecond)
	assert.Equal(t, updatedToken, reader.GetToken())

	// Test closing the reader
	reader.Close()

	// Update the token again
	finalToken := "final-token"
	err = os.WriteFile(tmpfile.Name(), []byte(finalToken), 0644)
	assert.NoError(t, err)

	// Wait and check that the token was NOT updated because the reader is closed
	time.Sleep(100 * time.Millisecond)
	assert.Equal(t, updatedToken, reader.GetToken())
}

func TestPeriodicTokenFileReader_FileNotInitiallyPresent(t *testing.T) {
	tmpfileName := "non-existent-file"
	os.Remove(tmpfileName) // ensure it doesn't exist

	// Create a new reader with a short interval
	_, err := NewPeriodicTokenFileReader(tmpfileName, 50*time.Millisecond)
	assert.Error(t, err)
}
