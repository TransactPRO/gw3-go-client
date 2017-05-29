package operations

import "bitbucket.transactpro.lv/tls/gw3-go-client/structures"

// OperationBuilder operation structure builder for specific request
type OperationBuilder struct {}

type operationRequestHTTPData struct {
	methodHTTP    string
	operationType structures.OperationType
}

// operationInterface contains
type OperationInterface interface {
	GetHTTPMethod() string
	GetOperationType() structures.OperationType
}


// GetHTTPMethod return HTTP method which will be used for send request
func (op *operationRequestHTTPData) GetHTTPMethod() string {
	return op.methodHTTP
}

// GetOperationType return part of route path which will be used for send request
func (op *operationRequestHTTPData) GetOperationType() structures.OperationType {
	return op.operationType
}

// SMSAssembly is structure for sms transaction type
type SMS struct {
	opHTTPData operationRequestHTTPData
	GeneralData structures.GeneralData
	// Command Data, isn't for any request type
	CommandData struct {
		// Inside form ID when selecting non-default form manually, allowed in sms, dms-hold
		FormID string `json:"form-id,omitempty"`
		// Terminal MID when selecting terminal manually, allowed in sms, dms-hold
		TerminalMID string `json:"terminal-mid,omitempty"`
	} `json:"command-data,omitempty"`
	//GeneralData   structures.GeneralData       `json:"general-data,omitempty"`
	PaymentMethod structures.PaymentMethodData `json:"payment-method-data"`
	Money         structures.MoneyData         `json:"money-data"`
	System        structures.SystemData        `json:"system"`
}

func (ob *OperationBuilder) NewSms() *SMS {
	var httpData operationRequestHTTPData
	httpData.methodHTTP = "POST"
	httpData.operationType = structures.SMS

	return &SMS{
		opHTTPData: httpData,
	}
}

func (sms *SMS) GetHTTPMethod() string {
	return sms.opHTTPData.GetHTTPMethod()
}

func (sms *SMS) GetOperationType() structures.OperationType {
	return sms.opHTTPData.GetOperationType()
}