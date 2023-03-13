package transactionbatcher

import (
	"github.com/StackVista/stackstate-receiver-go-client/pkg/model/check"
	"github.com/StackVista/stackstate-receiver-go-client/pkg/model/health"
	"github.com/StackVista/stackstate-receiver-go-client/pkg/model/telemetry"
	"github.com/StackVista/stackstate-receiver-go-client/pkg/model/topology"
	log "github.com/cihub/seelog"
)

// TransactionalBatcher interface can receive data for sending to the intake and will accumulate the data in batches. This does
// not work on a fixed schedule like the aggregator but flushes either when data exceeds a threshold, when
// data is complete.
type TransactionalBatcher interface {
	// Topology
	SubmitComponent(checkID check.CheckID, transactionID string, instance topology.Instance, component topology.Component)
	SubmitRelation(checkID check.CheckID, transactionID string, instance topology.Instance, relation topology.Relation)
	SubmitStartSnapshot(checkID check.CheckID, transactionID string, instance topology.Instance)
	SubmitStopSnapshot(checkID check.CheckID, transactionID string, instance topology.Instance)
	SubmitDelete(checkID check.CheckID, transactionID string, instance topology.Instance, topologyElementID string)

	// Health
	SubmitHealthCheckData(checkID check.CheckID, transactionID string, stream health.Stream, data health.CheckData)
	SubmitHealthStartSnapshot(checkID check.CheckID, transactionID string, stream health.Stream, intervalSeconds int, expirySeconds int)
	SubmitHealthStopSnapshot(checkID check.CheckID, transactionID string, stream health.Stream)

	// Raw Metrics
	SubmitRawMetricsData(checkID check.CheckID, transactionID string, data telemetry.RawMetric)

	// Events
	SubmitEvent(checkID check.CheckID, transactionID string, event telemetry.Event)

	// Transactional
	StartTransaction(checkID check.CheckID, transactionID string)
	SubmitCompleteTransaction(checkID check.CheckID, transactionID string)

	// lifecycle
	SubmitClearState(checkID check.CheckID)
	Stop()
}

// SubmitComponent is used to submit a component to the input channel
type SubmitComponent struct {
	CheckID       check.CheckID
	TransactionID string
	Instance      topology.Instance
	Component     topology.Component
}

// SubmitRelation is used to submit a relation to the input channel
type SubmitRelation struct {
	CheckID       check.CheckID
	TransactionID string
	Instance      topology.Instance
	Relation      topology.Relation
}

// SubmitStartSnapshot is used to submit a start of a snapshot to the input channel
type SubmitStartSnapshot struct {
	CheckID       check.CheckID
	TransactionID string
	Instance      topology.Instance
}

// SubmitStopSnapshot is used to submit a stop of a snapshot to the input channel
type SubmitStopSnapshot struct {
	CheckID       check.CheckID
	TransactionID string
	Instance      topology.Instance
}

// SubmitHealthCheckData is used to submit health check data to the input channel
type SubmitHealthCheckData struct {
	CheckID       check.CheckID
	TransactionID string
	Stream        health.Stream
	Data          health.CheckData
}

// SubmitHealthStartSnapshot is used to submit health check start snapshot to the input channel
type SubmitHealthStartSnapshot struct {
	CheckID         check.CheckID
	TransactionID   string
	Stream          health.Stream
	IntervalSeconds int
	ExpirySeconds   int
}

// SubmitHealthStopSnapshot is used to submit health check stop snapshot to the input channel
type SubmitHealthStopSnapshot struct {
	CheckID       check.CheckID
	TransactionID string
	Stream        health.Stream
}

// SubmitDelete is used to submit a topology delete to the input channel
type SubmitDelete struct {
	CheckID       check.CheckID
	TransactionID string
	Instance      topology.Instance
	DeleteID      string
}

// SubmitRawMetricsData is used to submit a raw metric value to the input channel
type SubmitRawMetricsData struct {
	CheckID       check.CheckID
	TransactionID string
	RawMetric     telemetry.RawMetric
}

// SubmitEvent is used to submit a event to the input channel
type SubmitEvent struct {
	CheckID       check.CheckID
	TransactionID string
	Event         telemetry.Event
}

// SubmitClearState is used to clear batcher state for a given CheckID
type SubmitClearState struct {
	CheckID check.CheckID
}

// StartTransaction is used to submit a start transaction to the input channel
type StartTransaction struct {
	CheckID       check.CheckID
	TransactionID string
}

// SubmitCompleteTransaction is used to submit a transaction complete to the input channel
type SubmitCompleteTransaction struct {
	CheckID       check.CheckID
	TransactionID string
}

// SubmitShutdown is used to submit a shutdown of the transactionbatcher to the input channel
type SubmitShutdown struct{}

