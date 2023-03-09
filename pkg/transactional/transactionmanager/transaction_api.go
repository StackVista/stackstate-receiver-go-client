package transactionmanager

import "github.com/StackVista/stackstate-receiver-go-client/pkg/model"

// TransactionAPI contains the functions required for transactional behaviour
type TransactionAPI interface {
	GetActiveTransaction(transactionID string) (*IntakeTransaction, error)
	TransactionCount() int
	StartTransaction(CheckID model.CheckID, TransactionID string, NotifyChannel chan interface{})
	CompleteTransaction(transactionID string)
	DiscardTransaction(transactionID, reason string)
	CommitAction(transactionID, actionID string)
	AcknowledgeAction(transactionID, actionID string)
	SetState(transactionID, key string, state string)
	RejectAction(transactionID, actionID, reason string)
}
