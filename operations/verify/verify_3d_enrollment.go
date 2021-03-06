package verify

import (
	"net/http"

	"github.com/TransactPRO/gw3-go-client/structures"
)

// ThreeDEnrollmentAssembly is default structure for card 3-D secure enrollment verification
type ThreeDEnrollmentAssembly struct {
	// HTTPData contains HTTP request method and operation action value for request in URL path
	opHTTPData structures.OperationRequestHTTPData
	// Request Data
	structures.Verify3dEnrollmentData
}

// NewVerify3dEnrollmentAssembly returns new instance with prepared HTTP request data ThreeDEnrollmentAssembly
func NewVerify3dEnrollmentAssembly() *ThreeDEnrollmentAssembly {
	// Predefine default HTTP request data for sms operations
	var opd structures.OperationRequestHTTPData

	opd.SetHTTPMethod(http.MethodPost)
	opd.SetOperationType(structures.Verify3dEnrollment)

	return &ThreeDEnrollmentAssembly{
		opHTTPData: opd,
	}
}

// Implement methods for ExploreAssembly structure, form pck structures OperationRequestInterface

// GetHTTPMethod return HTTP method which will be used for send request
func (op *ThreeDEnrollmentAssembly) GetHTTPMethod() string {
	return op.opHTTPData.GetHTTPMethod()
}

// GetOperationType return part of route path which will be used for send request
func (op *ThreeDEnrollmentAssembly) GetOperationType() structures.OperationType {
	return op.opHTTPData.GetOperationType()
}

// ParseResponse parses Gateway response into corresponding data structure
func (op *ThreeDEnrollmentAssembly) ParseResponse(response *structures.GatewayResponse) (result *structures.EnrollmentResponse, err error) {
	result = new(structures.EnrollmentResponse)
	err = response.ParseJSON(&result)
	return
}
