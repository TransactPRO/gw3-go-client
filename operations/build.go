package operations

import (
	"github.com/TransactPRO/gw3-go-client/operations/exploring"
	"github.com/TransactPRO/gw3-go-client/operations/transactions"
)

// Builder operation structure builder for specific request
type Builder struct{}

/*

	Transaction Types builders

*/

// @TODO In Transaction assembly refactor some data sets, cos not all bundles used for each transaction. Some used few fields in groups

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

// NewCredit returns new instance to new Credit structure
func (ob *Builder) NewCredit() *transactions.CreditAssembly {
	// Get new prepared Credit structure for assembly
	return transactions.NewCreditAssembly()
}

// NewP2P returns new instance to new P2P structure
func (ob *Builder) NewP2P() *transactions.P2PAssembly {
	// Get new prepared P2P structure for assembly
	return transactions.NewP2PAssembly()
}

// NewInitRecurrentSMS returns new instance to new Init Recurrent SMS structure
func (ob *Builder) NewInitRecurrentSMS() *transactions.InitRecurrentSMSAssembly {
	// Get new prepared sms structure for assembly
	return transactions.NewInitRecurrentSMSAssembly()
}

// NewRecurrentSMS returns new instance to new recurrent SMS structure
func (ob *Builder) NewRecurrentSMS() *transactions.RecurrentAssembly {
	// Get new prepared sms structure for assembly
	return transactions.NewRecurrentSMSAssembly()
}

// NewInitRecurrentDMS returns new instance to new Init Recurrent DMS structure
func (ob *Builder) NewInitRecurrentDMS() *transactions.InitRecurrentDMSAssembly {
	// Get new prepared sms structure for assembly
	return transactions.NewInitRecurrentDMSAssembly()
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

/*

	Exploring Past Payments builders

*/

// NewGetStatus returns new instance to new status structure
// allows to get status of past transaction in Transact Pro system
func (ob *Builder) NewGetStatus() *exploring.ExploreTransactionAssembly {
	return exploring.NewStatusAssembly()
}

// NewGetResult returns new instance to new status structure
// allows to get status of past transaction in Transact Pro system
func (ob *Builder) NewGetResult() *exploring.ExploreTransactionAssembly {
	return exploring.NewResultAssembly()
}

// NewGetHistory returns new instance to new status structure
// allows to get status of past transaction in Transact Pro system
func (ob *Builder) NewGetHistory() *exploring.ExploreTransactionAssembly {
	return exploring.NewHistoryAssembly()
}

// NewGetRecurrents returns new instance to new status structure
// allows to get status of past transaction in Transact Pro system
func (ob *Builder) NewGetRecurrents() *exploring.ExploreTransactionAssembly {
	return exploring.NewRecurrentsAssembly()
}

// NewGetRefunds returns new instance to new status structure
// allows to get status of past transaction in Transact Pro system
func (ob *Builder) NewGetRefunds() *exploring.ExploreTransactionAssembly {
	return exploring.NewHistoryAssembly()
}
