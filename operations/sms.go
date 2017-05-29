package operations

import "bitbucket.transactpro.lv/tls/gw3-go-client/structures"

// SMS is default structure for sms transaction operation
type SMS struct {
	// HTTPData contains HTTP request method and operation action value for request in URL path
	reqHTTPData operationRequestHTTPData
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

// NewSms returns new pointer to new SMS structure
func (ob *OperationBuilder) NewSms() *SMS {
	// Predefine default HTTP request data for sms operations
	var opHTTPData operationRequestHTTPData
	opHTTPData.methodHTTP = "POST"
	opHTTPData.operationType = structures.SMS

	return &SMS{
		reqHTTPData: opHTTPData,
	}
}

// GetHTTPMethod gets HTTP method for these operation
func (sms *SMS) GetHTTPMethod() string {
	return sms.reqHTTPData.GetHTTPMethod()
}

// GetHTTPMethod gets action URL path for these operation
func (sms *SMS) GetOperationType() structures.OperationType {
	return sms.reqHTTPData.GetOperationType()
}
