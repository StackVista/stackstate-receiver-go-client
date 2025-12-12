package openapiclient

import (
	"fmt"
	log "github.com/cihub/seelog"
	"os"
	"strings"
	"sync"
	"time"
)

// PeriodicTokenFileReader can be used to periodically read a token from a file.
// The contents of the file (a token) is exposed as a function which can be passed
// to NewApiClient as serviceAccount token.
// The helper is closeable properly to stop it and is thread-safe
type PeriodicTokenFileReader struct {
	filePath string
	interval time.Duration
	token    string
	mutex    sync.RWMutex
	stopChan chan struct{}
	wg       sync.WaitGroup
}

// NewPeriodicTokenFileReader creates a new PeriodicTokenFileReader and starts its background refresh loop.
func NewPeriodicTokenFileReader(filePath string, interval time.Duration) (*PeriodicTokenFileReader, error) {
	reader := &PeriodicTokenFileReader{
		filePath: filePath,
		interval: interval,
		stopChan: make(chan struct{}),
	}

	// Do an initial read of the token file, we will not fail if the file is not there yet
	err := reader.refreshToken()
	if err != nil {
		return nil, fmt.Errorf("failed to read token from file: %w", err)
	}

	reader.wg.Add(1)
	go reader.run()

	return reader, nil
}

// GetToken returns the current token. This function is safe for concurrent use.
func (r *PeriodicTokenFileReader) GetToken() string {
	r.mutex.RLock()
	defer r.mutex.RUnlock()
	return r.token
}

// Close stops the periodic reading of the token file.
func (r *PeriodicTokenFileReader) Close() {
	close(r.stopChan)
	r.wg.Wait()
}

// run is the background goroutine that periodically refreshes the token.
func (r *PeriodicTokenFileReader) run() {
	defer r.wg.Done()
	ticker := time.NewTicker(r.interval)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			err := r.refreshToken()
			if err != nil {
				log.Errorf("Failed to read token from file: %v", err)
			} else {
				log.Debugf("Read token from file: %v", err)
			}
		case <-r.stopChan:
			return
		}
	}
}

// refreshToken reads the token from the file and updates it.
func (r *PeriodicTokenFileReader) refreshToken() error {
	data, err := os.ReadFile(r.filePath)
	if err != nil {
		return err
	}

	newToken := strings.TrimSpace(string(data))

	r.mutex.Lock()
	r.token = newToken
	r.mutex.Unlock()

	return nil
}
