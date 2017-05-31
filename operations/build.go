package operations

import (
	"bitbucket.transactpro.lv/tls/gw3-go-client/operations/exploring_requests"
	"bitbucket.transactpro.lv/tls/gw3-go-client/operations/transactions"
)

// Builder operation structure builder for specific request
type Builder struct{}

// Transaction Types builders

// NewSms returns new instance to new SMS structure
func (ob *Builder) NewSms() *transactions.SMSAssembly {
	// Get new prepared sms structure for assembly
	return transactions.NewSMSAssembly()
}

// NewHoldDMS returns new instance to new HoldDMS structure
func (ob *Builder) NewHoldDMS() *transactions.HoldDMSAssembly {
	// Get new prepared sms structure for assembly
	return transactions.NewHoldDMSAssembly()
}

// NewChargeDMS returns new instance to new ChargeDMS  structure
func (ob *Builder) NewChargeDMS() *transactions.ChargeDMSAssembly {
	// Get new prepared sms structure for assembly
	return transactions.NewChargeDMSAssembly()
}

// NewCancel returns new instance to new ChargeDMS  structure
func (ob *Builder) NewCancel() *transactions.CancelAssembly {
	// Get new prepared sms structure for assembly
	return transactions.NewCancelAssembly()
}

// NewMOTOSMS returns new instance to new MOTO SMS structure
func (ob *Builder) NewMOTOSMS() *transactions.MOTOAssembly {
	// Get new prepared sms structure for assembly
	return transactions.NewMOTOSMSAssembly()
}

// NewMOTODMS returns new instance to new MOTO DMS structure
func (ob *Builder) NewMOTODMS() *transactions.MOTOAssembly {
	// Get new prepared sms structure for assembly
	return transactions.NewMOTODMSAssembly()
}

// NewRecurrentSMS returns new instance to new recurrent SMS structure
func (ob *Builder) NewRecurrentSMS() *transactions.RecurrentAssembly {
	// Get new prepared sms structure for assembly
	return transactions.NewRecurrentSMSAssembly()
}

// NewRecurrentDMS returns new instance to new recurrent DMS structure
func (ob *Builder) NewRecurrentDMS() *transactions.RecurrentAssembly {
	// Get new prepared sms structure for assembly
	return transactions.NewRecurrentDMSAssembly()
}

// NewRefund returns new instance to new refund structure
func (ob *Builder) NewRefund() *transactions.RefundAssembly {
	// Get new prepared sms structure for assembly
	return transactions.NewRefundAssembly()
}

// NewReversal returns new instance to new reversal structure
func (ob *Builder) NewReversal() *transactions.ReversalAssembly {
	// Get new prepared sms structure for assembly
	return transactions.NewReversalAssembly()
}

// Exploring Past Payments builders

// NewGetStatus returns new instance to new status structure
// allows to get status of past transaction in Transact Pro system
func (ob *Builder) NewGetStatus() *exploring_requests.StatusAssembly {
	return exploring_requests.NewStatusAssembly()
}