// SubmitComponent submits a component to the batch
func (ctb *transactionalBatcher) SubmitComponent(checkID check.CheckID, transactionID string, instance topology.Instance, component topology.Component) {
	ctb.Input <- SubmitComponent{
		CheckID:       checkID,
		TransactionID: transactionID,
		Instance:      instance,
		Component:     component,
	}
}

// SubmitRelation submits a relation to the batch
func (ctb *transactionalBatcher) SubmitRelation(checkID check.CheckID, transactionID string, instance topology.Instance, relation topology.Relation) {
	ctb.Input <- SubmitRelation{
		CheckID:       checkID,
		TransactionID: transactionID,
		Instance:      instance,
		Relation:      relation,
	}
}

// SubmitStartSnapshot submits start of a snapshot
func (ctb *transactionalBatcher) SubmitStartSnapshot(checkID check.CheckID, transactionID string, instance topology.Instance) {
	ctb.Input <- SubmitStartSnapshot{
		CheckID:       checkID,
		TransactionID: transactionID,
		Instance:      instance,
	}
}

// SubmitStopSnapshot submits a stop of a snapshot. This always causes a flush of the data downstream
func (ctb *transactionalBatcher) SubmitStopSnapshot(checkID check.CheckID, transactionID string, instance topology.Instance) {
	ctb.Input <- SubmitStopSnapshot{
		CheckID:       checkID,
		TransactionID: transactionID,
		Instance:      instance,
	}
}

// SubmitDelete submits a deletion of topology element.
func (ctb *transactionalBatcher) SubmitDelete(checkID check.CheckID, transactionID string, instance topology.Instance, topologyElementID string) {
	ctb.Input <- SubmitDelete{
		CheckID:       checkID,
		TransactionID: transactionID,
		Instance:      instance,
		DeleteID:      topologyElementID,
	}
}

// SubmitHealthCheckData submits a Health check data record to the batch
func (ctb *transactionalBatcher) SubmitHealthCheckData(checkID check.CheckID, transactionID string, stream health.Stream, data health.CheckData) {
	log.Debugf("Submitting Health check data for check [%s] stream [%s]: %s", checkID, stream.GoString(), data.JSONString())
	ctb.Input <- SubmitHealthCheckData{
		CheckID:       checkID,
		TransactionID: transactionID,
		Stream:        stream,
		Data:          data,
	}
}

// SubmitHealthStartSnapshot submits start of a Health snapshot
func (ctb *transactionalBatcher) SubmitHealthStartSnapshot(checkID check.CheckID, transactionID string, stream health.Stream, intervalSeconds int, expirySeconds int) {
	ctb.Input <- SubmitHealthStartSnapshot{
		CheckID:         checkID,
		TransactionID:   transactionID,
		Stream:          stream,
		IntervalSeconds: intervalSeconds,
		ExpirySeconds:   expirySeconds,
	}
}

// SubmitHealthStopSnapshot submits a stop of a Health snapshot. This always causes a flush of the data downstream
func (ctb *transactionalBatcher) SubmitHealthStopSnapshot(checkID check.CheckID, transactionID string, stream health.Stream) {
	ctb.Input <- SubmitHealthStopSnapshot{
		CheckID:       checkID,
		TransactionID: transactionID,
		Stream:        stream,
	}
}

// SubmitRawMetricsData submits a raw metrics data record to the batch
func (ctb *transactionalBatcher) SubmitRawMetricsData(checkID check.CheckID, transactionID string, rawMetric telemetry.RawMetric) {
	if rawMetric.HostName == "" {
		rawMetric.HostName = ctb.Hostname
	}

	ctb.Input <- SubmitRawMetricsData{
		CheckID:       checkID,
		TransactionID: transactionID,
		RawMetric:     rawMetric,
	}
}

// SubmitEvent submits an event to the batch
func (ctb *transactionalBatcher) SubmitEvent(checkID check.CheckID, transactionID string, event telemetry.Event) {
	ctb.Input <- SubmitEvent{
		CheckID:       checkID,
		TransactionID: transactionID,
		Event:         event,
	}
}

// StartTransaction submits a start transaction
func (ctb *transactionalBatcher) StartTransaction(checkID check.CheckID, transactionID string) {
	ctb.Input <- StartTransaction{
		CheckID:       checkID,
		TransactionID: transactionID,
	}
}

// SubmitCompleteTransaction submits a complete of a transaction
func (ctb *transactionalBatcher) SubmitCompleteTransaction(checkID check.CheckID, transactionID string) {
	ctb.Input <- SubmitCompleteTransaction{
		CheckID:       checkID,
		TransactionID: transactionID,
	}
}

// SubmitClearState signals completion of a check. May trigger a flush only if the check produced data
func (ctb *transactionalBatcher) SubmitClearState(checkID check.CheckID) {
	log.Debugf("Submitting clear state for check [%s]", checkID)
	ctb.Input <- SubmitClearState{
		CheckID: checkID,
	}
}
