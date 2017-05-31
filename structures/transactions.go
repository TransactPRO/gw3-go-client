package structures

// OperationType describes the operation action as string
type OperationType string

const (
	// SMS transactions type in route path
	SMS OperationType = "sms"
	// DMSHold transactions type in route path
	DMSHold OperationType = "hold-dms"
	// DMSCharge transactions type in route path
	DMSCharge OperationType = "charge-dms"
	// CANCEL transactions type in route path
	CANCEL OperationType = "cancel"
	// MOTOSMS transactions type in route path
	MOTOSMS OperationType = "moto/sms"
	// MOTOSMS transactions type in route path
	MOTODMS OperationType = "moto/dms"
	// RecurrentSMS transactions type in route path
	RecurrentSMS OperationType = "recurrent/sms"
	// RecurrentDMS transactions type in route path
	RecurrentDMS OperationType = "recurrent/dms"
	// Refund transactions type in route path
	Refund OperationType = "refund"
	// Reversal transactions type in route path
	Reversal OperationType = "reversal"
)
