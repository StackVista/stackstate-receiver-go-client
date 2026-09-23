package features

import "time"

// Class is a bounded, credential-safe query outcome.
type Class string

// Query outcome classes distinguish capability observations from failures.
const (
	Valid          Class = "valid"
	Unsupported    Class = "unsupported"
	Authentication Class = "authentication"
	Transient      Class = "transient"
	Timeout        Class = "timeout"
	Malformed      Class = "malformed"
	Configuration  Class = "configuration"
	Rejected       Class = "rejected"
	Canceled       Class = "canceled"
)

// Result describes a completed query without retaining raw errors or response bodies.
type Result struct {
	Class      Class
	Features   map[string]any
	StatusCode int
	Attempts   int
	FinishedAt time.Time
}
