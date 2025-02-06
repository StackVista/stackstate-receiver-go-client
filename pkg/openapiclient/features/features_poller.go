package features

import (
	"context"
	"github.com/StackVista/stackstate-receiver-go-client/generated/receiver_api"
	log "github.com/cihub/seelog"
	"net/http"
	"time"
)

func StartFeaturesPoller(clientCtx context.Context, featuresApi receiver_api.FeaturesAPI, interval time.Duration) (chan map[string]interface{}, func()) {

	outputChannel := make(chan map[string]interface{})
	stopChannel := make(chan interface{})
	ticker := time.NewTicker(interval)
	// Channel that produces just a single value
	init := make(chan bool, 1)
	init <- true

	go func() {
		for {
			select {
			case <-stopChannel:
				ticker.Stop()
				close(outputChannel)
				close(init)
				return
			case <-ticker.C:
			case <-init:
			}

			features, response, err := featuresApi.GetFeaturesExecute(featuresApi.GetFeatures(clientCtx))
			if err != nil {
				log.Errorf("Error retrieving features: %v", err)
				continue
			}

			if response.StatusCode != http.StatusOK {
				log.Errorf("Error retrieving features, statuscode was: %s", response.Status)
				continue
			}

			outputChannel <- features
		}
	}()

	return outputChannel, func() { close(stopChannel) }
}
