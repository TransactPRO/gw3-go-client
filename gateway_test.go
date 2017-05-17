package tprogateway

import (
	"testing"
)

// ma, Merchant authorization configuration
var ma *testCorrectMerchantAuth
type testCorrectMerchantAuth struct {
	AccID  int
	SecKey string
}

func init()  {
	ma = &testCorrectMerchantAuth{AccID:22, SecKey:"rg342QZSUaWzKHoCc5slyMGdAITk9LfR"}

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
	gc, err := NewGatewayClient(ma.AccID, ma.SecKey)
	ifErrHandleAsDefault(err, t)

	gc.API.Uri = "https://proxy.payment-tpro.co.uk"
	if gc.API.Uri == dAPIUri {
		t.Error("API uri not changed")
	}

	gc.API.Version = "1.0"
	if gc.API.Uri == dAPIVersion {
		t.Error("API Version not changed")
	}
}

func TestNewGatewayClientOperation(t *testing.T)  {
	gc, err := NewGatewayClient(ma.AccID, ma.SecKey)
	ifErrHandleAsDefault(err, t)

	sms := gc.operation.NewSMS()

	sms.Data.PaymentMethod.Pan = "2379183712983"
	sms.Data.PaymentMethod.ExpMmYy = "12/20"
	sms.Data.PaymentMethod.Cvv = "123"

	sms.Data.System.UserIP = "xxx.0.0.1"
	sms.Data.System.XForwardedFor = "xxx.66.33.12"

	if sms.Data.System.UserIP != "xxx.0.0.1" && sms.Data.System.XForwardedFor != "xxx.66.33.12" {
		t.Error("System structure not changed")
	}
}

// ifErrHandle, default err handler in unit test
func ifErrHandleAsDefault(err error, t *testing.T)  {
	if err != nil {
		t.Error(err)
	}
}
