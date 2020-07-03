package helpers

import (
	"errors"
	"net/http"
	"net/url"

	"github.com/TransactPRO/gw3-go-client/structures"
)

// RetrieveFormAssembly is default structure for retrieving an HTML form from Gateway for a Cardholder
type RetrieveFormAssembly struct {
	opHTTPData structures.OperationRequestHTTPData
}

// NewRetrieveFormAssembly returns new instance with prepared HTTP request data RetrieveFormAssembly
func NewRetrieveFormAssembly(paymentResponse *structures.TransactionResponse) (result *RetrieveFormAssembly, err error) {
	if paymentResponse == nil || paymentResponse.Gateway.RedirectURL == nil {
		err = errors.New("response doesn't contain link to an HTML form")
		return
	}

	var opd structures.OperationRequestHTTPData
	opd.SetHTTPMethod(http.MethodGet)
	opd.SetOperationType(structures.OperationType((*url.URL)(paymentResponse.Gateway.RedirectURL).String()))

	result = &RetrieveFormAssembly{opHTTPData: opd}
	return
}

// Implement methods for CreateTokenAssembly structure, form pck structures OperationRequestInterface

// GetHTTPMethod return HTTP method which will be used for send request
func (op *RetrieveFormAssembly) GetHTTPMethod() string {
	return op.opHTTPData.GetHTTPMethod()
}

// GetOperationType return part of route path which will be used for send request
func (op *RetrieveFormAssembly) GetOperationType() structures.OperationType {
	return op.opHTTPData.GetOperationType()
}
