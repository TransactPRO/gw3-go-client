package structures

// Transact Pro Gateway's request parameters data structures
type (
	AuthData struct {
		// Transact Pro Account ID
		AccountID int 		`json:"account-id"`
		// Transact Pro Merchant Password
		SecretKey string 	`json:"secret-key"`
	}

	CustomerData struct {
		// Customer (cardholder) email
		Email 		string  `json:"email,omitempty"`
		// Customer (cardholder) physical location
		BillingAddress 	Address `json:"billing-address,omitempty"`
		// Customer (cardholder) address where he want to receive orders
		ShippingAddress Address `json:"shipping-address,omitempty"`
	}

	OrderData struct {
		// Merchant-side transaction ID
		MerchantTransactionID 	string `json:"merchant-transaction-id,omitempty"`
		// Merchant-side user ID
		MerchantID 		string `json:"merchant-user-id,omitempty"`
		// Merchant-side order ID
		OrderID 		string `json:"order-id,omitempty"`
		// Merchant-side order short sms, dms-hold, moto description
		OrderDescription 	string `json:"order-description,omitempty"`
		// Merchant-side Key-Value order data
		OrderMeta		interface{} `json:"order-meta,omitempty"`
	}

	Address struct {
		// Billing\Shipping country
		Country string `json:"Country,omitempty"`
		// Billing\Shipping state
		State 	string `json:"State,omitempty"`
		// Billing\Shipping city
		City 	string `json:"city,omitempty"`
		// Billing\Shipping street
		Street 	string `json:"street,omitempty"`
		// Billing\Shipping house number
		House 	string `json:"house,omitempty"`
		// Billing\Shipping flat number
		Flat 	string `json:"flat,omitempty"`
		// Billing\Shipping zip number
		ZIP 	string `json:"zip,omitempty"`
	}

	PaymentMethodData struct {
		// Credit card number
		Pan 		string `json:"pan"`
		// Credit card expiry date in mm/yy format
		ExpMmYy 	string `json:"exp-mm-yy"`
		// Credit card protection code
		Cvv 		string `json:"cvv,omitempty"`
		// Cardholder Name and Surname (Name and Surname on credit card)
		CardholderName 	string `json:"cardholder-name,omitempty"`
	}

	MoneyData struct {
		// Money amount in minor units
		Amount 		int 	`json:"amount"`
		// Currency, ISO-4217 format
		Currency 	string 	`json:"currency"`
	}

	SystemData struct {
		// Cardholder IPv4 address
		UserIP		string `json:"user-ip"`
		// Cardholder real IPv4 address in case of proxy
		XForwardedFor 	string `json:"x-forwarded-for"`
	}

	// Single structures fields for COMMAND DATA BUNDLE, but not for any request type is allowed
	CommandDataGWTransactionID struct {
		// Previously created Transaction in Transact Pro system
		GWTransactionID string 		`json:"gateway-transaction-id"`
	}
)
