package structures

// Transact Pro Gateway's response structures
type (
	// UnauthorizedResponse will be return if Merchant authorization was incorrect, HTTP status code will be 401
	// Example of json: { "msg": "Unauthorized", "status": 401 }
	UnauthorizedResponse struct {
		Msg    string `json:"msg,omitempty"`
		Status int    `json:"status,omitempty"`
	}

	// TransactionResponse is structure of all transaction operations
	TransactionResponse struct {
		GateWay         gateWay         `json:"gw"`
		Error           gateWayError    `json:"error"`
		AcquirerDetails acquirerDetails `json:"acquirer-details"`
	}

	// Data parts for responses

	// gateWay Transact Pro system response
	gateWay struct {
		GatewayTransactionID string `json:"gateway-transaction-id,omitempty"`
		StatusCode           int    `json:"status-code,omitempty"`
		StatusText           string `json:"status-text,omitempty"`
	}
	// gateWayError error structure part
	gateWayError struct {
		// Gateway error code
		Code int `json:"code,omitempty"`
		// Gateway error description
		Msg string `json:"msg,omitempty"`
	}

	// acquirerDetails response translated via Transact Pro system
	acquirerDetails struct {
		EciSLi            int    `json:"eci-sli,omitempty"`
		TerminalID        string `json:"terminal-mid,omitempty"`
		TransactionID     string `json:"transaction-id,omitempty"`
		ResultCode        string `json:"result-code,omitempty"`
		StatusText        string `json:"status-text,omitempty"`
		StatusDescription string `json:"status-description,omitempty"`
	}

	// ExploringStatusResponse the structure of all exploring operations
	// contained asked Transact Pro transaction id and it's status data
	// note: the response will be Json object to parse response you must assign to []ExploringStatusResponse
	ExploringStatusResponse struct {
		// GatewayTransactionID the past Transact Pro gateway transaction id
		GatewayTransactionID string `json:"gateway-transaction-id,omitempty"`
		// Status contained informational data of transaction
		Status []ExploreStatus `json:"status"`
	}

	// ExploringResultResponse the structure of all exploring operations
	// contained asked Transact Pro transaction id and it's result data
	// note: the response will be Json object to parse response you must assign to []ExploringResultResponse
	ExploringResultResponse struct {
		// GatewayTransactionID the past Transact Pro gateway transaction id
		GatewayTransactionID string `json:"gateway-transaction-id,omitempty"`
		// Status contained informational data of transaction
		Status       []ExploreStatus `json:"status"`
		DateCreated  string          `json:"date-created"`
		DateFinished string          `json:"date-finished"`
		ResultData   ExploreResult   `json:"result-data"`
	}

	// ExploreStatus the structure of Transact Pro statuses for past transaction
	ExploreStatus struct {
		// GatewayTransactionID past transaction ID in Transact Pro system
		GatewayTransactionID string `json:"gateway-transaction-id,omitempty"`
		// StatusCode transaction status code
		StatusCode int `json:"status-code,omitempty"`
		// StatusText transaction status string representation
		StatusText string `json:"status-text,omitempty"`
		// StatusCodeGeneral transaction status code
		StatusCodeGeneral int `json:"status-code-general,omitempty"`
		// StatusTextGeneral transaction status code
		StatusTextGeneral string `json:"status-text-general,omitempty"`
	}

	// ExploreResult structure of past transaction result
	ExploreResult struct {
		GateWay         gateWay         `json:"gw"`
		Error           gateWayError    `json:"error"`
		AcquirerDetails acquirerDetails `json:"acquirer-details"`
	}
)
