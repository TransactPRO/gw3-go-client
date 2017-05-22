package builder

import "bitbucket.transactpro.lv/tls/gw3-go-client/structures"

type OperationType string

const (
	SMS OperationType 	= "sms"
	DMS_HOLD		= "hold-dms"
)

type OperationBuilder struct {}

// Combined general data structure
type generalData struct{
	CustomerData 	structures.CustomerData		`json:"customer-data,omitempty"`
	OrderData	structures.OrderData 		`json:"order-data,omitempty"`
}

// SMS bundle
type SMSDataSet struct {
	GeneralData 	generalData			`json:"general-data,omitempty"`
	PaymentMethod 	structures.PaymentMethodData	`json:"payment-method-data"`
	Money 		structures.MoneyData		`json:"money-data"`
	System 		structures.SystemData		`json:"system"`
}

// NewSMS, returns bundled structure of transaction SMS
func (ob *OperationBuilder) SMS() SMSDataSet {
	return SMSDataSet{}
}