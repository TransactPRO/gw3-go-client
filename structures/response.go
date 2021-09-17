package structures

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

// Transact Pro Gateway's response structures
type (
	// GatewayResponse represents generic Gateway response wrapper
	GatewayResponse struct {
		*http.Response
		Payload []byte
		Digest  *ResponseDigest
	}

	// Error is a generic error structure that might be returned with any response
	// collections responses (like refunds exploration) may contain errors
	// as response collection items
	Error struct {
		Code    ErrorCode `json:"code,omitempty"`
		Message string    `json:"message,omitempty"`
	}

	// TransactionResponse is structure of all transaction operations
	TransactionResponse struct {
		Gateway         Gateway         `json:"gw,omitempty"`
		Error           Error           `json:"error,omitempty"`
		AcquirerDetails AcquirerDetails `json:"acquirer-details,omitempty"`
		Warnings        []string        `json:"warnings,omitempty"`
	}

	// CallbackResult is a structure for callback payload JSON
	CallbackResult struct {
		Error      Error               `json:"error,omitempty"`
		ResultData TransactionResponse `json:"result-data,omitempty"`
	}

	// Gateway Transact Pro system response
	Gateway struct {
		GatewayTransactionID         string  `json:"gateway-transaction-id,omitempty"`
		MerchantTransactionID        string  `json:"merchant-transaction-id,omitempty"`
		StatusCode                   Status  `json:"status-code,omitempty"`
		StatusText                   string  `json:"status-text,omitempty"`
		RedirectURL                  *URL    `json:"redirect-url,omitempty"`
		OriginalGatewayTransactionID *string `json:"original-gateway-transaction-id,omitempty"`
		ParentGatewayTransactionID   *string `json:"parent-gateway-transaction-id,omitempty"`
	}

	// AcquirerDetails response translated via Transact Pro system
	AcquirerDetails struct {
		DynamicDescriptor string `json:"dynamic-descriptor,omitempty"`
		EciSli            string `json:"eci-sli,omitempty"`
		TerminalID        string `json:"terminal-mid,omitempty"`
		TransactionID     string `json:"transaction-id,omitempty"`
		ResultCode        string `json:"result-code,omitempty"`
		StatusText        string `json:"status-text,omitempty"`
		StatusDescription string `json:"status-description,omitempty"`
	}

	// Limit represents an object's limit instance
	Limit struct {
		CounterType          string      `json:"counter-type,omitempty"`
		Currency             string      `json:"currency,omitempty"`
		Limit                json.Number `json:"limit,omitempty"`
		PaymentMethodSubtype string      `json:"payment-method-subtype,omitempty"`
		PaymentMethodType    string      `json:"payment-method-type,omitempty"`
		Value                json.Number `json:"value,omitempty"`
	}

	// TransactionInfo represents one transaction's information
	TransactionInfo struct {
		AccountGUID           string      `json:"account-guid,omitempty"`
		AcqTerminalID         string      `json:"acq-terminal-id,omitempty"`
		AcqTransactionID      string      `json:"acq-transaction-id,omitempty"`
		Amount                json.Number `json:"amount,omitempty"`
		ApprovalCode          string      `json:"approval-code,omitempty"`
		CardholderName        string      `json:"cardholder-name,omitempty"`
		Currency              string      `json:"currency,omitempty"`
		DateFinished          Time        `json:"date-finished,omitempty"`
		EciSli                string      `json:"eci-sli,omitempty"`
		GatewayTransactionID  string      `json:"gateway-transaction-id,omitempty"`
		MerchantTransactionID string      `json:"merchant-transaction-id,omitempty"`
		StatusCode            Status      `json:"status-code,omitempty"`
		StatusCodeGeneral     Status      `json:"status-code-general,omitempty"`
		StatusText            string      `json:"status-text,omitempty"`
		StatusTextGeneral     string      `json:"status-text-general,omitempty"`
	}

	// EnrollmentResponse represents card's 3-D Secure verification response
	EnrollmentResponse struct {
		Error      *Error     `json:"error,omitempty"`
		Enrollment Enrollment `json:"enrollment,omitempty"`
	}

	// ExploringItem represents base item for esploring collection element
	exploringItem struct {
		Error                *Error `json:"error,omitempty"`
		GatewayTransactionID string `json:"gateway-transaction-id,omitempty"`
	}

	// ExploringStatusResponse represents transaction's status response
	ExploringStatusResponse struct {
		Error        *Error                  `json:"error,omitempty"`
		Transactions []TransactionStatusList `json:"transactions,omitempty"`
	}

	// TransactionStatusList represents one item for exploring statuses collection
	TransactionStatusList struct {
		exploringItem
		Status []TransactionStatus `json:"status,omitempty"`
	}

	// TransactionStatus represents one item for exploring results collection
	TransactionStatus struct {
		GatewayTransactionID string     `json:"gateway-transaction-id,omitempty"`
		StatusCode           Status     `json:"status-code,omitempty"`
		StatusText           string     `json:"status-text,omitempty"`
		StatusCodeGeneral    Status     `json:"status-code-general,omitempty"`
		StatusTextGeneral    string     `json:"status-text-general,omitempty"`
		CardMask             string     `json:"card-mask,omitempty"`
		CardFamily           CardFamily `json:"card-family,omitempty"`
	}

	// ExploringResultResponse represents transaction's result response
	ExploringResultResponse struct {
		Error        *Error              `json:"error,omitempty"`
		Transactions []TransactionResult `json:"transactions,omitempty"`
	}

	// TransactionResult represents one item for exploring results collection
	TransactionResult struct {
		exploringItem
		DateCreated  Time                `json:"date-created,omitempty"`
		DateFinished Time                `json:"date-finished,omitempty"`
		ResultData   TransactionResponse `json:"result-data,omitempty"`
	}

	// ExploringHistoryResponse represents transactions' history response
	ExploringHistoryResponse struct {
		Error        *Error               `json:"error,omitempty"`
		Transactions []TransactionHistory `json:"transactions,omitempty"`
	}

	// TransactionHistory represents one item for exploring history collection
	TransactionHistory struct {
		exploringItem
		History []HistoryEvent `json:"history,omitempty"`
	}

	// HistoryEvent represents one history event for a history transaction item
	HistoryEvent struct {
		DateUpdated   Time   `json:"date-updated,omitempty"`
		StatusCodeNew Status `json:"status-code-new,omitempty"`
		StatusCodeOld Status `json:"status-code-old,omitempty"`
		StatusTextNew string `json:"status-text-new,omitempty"`
		StatusTextOld string `json:"status-text-old,omitempty"`
	}

	// ExploringLimitsResponse represents limits response
	ExploringLimitsResponse struct {
		Error         *Error                    `json:"error,omitempty"`
		Type          string                    `json:"type,omitempty"`
		Title         string                    `json:"title,omitempty"`
		AcqTerminalID *string                   `json:"acq-terminal-id,omitempty"`
		Limits        []Limit                   `json:"counters,omitempty"`
		Children      []ExploringLimitsResponse `json:"childs,omitempty"`
	}

	// ExploringRecurrentsResponse represents subsequent recurring transactions response
	ExploringRecurrentsResponse struct {
		Error        *Error                  `json:"error,omitempty"`
		Transactions []TransactionRecurrings `json:"transactions,omitempty"`
	}

	// TransactionRecurrings represents one item for exploring recurring transactions collection
	TransactionRecurrings struct {
		exploringItem
		Subsequent []TransactionInfo `json:"recurrents,omitempty"`
	}

	// ExploringRefundsResponse represents refunds response
	ExploringRefundsResponse struct {
		Error        *Error               `json:"error,omitempty"`
		Transactions []TransactionRefunds `json:"transactions,omitempty"`
	}

	// TransactionRefunds represents one item for exploring refunds collection
	TransactionRefunds struct {
		exploringItem
		Refunds []TransactionInfo `json:"refunds,omitempty"`
	}

	// CsvReport represents parsed CSV report
	CsvReport struct {
		data    []byte
		Headers []string
	}
)

