package transactions

import "bitbucket.transactpro.lv/tls/gw3-go-client/structures"

// P2PAssembly is default structure for Offline Transactions (P2P) transactions operation
type P2PAssembly struct {
	// HTTPData contains HTTP request method and operation action value for request in URL path
	opHTTPData structures.OperationRequestHTTPData
	// Command Data, isn't for any request type and in that case it's combined
	CommandData struct {
		structures.CommandDataTerminalMID
	} `json:"command-data,omitempty"`
	GeneralData   structures.GeneralData       `json:"general-data,omitempty"`
	PaymentMethod structures.PaymentMethodData `json:"payment-method-data"`
	Money         structures.MoneyData         `json:"money-data"`
	// System data contains user(cardholder) IPv4 address and IPv4 address in case of proxy
	System structures.SystemData `json:"system"`
}

// NewP2PAssembly returns new instance with prepared HTTP request data P2PAssembly
func NewP2PAssembly() *P2PAssembly {
	// Predefine default HTTP request data for sms operations
	var opd structures.OperationRequestHTTPData

	opd.SetHTTPMethod("POST")
	opd.SetOperationType(structures.P2P)

	return &P2PAssembly{
		opHTTPData: opd,
	}
}

// Implement methods for P2PAssembly structure, form pck structures OperationRequestInterface

// GetHTTPMethod return HTTP method which will be used for send request
func (op *P2PAssembly) GetHTTPMethod() string {
	return op.opHTTPData.GetHTTPMethod()
}

// GetOperationType return part of route path which will be used for send request
func (op *P2PAssembly) GetOperationType() structures.OperationType {
	return op.opHTTPData.GetOperationType()
}
