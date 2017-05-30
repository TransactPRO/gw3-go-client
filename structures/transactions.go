package structures

// OperationType describes the operation action as string
type OperationType string

const (
	// SMS transaction type in route path
	SMS OperationType = "sms"
	// DMSHold transaction type in route path
	DMSHold = "hold-dms"
	// DMSCharge transaction type in route path
	DMSCharge = "charge-dms"
	// CANCEL transaction type in route path
	CANCEL = "cancel"
)