// NewGatewayResponse creates Gateway response wrapper over standard http.Response
func NewGatewayResponse(httpResponse *http.Response, payload []byte) *GatewayResponse {
	return &GatewayResponse{Response: httpResponse, Payload: payload}
}

// Successful returns TRUE for non-declined responses
func (o *GatewayResponse) Successful() bool {
	if o == nil {
		return false
	}

	return (o.StatusCode >= http.StatusOK && o.StatusCode < http.StatusBadRequest) || o.StatusCode == http.StatusPaymentRequired
}

// ParseJSON parses response payload as JSON into given object reference
func (o *GatewayResponse) ParseJSON(output interface{}) error {
	if o == nil {
		return errors.New("cannot unmarshal JSON response: response is nil")
	}

	if err := json.Unmarshal(o.Payload, output); err != nil {
		return fmt.Errorf("cannot unmarshal JSON response: %s", err)
	}

	return nil
}

// NewCsvReport create instance of CsvReport from response payload.
// Payload MUST contain headers line, otherwise an error will be returned.
func NewCsvReport(response *GatewayResponse) (result *CsvReport, err error) {
	result = &CsvReport{data: response.Payload}

	reader := csv.NewReader(bytes.NewReader(result.data))
	result.Headers, err = reader.Read()
	if err == io.EOF {
		err = errors.New("report format error: no headers line")
		return
	}

	return
}

// Iterate iterates over report.
// Each row will be passed to given "process" function as a map where key
// is taken from header line for corresponding position.
// If process returns FALSE, iteration will be stopped.
func (o *CsvReport) Iterate(process func(row map[string]string) bool) (err error) {
	reader := csv.NewReader(bytes.NewReader(o.data))
	reader.FieldsPerRecord = len(o.Headers)
	reader.ReuseRecord = true

	_, _ = reader.Read() // skip headers line
	for {
		record, readErr := reader.Read()
		if readErr == io.EOF {
			return
		}

		if readErr != nil {
			err = readErr
			return
		}

		if len(record) > 0 {
			row := make(map[string]string, len(o.Headers))
			for i := range o.Headers {
				if i < len(record) {
					row[o.Headers[i]] = record[i]
				} else {
					row[o.Headers[i]] = ""
				}
			}

			if !process(row) {
				return
			}
		}
	}
}
