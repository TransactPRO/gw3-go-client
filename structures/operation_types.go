package structures

// OperationType describes the operation action as string
type OperationType string

const (
	/*
		Transaction Types
	*/

	// SMS transactions request type for url route path
	SMS OperationType = "sms"
	// DMSHold transactions request type for url route path
	DMSHold OperationType = "hold-dms"
	// DMSCharge transactions request type for url route path
	DMSCharge OperationType = "charge-dms"
	// CANCEL transactions request type for url route path
	CANCEL OperationType = "cancel"
	// MOTOSMS transactions request type for url route path
	MOTOSMS OperationType = "moto/sms"
	// MOTODMS transactions request type for url route path
	MOTODMS OperationType = "moto/dms"
	// RecurrentSMS transactions request type for url route path
	RecurrentSMS OperationType = "recurrent/sms"
	// RecurrentDMS transactions request type for url route path
	RecurrentDMS OperationType = "recurrent/dms"
	// Refund transactions request type for url route path
	Refund OperationType = "refund"
	// Reversal transactions request type for url route path
	Reversal OperationType = "reversal"

	/*
		Exploring Past Payments types
	*/

	// Status is a transaction status request type for url route path
	Status OperationType = "status"
	// Result is a transaction status request type for url route path
	Result OperationType = "result"
)
