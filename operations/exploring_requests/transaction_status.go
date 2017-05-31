package exploring_requests

import "bitbucket.transactpro.lv/tls/gw3-go-client/structures"

// StatusAssembly is default structure for transaction status operation
type StatusAssembly struct {
	// HTTPData contains HTTP request method and operation action value for request in URL path
	opHTTPData structures.OperationRequestHTTPData
	// Command Data, isn't for any request type and in that case it's combined
	CommandData struct {
		structures.CommandDataExploreGWTransactionIDs
	} `json:"command-data,omitempty"`
	System structures.SystemData `json:"system"`
}

// NewChargeDMSAssembly returns new instance with prepared HTTP request data ChargeDMSAssembly
func NewStatusAssembly() *StatusAssembly {
	// Predefine default HTTP request data for sms operations
	var opd structures.OperationRequestHTTPData

	opd.SetHTTPMethod("POST")
	opd.SetOperationType(structures.Status)

	return &StatusAssembly{
		opHTTPData: opd,
	}
}

// Implement methods for StatusAssembly structure, form pck structures OperationRequestInterface

// GetHTTPMethod return HTTP method which will be used for send request
func (op *StatusAssembly) GetHTTPMethod() string {
	return op.opHTTPData.GetHTTPMethod()
}

// GetOperationType return part of route path which will be used for send request
func (op *StatusAssembly) GetOperationType() structures.OperationType {
	return op.opHTTPData.GetOperationType()
}
