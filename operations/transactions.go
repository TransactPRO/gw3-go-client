package operations

type OperationType string

const (
	SMS OperationType 	= "sms"
	DMS_HOLD		= "dms-hold"
)


type OperationBuilder struct {}

type SMSPayload struct {
	PaymentMethod 	PaymentMethodData		`json:"payment-method-data"`
	Money 		MoneyData			`json:"money-data"`
	System 		SystemData			`json:"system"`
}

// NewSMS, returns bundled structure of transaction SMS
func (ob *OperationBuilder) SMS() SMSPayload {
	return SMSPayload{}
}