package structures

// Card verification allowed modes
const (
	CardVerificationModeInit   = 1
	CardVerificationModeVerify = 2
)

// Payment data source allowed values
const (
	DataSourceCardholder = iota
	DataSourceSaveToGateway
	DataSourceUseGatewaySavedCardholderInitiated
	DataSourceSavingByMerchant
	DataSourceUseMerchantSavedCardholderInitiated
	DataSourceUseGatewaySavedMerchantInitiated
	DataSourceUseMerchantSavedMerchantInitiated
)

// Transact Pro Gateway's request parameters data structures
type (
	// RequestHTTPData contains HTTP request method and operationType to append in URL path
	RequestHTTPData struct {
		// HTTP method
		Method string `json:"-"`
		// Operation type
		OperationType OperationType `json:"-"`
	}

	// GeneralData combined structure about customer data and order data
	GeneralData struct {
		CustomerData CustomerData `json:"customer-data,omitempty"`
		OrderData    OrderData    `json:"order-data,omitempty"`
	}

	// CustomerData structure with detailed fields of customer(cardholder)
	CustomerData struct {
		// Customer (cardholder) email
		Email string `json:"email,omitempty"`
		// Customer (cardholder) phone
		Phone string `json:"phone,omitempty"`
		// Customer (cardholder) birth date in "MMDDYYYY" format
		BirthDate string `json:"birth-date,omitempty"`
		// Customer (cardholder) physical location
		BillingAddress Address `json:"billing-address,omitempty"`
		// Customer (cardholder) address where he want to receive orders
		ShippingAddress Address `json:"shipping-address,omitempty"`
	}

	// OrderData structure with detailed fields of merchant order (transactions)
	OrderData struct {
		// Merchant-side transactions ID
		MerchantTransactionID string `json:"merchant-transaction-id,omitempty"`
		// Merchant-side user ID
		MerchantUserID string `json:"merchant-user-id,omitempty"`
		// Merchant-side order ID
		OrderID string `json:"order-id,omitempty"`
		// Merchant-side order short sms, dms-hold, moto description
		OrderDescription string `json:"order-description,omitempty"`
		// Merchant-side Key-Value order data
		OrderMeta interface{} `json:"order-meta,omitempty"`
		// Merchant-side URL
		MerchantURL string `json:"merchant-side-url,omitempty"`
		// Recipient name
		RecipientName string `json:"recipient-name,omitempty"`
		// Merchant referring name for dynamic descriptor
		MerchantReferringName string `json:"merchant-referring-name,omitempty"`
		// Custom return URL after 3D Secure authentification
		Custom3dReturnURL string `json:"custom-3d-return-url,omitempty"`
		// Custom return URL
		CustomReturnURL string `json:"custom-return-url,omitempty"`
		// Date after which no further authorizations shall be performed (format: YYYYMMDD)
		RecurringExpiry string `json:"recurring-expiry,omitempty"`
		// The minimum number of days between authorizations (for variable frequency use 1)
		RecurringFrequency string `json:"recurring-frequency,omitempty"`
		// Must be set to true for UCOF initialization if any subsequent MIT transactions are supposed to be
		MITsExpected bool `json:"mits-expected,omitempty"`
		// Must be set to true for initial recurring transaction if amount will not be fixed for subsequent transactions
		VariableAmountRecurring bool `json:"variable-amount-recurring,omitempty"`
	}

	// Address structure with detailed fields of customer(cardholder) place
	Address struct {
		// Billing\Shipping country
		Country string `json:"country,omitempty"`
		// Billing\Shipping state
		State string `json:"state,omitempty"`
		// Billing\Shipping city
		City string `json:"city,omitempty"`
		// Billing\Shipping street
		Street string `json:"street,omitempty"`
		// Billing\Shipping house number
		House string `json:"house,omitempty"`
		// Billing\Shipping flat number
		Flat string `json:"flat,omitempty"`
		// Billing\Shipping zip number
		ZIP string `json:"zip,omitempty"`
	}

	// PaymentMethodData structure with detailed fields of PAN data
	PaymentMethodData struct {
		// Credit card number
		Pan string `json:"pan,omitempty"`
		// Credit card expiry date in mm/yy format
		ExpMmYy string `json:"exp-mm-yy,omitempty"`
		// Credit card protection code
		Cvv string `json:"cvv,omitempty"`
		// Cardholder Name and Surname (Name and Surname on credit card)
		CardholderName string `json:"cardholder-name,omitempty"`
		// External 3-D Secure data for an acquirer
		ExternalMpiData *ExternalMpiData `json:"external-mpi-data,omitempty"`
	}

	// ExternalMpiData is a structure for 3-D Secure data required for an acquirer
	ExternalMpiData struct {
		// protocolVersion received from 3-D Secure
		ProtocolVersion string `json:"protocolVersion,omitempty"`
		// dsTransID received from 3-D Secure
		DsTransID string `json:"dsTransID,omitempty"`
		// xid used for 3-D Secure
		XID string `json:"xid,omitempty"`
		// cavv received from 3-D Secure
		CAVV string `json:"cavv,omitempty"`
		// transStatus received from 3-D Secure
		TransStatus string `json:"transStatus,omitempty"`
	}

	// MoneyData structure with detailed fields about transactions amount and currency
	MoneyData struct {
		// Money amount in minor units
		Amount int `json:"amount,omitempty"`
		// Currency, ISO-4217 format
		Currency string `json:"currency,omitempty"`
	}

	// SystemData structure with fields with customer and merchants IP addressees
	SystemData struct {
		// Cardholder IPv4 address
		UserIP string `json:"user-ip"`
		// Cardholder real IPv4 address in case of proxy
		XForwardedFor string `json:"x-forwarded-for"`

		BrowserAcceptHeader      string `json:"browser-accept-header"`
		BrowserJavaEnabled       bool   `json:"browser-java-enabled"`
		BrowserJavascriptEnabled bool   `json:"browser-javascript-enabled"`
		BrowserLanguage          string `json:"browser-language"`
		BrowserColorDepth        string `json:"browser-color-depth"`
		BrowserScreenHeight      string `json:"browser-screen-height"`
		BrowserScreenWidth       string `json:"browser-screen-width"`
		BrowserTz                string `json:"browser-tz"`
		BrowserUserAgent         string `json:"browser-user-agent"`
	}

	// CommandData structure with fields to set various payment processing modes
	CommandData struct {
		CardVerificationMode    uint   `json:"card-verification,omitempty"`
		PaymentMethodDataSource uint   `json:"payment-method-data-source"`
		PaymentMethodDataToken  string `json:"payment-method-data-token,omitempty"`
	}

	// CommandDataGWTransactionID is single structures fields for CommandData, it's used not for any operation
	CommandDataGWTransactionID struct {
		// Previously created Transaction in Transact Pro system
		GWTransactionID string `json:"gateway-transaction-id,omitempty"`
	}

	// CommandDataFormID is single structures fields for CommandData, it's used not for any operation
	CommandDataFormID struct {
		// Inside form ID when selecting non-default form manually, allowed in sms, dms-hold
		FormID string `json:"form-id,omitempty"`
	}

	// CommandDataTerminalMID is single structures fields for CommandData, it's used not for any operation
	CommandDataTerminalMID struct {
		// TerminalMID when selecting terminal manually, allowed in sms, dms-hold
		TerminalMID string `json:"terminal-mid,omitempty"`
	}

	// CommandDataExploreGWTransactionIDs used in explore operations and contains the slice of GWTransactionID's
	CommandDataExploreGWTransactionIDs struct {
		// Previously created Transaction in Transact Pro system
		GWTransactionIDs []string `json:"gateway-transaction-ids,omitempty"`
	}

	// CommandDataExploreMerchantTransactionIDs used in explore operations and contains the slice of MerchantTransactionIDs's
	CommandDataExploreMerchantTransactionIDs struct {
		// Previously created Transaction in Transact Pro system
		MerchantTransactionIDs []string `json:"merchant-transaction-ids,omitempty"`
	}

	// Verify3dEnrollmentData structure for verify 3-D Secure enrollment request
	Verify3dEnrollmentData struct {
		// Credit card number
		Pan string `json:"pan,omitempty"`
		// TerminalMID
		TerminalMID string `json:"terminal-mid,omitempty"`
		// Currency, ISO-4217 format
		Currency string `json:"currency,omitempty"`
	}

	// VerifyCardData structure for card verification completion request
	VerifyCardData struct {
		// Previously created Transaction in Transact Pro system
		GWTransactionID string `json:"gateway-transaction-id,omitempty"`
	}
)

// OperationRequestInterface contains two methods, witch allows to get binned information about operation request
type OperationRequestInterface interface {
	GetHTTPMethod() string
	GetOperationType() OperationType
}

// OperationRequestHTTPData contained with HTTP method (POST, GET) and action URL path for operation (sms, dms_hold)
type OperationRequestHTTPData struct {
	methodHTTP    string
	operationType OperationType
}

// GetHTTPMethod return HTTP method which will be used for send request
func (op *OperationRequestHTTPData) GetHTTPMethod() string {
	return op.methodHTTP
}

// SetHTTPMethod HTTP method which will be used for send request
func (op *OperationRequestHTTPData) SetHTTPMethod(method string) {
	op.methodHTTP = method
}

// GetOperationType return part of route path which will be used for send request
func (op *OperationRequestHTTPData) GetOperationType() OperationType {
	return op.operationType
}

// SetOperationType part of route path which will be used for send request
func (op *OperationRequestHTTPData) SetOperationType(opt OperationType) {
	op.operationType = opt
}
