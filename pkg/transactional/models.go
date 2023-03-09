package transactional

import (
	"encoding/json"
	"fmt"
	"reflect"
)

const (
	// IntakePath is the path for the intake route on the receiver API
	IntakePath = "intake"
)

// PayloadTransaction is used to keep track of a given actionID and completion status of a transaction when submitting
// payloads
type PayloadTransaction struct {
	ActionID             string
	CompletedTransaction bool
}

// IntakePayload is a Go representation of the Receiver Intake structure
type IntakePayload struct {
	InternalHostname string                     `json:"internalHostname"`
	Topologies       []model.Topology           `json:"topologies"`
	Health           []model.Health             `json:"health"`
	Metrics          []interface{}              `json:"metrics"`
	Events           map[string][]metrics.Event `json:"events"`
}

// JSONString returns a JSON string of the Component
func (ip *IntakePayload) JSONString() string {
	b, err := json.Marshal(ip)
	if err != nil {
		fmt.Println(err)
		return fmt.Sprintf("{\"error\": \"%s\"}", err.Error())
	}
	return string(b)
}

// EqualDataPayload compares the topology, health and metrics of two IntakePayloads and returns a bool indicating
// whether the intake payloads are equal
func (ip *IntakePayload) EqualDataPayload(ip2 IntakePayload) bool {
	return reflect.DeepEqual(ip.Topologies, ip2.Topologies) &&
		reflect.DeepEqual(ip.Health, ip2.Health) &&
		reflect.DeepEqual(ip.Metrics, ip2.Metrics) &&
		reflect.DeepEqual(ip.Events, ip2.Events)
}

// NewIntakePayload returns a IntakePayload with default values
func NewIntakePayload() IntakePayload {
	return IntakePayload{
		Topologies: make([]model.Topology, 0),
		Health:     make([]model.Health, 0),
		Metrics:    make([]interface{}, 0),
		Events:     make(map[string][]metrics.Event, 0),
	}
}
