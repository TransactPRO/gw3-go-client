package transaction

import "bitbucket.transactpro.lv/tls/gw3-go-client/structures"

// SMSAssembly is default structure for sms transaction operation
type SMSAssembly struct {
	// HTTPData contains HTTP request method and operation action value for request in URL path
	opHTTPData structures.OperationRequestHTTPData
	// Command Data, isn't for any request type
	CommandData struct {
		// Inside form ID when selecting non-default form manually, allowed in sms, dms-hold
		FormID string `json:"form-id,omitempty"`
		// Terminal MID when selecting terminal manually, allowed in sms, dms-hold
		TerminalMID string `json:"terminal-mid,omitempty"`
	} `json:"command-data,omitempty"`

	GeneralData   structures.GeneralData       `json:"general-data,omitempty"`
	PaymentMethod structures.PaymentMethodData `json:"payment-method-data"`
	Money         structures.MoneyData         `json:"money-data"`
	System        structures.SystemData        `json:"system"`
}

// NewSMSAssembly returns new instance with prepared HTTP request data SMSAssembly
func NewSMSAssembly() *SMSAssembly {
	// Predefine default HTTP request data for sms operations
	var opd structures.OperationRequestHTTPData

	opd.SetHTTPMethod("POST")
	opd.SetOperationType(structures.SMS)

	return &SMSAssembly{
		opHTTPData: opd,
	}
}

// Implement methods for SMSAssembly structure, form pck structures OperationRequestInterface

// GetHTTPMethod return HTTP method which will be used for send request
func (op *SMSAssembly) GetHTTPMethod() string {
	return op.opHTTPData.GetHTTPMethod()
}

// GetOperationType return part of route path which will be used for send request
func (op *SMSAssembly) GetOperationType() structures.OperationType {
	return op.opHTTPData.GetOperationType()
}
