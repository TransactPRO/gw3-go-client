package verify

import "github.com/TransactPRO/gw3-go-client/structures"

// VerifyCardAssembly is default structure for card verification completion operation
type VerifyCardAssembly struct {
	// HTTPData contains HTTP request method and operation action value for request in URL path
	opHTTPData structures.OperationRequestHTTPData
	// Request Data
	structures.VerifyCardData
}

// NewVerifyCardAssembly returns new instance with prepared HTTP request data VerifyCardAssembly
func NewVerifyCardAssembly() *VerifyCardAssembly {
	// Predefine default HTTP request data for sms operations
	var opd structures.OperationRequestHTTPData

	opd.SetHTTPMethod("POST")
	opd.SetOperationType(structures.VerifyCard)

	return &VerifyCardAssembly{
		opHTTPData: opd,
	}
}

// Implement methods for ExploreAssembly structure, form pck structures OperationRequestInterface

// GetHTTPMethod return HTTP method which will be used for send request
func (op *VerifyCardAssembly) GetHTTPMethod() string {
	return op.opHTTPData.GetHTTPMethod()
}

// GetOperationType return part of route path which will be used for send request
func (op *VerifyCardAssembly) GetOperationType() structures.OperationType {
	return op.opHTTPData.GetOperationType()
}
