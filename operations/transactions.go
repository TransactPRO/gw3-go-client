package operations

type Operation struct {}

type SMSPayload struct {
	Data struct {
		PaymentMethod 	PaymentMethodData		`json:"payment-method-data"`
		Money 		MoneyData			`json:"money-data"`
		System 		SystemData			`json:"system"`
	} `json:"data"`
}

// NewSMS, returns bundled structure of transaction SMS
func (*Operation) NewSMS() SMSPayload {
	return SMSPayload{}
}