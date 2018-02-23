package transactions

import "github.com/TransactPRO/gw3-go-client/structures"

// RecurrentAssembly is default structure for recurrent transactions operation
type RecurrentAssembly struct {
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

// NewRecurrentSMSAssembly returns new instance with prepared HTTP request data RecurrentAssembly
func NewRecurrentSMSAssembly() *RecurrentAssembly {
	// Predefine default HTTP request data for sms operations
	var opd structures.OperationRequestHTTPData

	opd.SetHTTPMethod("POST")
	opd.SetOperationType(structures.RecurrentSMS)

	return &RecurrentAssembly{
		opHTTPData: opd,
	}
}

// NewRecurrentDMSAssembly returns new instance with prepared HTTP request data RecurrentAssembly
func NewRecurrentDMSAssembly() *RecurrentAssembly {
	// Predefine default HTTP request data for sms operations
	var opd structures.OperationRequestHTTPData

	opd.SetHTTPMethod("POST")
	opd.SetOperationType(structures.RecurrentDMS)

	return &RecurrentAssembly{
		opHTTPData: opd,
	}
}

// Implement methods for RecurrentAssembly structure, form pck structures OperationRequestInterface

// GetHTTPMethod return HTTP method which will be used for send request
func (op *RecurrentAssembly) GetHTTPMethod() string {
	return op.opHTTPData.GetHTTPMethod()
}

// GetOperationType return part of route path which will be used for send request
func (op *RecurrentAssembly) GetOperationType() structures.OperationType {
	return op.opHTTPData.GetOperationType()
}
