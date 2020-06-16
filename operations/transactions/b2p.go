package transactions

import (
	"net/http"

	"github.com/TransactPRO/gw3-go-client/structures"
)

// B2PAssembly is default structure for Offline Transactions (B2P) transactions operation
type B2PAssembly struct {
	// HTTPData contains HTTP request method and operation action value for request in URL path
	opHTTPData structures.OperationRequestHTTPData
	// Command Data, isn't for any request type and in that case it's combined
	CommandData struct {
		structures.CommandData
		structures.CommandDataFormID
		structures.CommandDataTerminalMID
	} `json:"command-data,omitempty"`
	GeneralData   structures.GeneralData       `json:"general-data,omitempty"`
	PaymentMethod structures.PaymentMethodData `json:"payment-method-data,omitempty"`
	Money         structures.MoneyData         `json:"money-data"`
	// System data contains user(cardholder) IPv4 address and IPv4 address in case of proxy
	System structures.SystemData `json:"system"`
}

// NewB2PAssembly returns new instance with prepared HTTP request data B2PAssembly
func NewB2PAssembly() *B2PAssembly {
	// Predefine default HTTP request data for sms operations
	var opd structures.OperationRequestHTTPData

	opd.SetHTTPMethod(http.MethodPost)
	opd.SetOperationType(structures.B2P)

	return &B2PAssembly{
		opHTTPData: opd,
	}
}

// Implement methods for B2PAssembly structure, form pck structures OperationRequestInterface

// GetHTTPMethod return HTTP method which will be used for send request
func (op *B2PAssembly) GetHTTPMethod() string {
	return op.opHTTPData.GetHTTPMethod()
}

// GetOperationType return part of route path which will be used for send request
func (op *B2PAssembly) GetOperationType() structures.OperationType {
	return op.opHTTPData.GetOperationType()
}

// ParseResponse parses Gateway response into corresponding data structure
func (op *B2PAssembly) ParseResponse(response *structures.GatewayResponse) (result *structures.TransactionResponse, err error) {
	result = new(structures.TransactionResponse)
	err = response.ParseJSON(&result)
	return
}
