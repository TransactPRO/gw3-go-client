package operations

// OperationType describes the operation action as string
type OperationType string

const (
	// SMS transaction type
	SMS OperationType = "sms"

	// DMSHOLD transaction type
	DMSHOLD = "hold-dms"
)

// OperationBuilder operation structure builder for specific request
type OperationBuilder struct{}

// OperationInterface contains two methods, witch allows to get binned information about operation request
type OperationInterface interface {
	GetHTTPMethod() string
	GetOperationType() OperationType
}

// OperationRequestHTTPData contained with HTTP method (POST, GET) and action URL path for operation (sms, dms_hold)
type operationRequestHTTPData struct {
	methodHTTP    string
	operationType OperationType
}

// GetHTTPMethod return HTTP method which will be used for send request
func (op *operationRequestHTTPData) GetHTTPMethod() string {
	return op.methodHTTP
}

// GetOperationType return part of route path which will be used for send request
func (op *operationRequestHTTPData) GetOperationType() OperationType {
	return op.operationType
}
