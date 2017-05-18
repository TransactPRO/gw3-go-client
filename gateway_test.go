package tprogateway

import (
	"testing"
	"fmt"
)

// ma, Merchant authorization configuration
var (
	// Merchant auth structure
	ma *testCorrectMerchantAuth
	// Correct instance of GatewayClient
	gc *GatewayClient
)

type testCorrectMerchantAuth struct {
	AccID  int
	SecKey string
}

func init()  {
	ma = &testCorrectMerchantAuth{AccID:22, SecKey:"rg342QZSUaWzKHoCc5slyMGdAITk9LfR"}
	gc, _ = NewGatewayClient(ma.AccID, ma.SecKey)
}

func TestNewGatewayClient(t *testing.T) {
	_, err := NewGatewayClient(ma.AccID, ma.SecKey)
	ifErrHandleAsDefault(err, t)
}

func TestNewGatewayClientIncorrectAccountID(t *testing.T) {
	_, err := NewGatewayClient(0, ma.SecKey)
	if err == nil {
		t.Error(err)
	}
}

func TestNewGatewayClientIncorrectSecretKey(t *testing.T) {
	_, err := NewGatewayClient(ma.AccID, "")
	if err == nil {
		t.Error(err)
	}
}

func TestNewGatewayClientRedefineDefaultAPISettings(t *testing.T)  {
	gc.API.Uri = "https://proxy.payment-tpro.co.uk"
	if gc.API.Uri == dAPIUri {
		t.Error("API uri not changed")
	}

	gc.API.Version = "1.0"
	if gc.API.Uri == dAPIVersion {
		t.Error("API Version not changed")
	}
}

func TestNewOperation(t *testing.T) {
	sms := gc.NewOp().SMS()
	sms.PaymentMethod.Pan = "2379183712983"
	sms.PaymentMethod.ExpMmYy = "12/20"
	sms.PaymentMethod.Cvv = "123"
	sms.Money.Amount = 300
	sms.Money.Currency = "EUR"

	sms.System.UserIP = "xxx.0.0.1"
	sms.System.XForwardedFor = "xxx.66.33.12"

	if sms.System.UserIP != "xxx.0.0.1" && sms.System.XForwardedFor != "xxx.66.33.12" {
		t.Error("System structure not changed")
	}
}

func TestNewRequest(t *testing.T) {
	sms := gc.NewOp().SMS()
	sms.PaymentMethod.Pan = "5262482284416445"
	sms.PaymentMethod.ExpMmYy = "12/20"
	sms.PaymentMethod.Cvv = "403"
	sms.Money.Amount = 300
	sms.Money.Currency = "EUR"

	_, err := gc.NewRequest(sms)
	if err != nil {
		t.Error(err)
	}
}

// ifErrHandle, default err handler in unit test
func ifErrHandleAsDefault(err error, t *testing.T)  {
	if err != nil {
		t.Error(err)
	}
}
