package builder

import (
	"bitbucket.transactpro.lv/tls/gw3-go-client/structures"
)

// OperationBuilder operation structure builder for specific request
type OperationBuilder struct{}

type (
	// SMSAssembly is structure for sms transaction type
	SMSAssembly struct {
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
)

// SMS method returns bundled structure for SMS transaction request
//func (ob *OperationBuilder) SMS() *SMSAssembly {
//	// Prepare new operation container
//	// Set default settings for that operation
//	ob.opContainer.RequestHTTPData.Method = "POST"
//	ob.opContainer.RequestHTTPData.OperationType = structures.SMS
//
//	ob.opContainer.Data = &SMSAssembly{}
//	fmt.Println(fmt.Sprintf("%f", &ob.opContainer.Data))
//
//	return ob.opContainer.Data.(&SMSAssembly)
//}
