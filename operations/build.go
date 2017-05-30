package operations

import (
	"bitbucket.transactpro.lv/tls/gw3-go-client/operations/transaction"
)

// Builder operation structure builder for specific request
type Builder struct{}

// NewSms returns new instance to new SMS structure
func (ob *Builder) NewSms() *transaction.SMSAssembly {
	// Get new prepared sms structure for assembly
	return transaction.NewSMSAssembly()
}

// NewHoldDMS returns new instance to new HoldDMS structure
func (ob *Builder) NewHoldDMS() *transaction.HoldDMSAssembly {
	// Get new prepared sms structure for assembly
	return transaction.NewHoldDMSAssembly()
}

// NewChargeDMS returns new instance to new ChargeDMS  structure
func (ob *Builder) NewChargeDMS() *transaction.ChargeDMSAssembly {
	// Get new prepared sms structure for assembly
	return transaction.NewChargeDMSAssembly()
}

// CancelAssembly returns new instance to new ChargeDMS  structure
func (ob *Builder) NewCancel() *transaction.CancelAssembly {
	// Get new prepared sms structure for assembly
	return transaction.NewCancelAssembly()
}
