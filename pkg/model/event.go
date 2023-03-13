// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2016-present Datadog, Inc.

package model

import (
	"encoding/json"
	"fmt"
)

// EventPriority represents the priority of an event
type EventPriority string

// Enumeration of the existing event priorities, and their values
const (
	EventPriorityNormal EventPriority = "normal"
	EventPriorityLow    EventPriority = "low"
)

// GetEventPriorityFromString returns the EventPriority from its string representation
func GetEventPriorityFromString(val string) (EventPriority, error) {
	switch val {
	case string(EventPriorityNormal):
		return EventPriorityNormal, nil
	case string(EventPriorityLow):
		return EventPriorityLow, nil
	default:
		return "", fmt.Errorf("Invalid event priority: '%s'", val)
	}
}

// EventAlertType represents the alert type of an event
type EventAlertType string

// Enumeration of the existing event alert types, and their values
const (
	EventAlertTypeError   EventAlertType = "error"
	EventAlertTypeWarning EventAlertType = "warning"
	EventAlertTypeInfo    EventAlertType = "info"
	EventAlertTypeSuccess EventAlertType = "success"
)

// GetAlertTypeFromString returns the EventAlertType from its string representation
func GetAlertTypeFromString(val string) (EventAlertType, error) {
	switch val {
	case string(EventAlertTypeError):
		return EventAlertTypeError, nil
	case string(EventAlertTypeWarning):
		return EventAlertTypeWarning, nil
	case string(EventAlertTypeInfo):
		return EventAlertTypeInfo, nil
	case string(EventAlertTypeSuccess):
		return EventAlertTypeSuccess, nil
	default:
		return EventAlertTypeInfo, fmt.Errorf("Invalid alert type: '%s'", val)
	}
}

// Event holds an event (w/ serialization to DD agent 5 intake format)
type Event struct {
	Title          string         `json:"msg_title"`
	Text           string         `json:"msg_text"`
	Ts             int64          `json:"timestamp"`
	Priority       EventPriority  `json:"priority,omitempty"`
	Host           string         `json:"host"`
	Tags           []string       `json:"tags,omitempty"`
	AlertType      EventAlertType `json:"alert_type,omitempty"`
	AggregationKey string         `json:"aggregation_key,omitempty"`
	SourceTypeName string         `json:"source_type_name,omitempty"`
	EventType      string         `json:"event_type,omitempty"`
	OriginID       string         `json:"-"`
	K8sOriginID    string         `json:"-"`
	Cardinality    string         `json:"-"`
	EventContext   *EventContext  `json:"context,omitempty"`
}

// EventContext enriches the event with some more context and allows correlation to topology in StackState
type EventContext struct {
	SourceIdentifier   string                 `json:"source_identifier,omitempty" msg:"source_identifier,omitempty"`
	ElementIdentifiers []string               `json:"element_identifiers" msg:"element_identifiers"`
	Source             string                 `json:"source" msg:"source"`
	Category           string                 `json:"category" msg:"category"`
	Data               map[string]interface{} `json:"data,omitempty" msg:"data,omitempty"`
	SourceLinks        []SourceLink           `json:"source_links" msg:"source_links"`
} // [sts]

// SourceLink points to links that may contain more information about this event
type SourceLink struct {
	Title string `json:"title" msg:"title"`
	URL   string `json:"url" msg:"url"`
} // [sts]

// Return a JSON string
func (e *Event) String() string {
	s, err := json.Marshal(e)
	if err != nil {
		fmt.Println(err)
		return fmt.Sprintf("{\"error\": \"%s\"}", err.Error())
	}
	return string(s)
} // [sts]

// IntakeEvents are used in the transactional batcher to keep state of events in a payload [sts]
type IntakeEvents struct {
	Events []Event
}

// IntakeFormat returns a map of events grouped by source type name
func (ie IntakeEvents) IntakeFormat() map[string][]Event {
	eventsBySourceType := make(map[string][]Event)
	for _, e := range ie.Events {
		sourceTypeName := e.SourceTypeName
		if sourceTypeName == "" {
			sourceTypeName = "api"
		}

		// ensure that event context lists are not empty. ie serialized to null
		if e.EventContext != nil {
			if e.EventContext.SourceLinks == nil {
				e.EventContext.SourceLinks = make([]SourceLink, 0)
			}

			if e.EventContext.ElementIdentifiers == nil {
				e.EventContext.ElementIdentifiers = make([]string, 0)
			}
		}

		eventsBySourceType[sourceTypeName] = append(eventsBySourceType[sourceTypeName], e)
	}
	return eventsBySourceType
}
