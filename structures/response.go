package structures

// Transact Pro Gateway's response data structures
type (
	// ResponseSMS is structure of sms response
	ResponseSMS struct {
		GateWay         `json:"gw"`
		Error           interface{} `json:"error"`
		AcquirerDetails `json:"acquirer-details"`
	}

	// Data parts for different responses

	// GateWay Transact Pro system response
	GateWay struct {
		GatewayTransactionID string `json:"gateway-transaction-id,omitempty"`
		StatusCode           int    `json:"status-code,omitempty"`
		StatusText           string `json:"status-text,omitempty"`
	}

	// AcquirerDetails response translated via Transact Pro system
	AcquirerDetails struct {
		EciSLi            int    `json:"eci-sli,omitempty"`
		TerminalID        string `json:"terminal-mid,omitempty"`
		TransactionID     string `json:"transaction-id,omitempty"`
		ResultCode        string `json:"result-code,omitempty"`
		StatusText        string `json:"status-text,omitempty"`
		StatusDescription string `json:"status-description,omitempty"`
	}
)
