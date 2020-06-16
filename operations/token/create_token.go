package token

import (
	"net/http"

	"github.com/TransactPRO/gw3-go-client/structures"
)

// CreateTokenAssembly is default structure for payment data tokenization operation
type CreateTokenAssembly struct {
	opHTTPData structures.OperationRequestHTTPData

	CommandData struct {
		structures.CommandData
		structures.CommandDataFormID
		structures.CommandDataTerminalMID
	} `json:"command-data,omitempty"`
	GeneralData   structures.GeneralData       `json:"general-data,omitempty"`
	PaymentMethod structures.PaymentMethodData `json:"payment-method-data,omitempty"`
	Money         structures.MoneyData         `json:"money-data"`
	System        structures.SystemData        `json:"system"`
}

// NewCreateTokenAssembly returns new instance with prepared HTTP request data CreateTokenAssembly
func NewCreateTokenAssembly() *CreateTokenAssembly {
	// Predefine default HTTP request data for create token operations
	var opd structures.OperationRequestHTTPData

	opd.SetHTTPMethod(http.MethodPost)
	opd.SetOperationType(structures.CreateToken)

	return &CreateTokenAssembly{
		opHTTPData: opd,
	}
}

// Implement methods for CreateTokenAssembly structure, form pck structures OperationRequestInterface

// GetHTTPMethod return HTTP method which will be used for send request
func (op *CreateTokenAssembly) GetHTTPMethod() string {
	return op.opHTTPData.GetHTTPMethod()
}

// GetOperationType return part of route path which will be used for send request
func (op *CreateTokenAssembly) GetOperationType() structures.OperationType {
	return op.opHTTPData.GetOperationType()
}

// ParseResponse parses Gateway response into corresponding data structure
func (op *CreateTokenAssembly) ParseResponse(response *structures.GatewayResponse) (result *structures.TransactionResponse, err error) {
	result = new(structures.TransactionResponse)
	err = response.ParseJSON(result)
	return
}
