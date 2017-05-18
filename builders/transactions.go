package builders

import "bitbucket.transactpro.lv/tls/gw3-go-client/structures"

type OperationType string

const (
	SMS OperationType 	= "sms"
	DMS_HOLD		= "dms-hold"
)


type OperationBuilder struct {}

type SMSPayload struct {
	PaymentMethod 	structures.PaymentMethodData		`json:"payment-method-data"`
	Money 		structures.MoneyData			`json:"money-data"`
	System 		structures.SystemData			`json:"system"`
}

// NewSMS, returns bundled structure of transaction SMS
func (ob *OperationBuilder) SMS() SMSPayload {
	return SMSPayload{}
}