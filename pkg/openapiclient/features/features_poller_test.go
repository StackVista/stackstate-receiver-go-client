package features

import (
	"context"
	"github.com/StackVista/stackstate-receiver-go-client/generated/receiver_api"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
	"time"
)

func TestFeaturePollerProducesRepeatedResults(t *testing.T) {
	featuresAPI := receiver_api.NewFeaturesAPIMock()
	features := make(map[string]interface{})
	featuresAPI.GetFeaturesResponse = receiver_api.GetFeaturesMockResponse{
		Result:   features,
		Response: &http.Response{StatusCode: http.StatusOK},
		Error:    nil,
	}

	outputChannel, tearDown := StartFeaturesPoller(context.Background(), featuresAPI, 1*time.Second)
	result := <-outputChannel
	assert.Equal(t, features, result)

	result = <-outputChannel
	assert.Equal(t, features, result)

	tearDown()

	_, ok := <-outputChannel
	assert.False(t, ok)
}

func TestDoesNotProduceWhenBrokenAndBeAbletoTearDown(t *testing.T) {
	featuresAPI := receiver_api.NewFeaturesAPIMock()
	features := make(map[string]interface{})
	featuresAPI.GetFeaturesResponse = receiver_api.GetFeaturesMockResponse{
		Result:   features,
		Response: &http.Response{StatusCode: http.StatusBadRequest},
		Error:    nil,
	}

	outputChannel, tearDown := StartFeaturesPoller(context.Background(), featuresAPI, 1*time.Second)
	time.Sleep(1 * time.Second)

	tearDown()

	_, ok := <-outputChannel
	assert.False(t, ok)
}
