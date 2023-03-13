// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2016-2020 Datadog, Inc.

package telemetry

import (
	"encoding/json"
	"fmt"
	"time"
)

// Metrics is a container structure for a list of raw metric values. This allows us to set Metrics of a batch payload as
// a pointer and append more metrics to the structure
type Metrics struct {
	Values []RawMetric
}

// RawMetric single payload structure
type RawMetric struct {
	Name      string   `json:"name,omitempty"`
	Timestamp int64    `json:"timestamp,omitempty"`
	HostName  string   `json:"host_name,omitempty"`
	Value     float64  `json:"value,omitempty"`
	Tags      []string `json:"tags,omitempty"`
}

// RawMetricsMetaData payload containing meta data for the metric
type RawMetricsMetaData struct {
	Hostname string   `json:"hostname,omitempty"`
	Tags     []string `json:"tags,omitempty"`
	Type     string   `json:"type,omitempty"`
}

// MakeRawMetric helper function for creating a RawMetric value
func MakeRawMetric(name string, hostName string, value float64, tags []string) RawMetric {
	timestamp := time.Now().Unix()
	return RawMetric{
		Name:      name,
		Timestamp: timestamp,
		HostName:  hostName,
		Value:     value,
		Tags:      tags,
	}
}

// ConvertToIntakeMetric Converts RawMetricsCheckData struct to an older v1 metrics structure
func (r RawMetric) ConvertToIntakeMetric() []interface{} {
	data := []interface{}{
		r.Name,
		r.Timestamp,
		r.Value,
		RawMetricsMetaData{
			Hostname: r.HostName,
			Type:     "raw",
			Tags:     r.Tags,
		},
	}
	return data
}

// IntakeMetricJSON Converts RawMetricsCheckData struct to an older v1 metrics structure, parses it to JSON and returns
// it as a interface. This is only used in batcher test assertions.
func (r RawMetric) IntakeMetricJSON() (jsonObject []interface{}) {
	jsonString, _ := json.Marshal(r.ConvertToIntakeMetric())
	_ = json.Unmarshal(jsonString, &jsonObject)
	return jsonObject
}

// JSONString returns a JSON string of the Component
func (r RawMetric) JSONString() string {
	b, err := json.Marshal(r)
	if err != nil {
		fmt.Println(err)
		return fmt.Sprintf("{\"error\": \"%s\"}", err.Error())
	}
	return string(b)
}
