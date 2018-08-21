package verify

import "github.com/TransactPRO/gw3-go-client/structures"

// ExploreTransactionAssembly is default structure for transaction status operation
type Verify3dEnrollmentAssembly struct {
	// HTTPData contains HTTP request method and operation action value for request in URL path
	opHTTPData structures.OperationRequestHTTPData
	// Request Data
	structures.Verify3dEnrollmentData
}

// NewStatusAssembly returns new instance with prepared HTTP request data ExploreTransactionAssembly
func NewVerify3dEnrollmentAssembly() *Verify3dEnrollmentAssembly {
	// Predefine default HTTP request data for sms operations
	var opd structures.OperationRequestHTTPData

	opd.SetHTTPMethod("POST")
	opd.SetOperationType(structures.Verify3dEnrollment)

	return &Verify3dEnrollmentAssembly{
		opHTTPData: opd,
	}
}

// Implement methods for ExploreAssembly structure, form pck structures OperationRequestInterface

// GetHTTPMethod return HTTP method which will be used for send request
func (op *Verify3dEnrollmentAssembly) GetHTTPMethod() string {
	return op.opHTTPData.GetHTTPMethod()
}

// GetOperationType return part of route path which will be used for send request
func (op *Verify3dEnrollmentAssembly) GetOperationType() structures.OperationType {
	return op.opHTTPData.GetOperationType()
}
