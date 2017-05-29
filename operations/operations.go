package operations

import (
	"bitbucket.transactpro.lv/tls/gw3-go-client/structures"
)

// OperationBuilder operation structure builder for specific request
type OperationBuilder struct{}

// OperationInterface contains two methods, witch allows to get binned information about operation request
type OperationInterface interface {
	GetHTTPMethod() string
	GetOperationType() structures.OperationType
}

// OperationRequestHTTPData contained with HTTP method (POST, GET) and action URL path for operation (sms, dms_hold)
type operationRequestHTTPData struct {
	methodHTTP    string
	operationType structures.OperationType
}

// GetHTTPMethod return HTTP method which will be used for send request
func (op *operationRequestHTTPData) GetHTTPMethod() string {
	return op.methodHTTP
}

// GetOperationType return part of route path which will be used for send request
func (op *operationRequestHTTPData) GetOperationType() structures.OperationType {
	return op.operationType
}
