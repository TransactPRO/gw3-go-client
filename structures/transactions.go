package structures

// OperationType describes the operation action as string
type OperationType string

const (
	// SMS transaction type in route path
	SMS OperationType = "sms"
	// DMSHOLD transaction type in route path
	DMSHOLD = "hold-dms"
	// DMSCHARGE transaction type in route path
	DMSCHARGE = "charge-dms"
)
