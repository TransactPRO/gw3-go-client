package exploring

import "github.com/TransactPRO/gw3-go-client/structures"

// ExploreTransactionAssembly is default structure for transaction status operation
type ExploreTransactionAssembly struct {
	// HTTPData contains HTTP request method and operation action value for request in URL path
	opHTTPData structures.OperationRequestHTTPData
	// Command Data, isn't for any request type and in that case it's combined
	CommandData struct {
		structures.CommandDataExploreGWTransactionIDs
		structures.CommandDataExploreMerchantTransactionIDs
	} `json:"command-data,omitempty"`
	// System data contains user(cardholder) IPv4 address and IPv4 address in case of proxy
	System structures.SystemData `json:"system"`
}

// NewStatusAssembly returns new instance with prepared HTTP request data ExploreTransactionAssembly
func NewStatusAssembly() *ExploreTransactionAssembly {
	// Predefine default HTTP request data for sms operations
	var opd structures.OperationRequestHTTPData

	opd.SetHTTPMethod("POST")
	opd.SetOperationType(structures.ExploringStatus)

	return &ExploreTransactionAssembly{
		opHTTPData: opd,
	}
}

// NewResultAssembly returns new instance with prepared HTTP request data ExploreTransactionAssembly
func NewResultAssembly() *ExploreTransactionAssembly {
	// Predefine default HTTP request data for sms operations
	var opd structures.OperationRequestHTTPData

	opd.SetHTTPMethod("POST")
	opd.SetOperationType(structures.ExploringResult)

	return &ExploreTransactionAssembly{
		opHTTPData: opd,
	}
}

// NewHistoryAssembly returns new instance with prepared HTTP request data ExploreTransactionAssembly
func NewHistoryAssembly() *ExploreTransactionAssembly {
	// Predefine default HTTP request data for sms operations
	var opd structures.OperationRequestHTTPData

	opd.SetHTTPMethod("POST")
	opd.SetOperationType(structures.ExploringHistory)

	return &ExploreTransactionAssembly{
		opHTTPData: opd,
	}
}

// NewRecurrentsAssembly returns new instance with prepared HTTP request data ExploreTransactionAssembly
func NewRecurrentsAssembly() *ExploreTransactionAssembly {
	// Predefine default HTTP request data for sms operations
	var opd structures.OperationRequestHTTPData

	opd.SetHTTPMethod("POST")
	opd.SetOperationType(structures.ExploringRecurrents)

	return &ExploreTransactionAssembly{
		opHTTPData: opd,
	}
}

// NewRefundsAssembly returns new instance with prepared HTTP request data ExploreTransactionAssembly
func NewRefundsAssembly() *ExploreTransactionAssembly {
	// Predefine default HTTP request data for sms operations
	var opd structures.OperationRequestHTTPData

	opd.SetHTTPMethod("POST")
	opd.SetOperationType(structures.ExploringRefunds)

	return &ExploreTransactionAssembly{
		opHTTPData: opd,
	}
}

// Implement methods for ExploreAssembly structure, form pck structures OperationRequestInterface

// GetHTTPMethod return HTTP method which will be used for send request
func (op *ExploreTransactionAssembly) GetHTTPMethod() string {
	return op.opHTTPData.GetHTTPMethod()
}

// GetOperationType return part of route path which will be used for send request
func (op *ExploreTransactionAssembly) GetOperationType() structures.OperationType {
	return op.opHTTPData.GetOperationType()
}
