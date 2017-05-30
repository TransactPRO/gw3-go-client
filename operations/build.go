package operations

import (
	"bitbucket.transactpro.lv/tls/gw3-go-client/operations/transaction"
)

// Builder operation structure builder for specific request
type Builder struct{}

// NewSms returns new pointer to new SMS structure
func (ob *Builder) NewSms() *transaction.SMSAssembly {
	// Get new prepared sms structure for assembly
	return transaction.NewSMSAssembly()
}
