package structures

// OperationType describes the operation action as string
type OperationType string

const (
	// SMS transaction type
	SMS OperationType = "sms"

	// DMSHOLD transaction type
	DMSHOLD = "hold-dms"
)
