package transaction

import "bitbucket.transactpro.lv/tls/gw3-go-client/structures"

// CancelAssembly is default structure for Double-Message Transactions (DMS) Cancel transaction operation
type CancelAssembly struct {
	// HTTPData contains HTTP request method and operation action value for request in URL path
	opHTTPData structures.OperationRequestHTTPData
	// Command Data, isn't for any request type and in that case it's combined
	CommandData struct {
		structures.CommandDataGWTransactionID
	} `json:"command-data,omitempty"`
	System structures.SystemData `json:"system"`
}

// NewCancelAssembly returns new instance with prepared HTTP request data SMSAssembly
func NewCancelAssembly() *CancelAssembly {
	// Predefine default HTTP request data for sms operations
	var opd structures.OperationRequestHTTPData

	opd.SetHTTPMethod("POST")
	opd.SetOperationType(structures.CANCEL)

	return &CancelAssembly{
		opHTTPData: opd,
	}
}

// Implement methods for CancelDMSAssembly structure, form pck structures OperationRequestInterface

// GetHTTPMethod return HTTP method which will be used for send request
func (op *CancelAssembly) GetHTTPMethod() string {
	return op.opHTTPData.GetHTTPMethod()
}

// GetOperationType return part of route path which will be used for send request
func (op *CancelAssembly) GetOperationType() structures.OperationType {
	return op.opHTTPData.GetOperationType()
}
