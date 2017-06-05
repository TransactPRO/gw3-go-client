package transactions

import "bitbucket.transactpro.lv/tls/gw3-go-client/structures"

// ReversalAssembly is default structure for reversal transactions operation
type ReversalAssembly struct {
	// HTTPData contains HTTP request method and operation action value for request in URL path
	opHTTPData structures.OperationRequestHTTPData
	// Command Data, isn't for any request type and in that case it's combined
	CommandData struct {
		structures.CommandDataGWTransactionID
	} `json:"command-data,omitempty"`
	GeneralData   structures.GeneralData       `json:"general-data,omitempty"`
	PaymentMethod structures.PaymentMethodData `json:"payment-method-data"`
	Money         structures.MoneyData         `json:"money-data"`
	// System data contains user(cardholder) IPv4 address and IPv4 address in case of proxy
	System structures.SystemData `json:"system"`
}

// NewReversalAssembly returns new instance with prepared HTTP request data RefundAssembly
func NewReversalAssembly() *ReversalAssembly {
	// Predefine default HTTP request data for sms operations
	var opd structures.OperationRequestHTTPData

	opd.SetHTTPMethod("POST")
	opd.SetOperationType(structures.Reversal)

	return &ReversalAssembly{
		opHTTPData: opd,
	}
}

// Implement methods for ReversalAssembly structure, form pck structures OperationRequestInterface

// GetHTTPMethod return HTTP method which will be used for send request
func (op *ReversalAssembly) GetHTTPMethod() string {
	return op.opHTTPData.GetHTTPMethod()
}

// GetOperationType return part of route path which will be used for send request
func (op *ReversalAssembly) GetOperationType() structures.OperationType {
	return op.opHTTPData.GetOperationType()
}
