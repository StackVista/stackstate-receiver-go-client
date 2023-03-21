package transactionforwarder

import (
	"encoding/json"
	"fmt"
	"github.com/StackVista/stackstate-receiver-go-client/pkg/transactional"
	"github.com/StackVista/stackstate-receiver-go-client/pkg/transactional/transactionmanager"
	"github.com/fatih/color"
)

func NewPrintingTransactionalForwarder(manager transactionmanager.TransactionManager) *PrintingTransactionalForwarder {
	return &PrintingTransactionalForwarder{PayloadChan: make(chan TransactionalPayload, 100), manager: manager}
}

// PrintingTransactionalForwarder is a implementation of the transactional forwarder that prints the payload
type PrintingTransactionalForwarder struct {
	PayloadChan chan TransactionalPayload
	manager     transactionmanager.TransactionManager
}

// Start is a noop
func (mf *PrintingTransactionalForwarder) Start() {}

// SubmitTransactionalIntake receives a TransactionalPayload and keeps it in the PayloadChan to be used in assertions
func (mf *PrintingTransactionalForwarder) SubmitTransactionalIntake(payload TransactionalPayload) {

	// Acknowledge actions and succeed transactions
	for transactionID, payloadTransaction := range payload.TransactionActionMap {
		mf.manager.AcknowledgeAction(transactionID, payloadTransaction.ActionID)

		// if the transaction of the payload is completed, submit a transaction complete
		if payloadTransaction.CompletedTransaction {
			mf.manager.CompleteTransaction(transactionID)
		}
	}

	actualPayload := transactional.NewIntakePayload()
	_ = json.Unmarshal(payload.Body, &actualPayload)

	fmt.Fprintln(color.Output, fmt.Sprintf("=== %s ===", color.BlueString("Topology")))
	j, _ := json.MarshalIndent(actualPayload, "", "  ")
	fmt.Println(string(j))
}

// Stop is a noop
func (mf *PrintingTransactionalForwarder) Stop() {
	close(mf.PayloadChan)
}
