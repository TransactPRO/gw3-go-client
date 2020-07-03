package transactions

import (
	"net/http"

	"github.com/TransactPRO/gw3-go-client/structures"
)

// HoldDMSAssembly is default structure for Double-Message Transactions (DMS) Hold transactions operation
type HoldDMSAssembly struct {
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

// NewHoldDMSAssembly returns new instance with prepared HTTP request data HoldDMSAssembly
func NewHoldDMSAssembly() *HoldDMSAssembly {
	// Predefine default HTTP request data for sms operations
	var opd structures.OperationRequestHTTPData

	opd.SetHTTPMethod(http.MethodPost)
	opd.SetOperationType(structures.DMSHold)

	return &HoldDMSAssembly{
		opHTTPData: opd,
	}
}

// Implement methods for HoldDMSAssembly structure, form pck structures OperationRequestInterface

// GetHTTPMethod return HTTP method which will be used for send request
func (op *HoldDMSAssembly) GetHTTPMethod() string {
	return op.opHTTPData.GetHTTPMethod()
}

// GetOperationType return part of route path which will be used for send request
func (op *HoldDMSAssembly) GetOperationType() structures.OperationType {
	return op.opHTTPData.GetOperationType()
}

// ParseResponse parses Gateway response into corresponding data structure
func (op *HoldDMSAssembly) ParseResponse(response *structures.GatewayResponse) (result *structures.TransactionResponse, err error) {
	result = new(structures.TransactionResponse)
	err = response.ParseJSON(&result)
	return
}
