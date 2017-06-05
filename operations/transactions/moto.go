package transactions

import "bitbucket.transactpro.lv/tls/gw3-go-client/structures"

// MOTOAssembly is default structure for Offline Transactions (MOTO) transactions operation
type MOTOAssembly struct {
	// HTTPData contains HTTP request method and operation action value for request in URL path
	opHTTPData structures.OperationRequestHTTPData
	// Command Data, isn't for any request type and in that case it's combined
	CommandData struct {
		structures.CommandDataTerminalMID
	} `json:"command-data,omitempty"`
	GeneralData   structures.GeneralData       `json:"general-data,omitempty"`
	PaymentMethod structures.PaymentMethodData `json:"payment-method-data"`
	Money         structures.MoneyData         `json:"money-data"`
	// System data contains user(cardholder) IPv4 address and IPv4 address in case of proxy
	System structures.SystemData `json:"system"`
}

// NewMOTOSMSAssembly returns new instance with prepared HTTP request data MOTOAssembly
func NewMOTOSMSAssembly() *MOTOAssembly {
	// Predefine default HTTP request data for sms operations
	var opd structures.OperationRequestHTTPData

	opd.SetHTTPMethod("POST")
	opd.SetOperationType(structures.MOTOSMS)

	return &MOTOAssembly{
		opHTTPData: opd,
	}
}

// NewMOTODMSAssembly returns new instance with prepared HTTP request data MOTOAssembly
func NewMOTODMSAssembly() *MOTOAssembly {
	// Predefine default HTTP request data for sms operations
	var opd structures.OperationRequestHTTPData

	opd.SetHTTPMethod("POST")
	opd.SetOperationType(structures.MOTODMS)

	return &MOTOAssembly{
		opHTTPData: opd,
	}
}

// Implement methods for MOTOAssembly structure, form pck structures OperationRequestInterface

// GetHTTPMethod return HTTP method which will be used for send request
func (op *MOTOAssembly) GetHTTPMethod() string {
	return op.opHTTPData.GetHTTPMethod()
}

// GetOperationType return part of route path which will be used for send request
func (op *MOTOAssembly) GetOperationType() structures.OperationType {
	return op.opHTTPData.GetOperationType()
}
