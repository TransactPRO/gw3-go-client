package operations

import "bitbucket.transactpro.lv/tls/gw3-go-client/operations/transaction"

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

// NewCancel returns new instance to new ChargeDMS  structure
func (ob *Builder) NewCancel() *transaction.CancelAssembly {
	// Get new prepared sms structure for assembly
	return transaction.NewCancelAssembly()
}

// NewMOTOSMS returns new instance to new MOTO SMS structure
func (ob *Builder) NewMOTOSMS() *transaction.MOTOAssembly {
	// Get new prepared sms structure for assembly
	return transaction.NewMOTOSMSAssembly()
}

// NewMOTODMS returns new instance to new MOTO DMS structure
func (ob *Builder) NewMOTODMS() *transaction.MOTOAssembly {
	// Get new prepared sms structure for assembly
	return transaction.NewMOTODMSAssembly()
}

// NewRecurrentSMS returns new instance to new recurrent SMS structure
func (ob *Builder) NewRecurrentSMS() *transaction.RecurrentAssembly {
	// Get new prepared sms structure for assembly
	return transaction.NewRecurrentSMSAssembly()
}

// NewRecurrentDMS returns new instance to new recurrent DMS structure
func (ob *Builder) NewRecurrentDMS() *transaction.RecurrentAssembly {
	// Get new prepared sms structure for assembly
	return transaction.NewRecurrentDMSAssembly()
}

// NewRefund returns new instance to new refund structure
func (ob *Builder) NewRefund() *transaction.RefundAssembly {
	// Get new prepared sms structure for assembly
	return transaction.NewRefundAssembly()
}

// NewReversal returns new instance to new reversal structure
func (ob *Builder) NewReversal() *transaction.ReversalAssembly {
	// Get new prepared sms structure for assembly
	return transaction.NewReversalAssembly()
}
