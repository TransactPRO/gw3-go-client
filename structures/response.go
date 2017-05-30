package structures

// Transact Pro Gateway's response data structures
type (
	// Unauthorized response will be return if Merchant authorization was incorrect, HTTP status code will be 401
	// Example of json: { "msg": "Unauthorized", "status": 401 }
	UnauthorizedResponse struct {
		Msg    string `json:"msg"`
		Status int    `json:"status"`
	}

	// SMSResponse is structure of SMS operation response
	SMSResponse struct {
		GateWay         gateWay         `json:"gw"`
		Error           interface{}     `json:"error"`
		AcquirerDetails acquirerDetails `json:"acquirer-details"`
	}

	// Data parts for different responses

	// GateWay Transact Pro system response
	gateWay struct {
		GatewayTransactionID string `json:"gateway-transaction-id,omitempty"`
		StatusCode           int    `json:"status-code,omitempty"`
		StatusText           string `json:"status-text,omitempty"`
	}

	// AcquirerDetails response translated via Transact Pro system
	acquirerDetails struct {
		EciSLi            int    `json:"eci-sli,omitempty"`
		TerminalID        string `json:"terminal-mid,omitempty"`
		TransactionID     string `json:"transaction-id,omitempty"`
		ResultCode        string `json:"result-code,omitempty"`
		StatusText        string `json:"status-text,omitempty"`
		StatusDescription string `json:"status-description,omitempty"`
	}
)
