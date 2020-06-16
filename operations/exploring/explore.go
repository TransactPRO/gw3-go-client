package exploring

import (
	"net/http"

	"github.com/TransactPRO/gw3-go-client/structures"
)

type (
	// ExploreTransactionAssembly is base structure for transactions exploring operation
	ExploreTransactionAssembly struct {
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

	// ExploreStatusAssembly is a transactions status exploring operation
	ExploreStatusAssembly ExploreTransactionAssembly
	// ExploreResultAssembly is a transactions result exploring operation
	ExploreResultAssembly ExploreTransactionAssembly
	// ExploreHistoryAssembly is a transactions history exploring operation
	ExploreHistoryAssembly ExploreTransactionAssembly
	// ExploreRefundsAssembly is a transactions refunds exploring operation
	ExploreRefundsAssembly ExploreTransactionAssembly
	// ExploreRecurrentsAssembly is a transactions subsequent recurring transactions exploring operation
	ExploreRecurrentsAssembly ExploreTransactionAssembly

	// ExploreLimitsAssembly is a structure for object limits exploring operation
	ExploreLimitsAssembly struct {
		// HTTPData contains HTTP request method and operation action value for request in URL path
		opHTTPData structures.OperationRequestHTTPData
	}
)

// NewStatusAssembly returns new instance with prepared HTTP request data ExploreStatusAssembly
func NewStatusAssembly() *ExploreStatusAssembly {
	// Predefine default HTTP request data for sms operations
	var opd structures.OperationRequestHTTPData

	opd.SetHTTPMethod(http.MethodPost)
	opd.SetOperationType(structures.ExploringStatus)

	return &ExploreStatusAssembly{
		opHTTPData: opd,
	}
}

// NewResultAssembly returns new instance with prepared HTTP request data ExploreResultAssembly
func NewResultAssembly() *ExploreResultAssembly {
	// Predefine default HTTP request data for sms operations
	var opd structures.OperationRequestHTTPData

	opd.SetHTTPMethod(http.MethodPost)
	opd.SetOperationType(structures.ExploringResult)

	return &ExploreResultAssembly{
		opHTTPData: opd,
	}
}

// NewHistoryAssembly returns new instance with prepared HTTP request data ExploreHistoryAssembly
func NewHistoryAssembly() *ExploreHistoryAssembly {
	// Predefine default HTTP request data for sms operations
	var opd structures.OperationRequestHTTPData

	opd.SetHTTPMethod(http.MethodPost)
	opd.SetOperationType(structures.ExploringHistory)

	return &ExploreHistoryAssembly{
		opHTTPData: opd,
	}
}

// NewRecurrentsAssembly returns new instance with prepared HTTP request data ExploreRecurrentsAssembly
func NewRecurrentsAssembly() *ExploreRecurrentsAssembly {
	// Predefine default HTTP request data for sms operations
	var opd structures.OperationRequestHTTPData

	opd.SetHTTPMethod(http.MethodPost)
	opd.SetOperationType(structures.ExploringRecurrents)

	return &ExploreRecurrentsAssembly{
		opHTTPData: opd,
	}
}

// NewRefundsAssembly returns new instance with prepared HTTP request data ExploreRefundsAssembly
func NewRefundsAssembly() *ExploreRefundsAssembly {
	// Predefine default HTTP request data for sms operations
	var opd structures.OperationRequestHTTPData

	opd.SetHTTPMethod(http.MethodPost)
	opd.SetOperationType(structures.ExploringRefunds)

	return &ExploreRefundsAssembly{
		opHTTPData: opd,
	}
}

// NewLimitsAssembly returns new instance with prepared HTTP request data ExploreLimitsAssembly
func NewLimitsAssembly() *ExploreLimitsAssembly {
	// Predefine default HTTP request data for sms operations
	var opd structures.OperationRequestHTTPData

	opd.SetHTTPMethod(http.MethodPost)
	opd.SetOperationType(structures.ExploringLimits)

	return &ExploreLimitsAssembly{
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

// ParseResponse parses Gateway response into corresponding data structure
func (op *ExploreStatusAssembly) ParseResponse(response *structures.GatewayResponse) (result *structures.ExploringStatusResponse, err error) {
	result = new(structures.ExploringStatusResponse)
	err = response.ParseJSON(&result)
	return
}

// ParseResponse parses Gateway response into corresponding data structure
func (op *ExploreResultAssembly) ParseResponse(response *structures.GatewayResponse) (result *structures.ExploringResultResponse, err error) {
	result = new(structures.ExploringResultResponse)
	err = response.ParseJSON(&result)
	return
}

// ParseResponse parses Gateway response into corresponding data structure
func (op *ExploreHistoryAssembly) ParseResponse(response *structures.GatewayResponse) (result *structures.ExploringHistoryResponse, err error) {
	result = new(structures.ExploringHistoryResponse)
	err = response.ParseJSON(&result)
	return
}

// ParseResponse parses Gateway response into corresponding data structure
func (op *ExploreRecurrentsAssembly) ParseResponse(response *structures.GatewayResponse) (result *structures.ExploringRecurrentsResponse, err error) {
	result = new(structures.ExploringRecurrentsResponse)
	err = response.ParseJSON(&result)
	return
}

// ParseResponse parses Gateway response into corresponding data structure
func (op *ExploreRefundsAssembly) ParseResponse(response *structures.GatewayResponse) (result *structures.ExploringRefundsResponse, err error) {
	result = new(structures.ExploringRefundsResponse)
	err = response.ParseJSON(&result)
	return
}

// GetHTTPMethod return HTTP method which will be used for send request
func (op *ExploreLimitsAssembly) GetHTTPMethod() string {
	return op.opHTTPData.GetHTTPMethod()
}

// GetOperationType return part of route path which will be used for send request
func (op *ExploreLimitsAssembly) GetOperationType() structures.OperationType {
	return op.opHTTPData.GetOperationType()
}

// ParseResponse parses Gateway response into corresponding data structure
func (op *ExploreLimitsAssembly) ParseResponse(response *structures.GatewayResponse) (result *structures.ExploringLimitsResponse, err error) {
	result = new(structures.ExploringLimitsResponse)
	err = response.ParseJSON(&result)
	return
}
