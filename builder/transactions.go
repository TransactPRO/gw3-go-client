package builder

import "bitbucket.transactpro.lv/tls/gw3-go-client/structures"

// OperationType describes the operation action as string
type OperationType string

const (
	// SMS transaction type
	SMS OperationType = "sms"
	// DMSHOLD transaction type
	DMSHOLD = "hold-dms"
)

// OperationBuilder operation structure builder for specific request
type OperationBuilder struct{}

type (
	requestHTTPData struct {
		// HTTP method
		method string
		// URL path part
		routePath string
	}

	// SMSAssembly is structure for sms transaction type
	SMSAssembly struct {
		requestHTTPData
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

	// DMSHoldAssembly is structure for dms hold transaction type
	DMSHoldAssembly struct {
	}
)

// GetRoutePath return part of route path which will be used for send request
func (rd *requestHTTPData) GetRoutePath() string {
	return rd.routePath
}

// GetHTTPMethod return HTTP method which will be used for send request
func (rd *requestHTTPData) GetHTTPMethod() string {
	return rd.method
}

// SMS method returns bundled structure for SMS transaction request
func (ob *OperationBuilder) SMS() SMSAssembly {
	return SMSAssembly{}
}

// DMSHold method returns bundled structure for DMS HOLD transaction request
func (ob *OperationBuilder) DMSHold() DMSHoldAssembly {
	return DMSHoldAssembly{}
}
