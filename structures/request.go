package structures

// Transact Pro Gateway's request parameters data structures
type (
	AuthData struct {
		// Transact Pro Account ID
		AccountID int 		`json:"account-id"`
		// Transact Pro Merchant Password
		SecretKey string 	`json:"secret-key"`
	}

	CommandData struct {
		// Previously created Transaction ID
		GWTranID 	string 		`json:"gateway-transaction-id"`
		// Inside form ID when selecting non-default form manually
		FormID 		string		`json:"form-id"`
		// Terminal MID when selecting terminal manually
		TerminalMID	string		`json:"terminal-mid"`
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
)
