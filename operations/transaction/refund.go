package transaction

import "bitbucket.transactpro.lv/tls/gw3-go-client/structures"

// RefundAssembly is default structure for Refunds transaction operation
type RefundAssembly struct {
	// HTTPData contains HTTP request method and operation action value for request in URL path
	opHTTPData structures.OperationRequestHTTPData
	// Command Data, isn't for any request type and in that case it's combined
	CommandData struct {
		structures.CommandDataGWTransactionID
	} `json:"command-data,omitempty"`
	GeneralData   structures.GeneralData       `json:"general-data,omitempty"`
	PaymentMethod structures.PaymentMethodData `json:"payment-method-data"`
	Money         structures.MoneyData         `json:"money-data"`
	System        structures.SystemData        `json:"system"`
}

// NewRefundAssembly returns new instance with prepared HTTP request data RefundAssembly
func NewRefundAssembly() *RefundAssembly {
	// Predefine default HTTP request data for sms operations
	var opd structures.OperationRequestHTTPData

	opd.SetHTTPMethod("POST")
	opd.SetOperationType(structures.Refund)

	return &RefundAssembly{
		opHTTPData: opd,
	}
}

// Implement methods for RecurrentAssembly structure, form pck structures OperationRequestInterface

// GetHTTPMethod return HTTP method which will be used for send request
func (op *RefundAssembly) GetHTTPMethod() string {
	return op.opHTTPData.GetHTTPMethod()
}

// GetOperationType return part of route path which will be used for send request
func (op *RefundAssembly) GetOperationType() structures.OperationType {
	return op.opHTTPData.GetOperationType()
}
