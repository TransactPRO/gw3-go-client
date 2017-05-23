package builder

import "bitbucket.transactpro.lv/tls/gw3-go-client/structures"

type OperationType string

const (
	SMS OperationType 	= "sms"
	DMS_HOLD		= "hold-dms"
)

// OperationBuilder, operation structure builder for specific request
//
// Allowed methods: SMS
type OperationBuilder struct {}

type (
	// Combined general data structure
	generalData struct{
		CustomerData 	structures.CustomerData		`json:"customer-data,omitempty"`
		OrderData	structures.OrderData 		`json:"order-data,omitempty"`
	}

	// SMS data bundle
	SMSAssembly struct {
		// Command Data, isn't for any request type
		CommandData 	struct {
			// Inside form ID when selecting non-default form manually, allowed in sms, dms-hold
			FormID 		string		`json:"form-id,omitempty"`
			// Terminal MID when selecting terminal manually, allowed in sms, dms-hold
			TerminalMID	string		`json:"terminal-mid,omitempty"`
		} `json:"command-data,omitempty"`
		GeneralData 	generalData			`json:"general-data,omitempty"`
		PaymentMethod 	structures.PaymentMethodData	`json:"payment-method-data"`
		Money 		structures.MoneyData		`json:"money-data"`
		System 		structures.SystemData		`json:"system"`
	}
)

// SMS, method returns bundled structure of SMS transaction request
func (ob *OperationBuilder) SMS() SMSAssembly {
	return SMSAssembly{}
}