package structures

// OperationType describes the operation action as string
type OperationType string

// OperationType possible values
const (
	/*
		Transaction Types
	*/

	// SMS transactions request type for url route path
	SMS OperationType = "sms"
	// DMSHold transactions request type for url route path
	DMSHold OperationType = "hold-dms"
	// DMSCharge transactions request type for url route path
	DMSCharge OperationType = "charge-dms"
	// CANCEL transactions request type for url route path
	CANCEL OperationType = "cancel"
	// MOTOSMS transactions request type for url route path
	MOTOSMS OperationType = "moto/sms"
	// MOTODMS transactions request type for url route path
	MOTODMS OperationType = "moto/dms"
	// CREDIT transactions request type for url route path
	CREDIT OperationType = "credit"
	// P2P transactions request type for url route path
	P2P OperationType = "p2p"
	// B2P transactions request type for url route path
	B2P OperationType = "b2p"
	// RecurrentSMS transactions request type for url route path
	InitRecurrentSMS OperationType = "recurrent/sms/init"
	// RecurrentSMS transactions request type for url route path
	RecurrentSMS OperationType = "recurrent/sms"
	// RecurrentDMS transactions request type for url route path
	InitRecurrentDMS OperationType = "recurrent/dms/init"
	// Refund transactions request type for url route path
	RecurrentDMS OperationType = "recurrent/dms"
	// Refund transactions request type for url route path
	Refund OperationType = "refund"
	// Reversal transactions request type for url route path
	Reversal OperationType = "reversal"

	/*
		Exploring Past Payments types
	*/

	// ExploringStatus is a transaction status request type for url route path
	ExploringStatus OperationType = "status"
	// ExploringResult is a transaction result request type for url route path
	ExploringResult OperationType = "result"
	// ExploringHistory is a transaction history request type for url route path
	ExploringHistory OperationType = "history"
	// ExploringRecurrents is a transaction history request type for url route path
	ExploringRecurrents OperationType = "recurrents"
	// ExploringRefunds is a transaction history request type for url route path
	ExploringRefunds OperationType = "refunds"
	// ExploringLimits is limits request type for url route path
	ExploringLimits OperationType = "limits"

	/*
		Verification requests
	*/
	// Verify card 3-D Secure enrollment request type for url route path
	Verify3dEnrollment OperationType = "verify/3d-enrollment"
	// Verify card completion request type for url route path
	VerifyCard OperationType = "verify/card"

	/*
		Tokenization requests
	*/
	// Create payment data token url route path
	CreateToken OperationType = "token/create"

	/*
		Reporting requests
	*/
	Report OperationType = "report"
)
