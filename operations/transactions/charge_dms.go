package transactions

import "bitbucket.transactpro.lv/tls/gw3-go-client/structures"

// ChargeDMSAssembly is default structure for Double-Message Transactions (DMS) Hold transactions operation
type ChargeDMSAssembly struct {
	// HTTPData contains HTTP request method and operation action value for request in URL path
	opHTTPData structures.OperationRequestHTTPData
	// Command Data, isn't for any request type and in that case it's combined
	CommandData struct {
		structures.CommandDataGWTransactionID
	} `json:"command-data,omitempty"`
	Money  structures.MoneyData  `json:"money-data"`
	// System data contains user(cardholder) IPv4 address and IPv4 address in case of proxy
	System structures.SystemData `json:"system"`
}

// NewChargeDMSAssembly returns new instance with prepared HTTP request data ChargeDMSAssembly
func NewChargeDMSAssembly() *ChargeDMSAssembly {
	// Predefine default HTTP request data for sms operations
	var opd structures.OperationRequestHTTPData

	opd.SetHTTPMethod("POST")
	opd.SetOperationType(structures.DMSCharge)

	return &ChargeDMSAssembly{
		opHTTPData: opd,
	}
}

// Implement methods for ChargeDMSAssembly structure, form pck structures OperationRequestInterface

// GetHTTPMethod return HTTP method which will be used for send request
func (op *ChargeDMSAssembly) GetHTTPMethod() string {
	return op.opHTTPData.GetHTTPMethod()
}

// GetOperationType return part of route path which will be used for send request
func (op *ChargeDMSAssembly) GetOperationType() structures.OperationType {
	return op.opHTTPData.GetOperationType()
}
