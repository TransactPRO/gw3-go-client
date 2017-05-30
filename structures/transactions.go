package structures

// OperationType describes the operation action as string
type OperationType string

const (
	// SMS transaction type in route path
	SMS OperationType = "sms"
	// DMSHold transaction type in route path
	DMSHold OperationType = "hold-dms"
	// DMSCharge transaction type in route path
	DMSCharge OperationType = "charge-dms"
	// CANCEL transaction type in route path
	CANCEL OperationType = "cancel"
	// MOTOSMS transaction type in route path
	MOTOSMS OperationType = "moto/sms"
	// MOTOSMS transaction type in route path
	MOTODMS OperationType = "moto/dms"
)
