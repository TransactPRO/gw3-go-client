package request

// The list of all Transact Pro Gateway's structure bundles
type (
	// AuthData (Auhentication data), than data is mandatory to authorize merchant request in Transact Pro system
	AuthData struct {
		AccountID int
		SecretKey string
		SessionID string
	}
)