package verify

import (
	"net/http"

	"github.com/TransactPRO/gw3-go-client/structures"
)

// CardAssembly is default structure for card verification completion operation
type CardAssembly struct {
	// HTTPData contains HTTP request method and operation action value for request in URL path
	opHTTPData structures.OperationRequestHTTPData
	// Request Data
	structures.VerifyCardData
}

// NewVerifyCardAssembly returns new instance with prepared HTTP request data CardAssembly
func NewVerifyCardAssembly() *CardAssembly {
	// Predefine default HTTP request data for sms operations
	var opd structures.OperationRequestHTTPData

	opd.SetHTTPMethod(http.MethodPost)
	opd.SetOperationType(structures.VerifyCard)

	return &CardAssembly{
		opHTTPData: opd,
	}
}

// Implement methods for ExploreAssembly structure, form pck structures OperationRequestInterface

// GetHTTPMethod return HTTP method which will be used for send request
func (op *CardAssembly) GetHTTPMethod() string {
	return op.opHTTPData.GetHTTPMethod()
}

// GetOperationType return part of route path which will be used for send request
func (op *CardAssembly) GetOperationType() structures.OperationType {
	return op.opHTTPData.GetOperationType()
}
