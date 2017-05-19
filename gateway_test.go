package tprogateway

import (
	"testing"
	"bitbucket.transactpro.lv/tls/gw3-go-client/builder"
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
	if err != nil {
		t.Error(err)
	}
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
	apiGC, err := NewGatewayClient(ma.AccID, ma.SecKey)
	if err != nil {
		t.Error(err)
	}
	apiGC.API.Uri = "https://proxy.payment-tpro.co.uk"
	if apiGC.API.Uri == dAPIBaseUri {
		t.Error("API uri not changed")
	}

	apiGC.API.Version = "1.0"
	if apiGC.API.Uri == dAPIVersion {
		t.Error("API Version not changed")
	}
}

func TestNewOperation(t *testing.T) {
	sms := gc.NewOp().SMS()
	sms.PaymentMethod.Pan = "5262482284416445"
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

	newReq, err := gc.NewRequest(builder.SMS, sms)
	if err != nil {
		t.Error(err)
	}

	if newReq == nil {
		t.Error("HTTP NewRequest structure is empty.")
	}
}

func TestDetermineAPIActionUriError(t *testing.T) {
	apiGC, errGC := NewGatewayClient(ma.AccID, ma.SecKey)
	if errGC != nil {
		t.Error(errGC)
	}

	apiGC.API.Uri = ""
	_, _, err := determineAPIAction(apiGC.API.Uri, apiGC.API.Version, builder.SMS)
	if err == nil {
		t.Error(err)
	}
}

func TestDetermineAPIActionVersionError(t *testing.T) {
	apiGC, errGC := NewGatewayClient(ma.AccID, ma.SecKey)
	if errGC != nil {
		t.Error(errGC)
	}

	apiGC.API.Version = ""
	_, _, err := determineAPIAction(apiGC.API.Uri, apiGC.API.Version, builder.SMS)
	if err == nil {
		t.Error(err)
	}
}

func TestDetermineAPIActionHttpMethodError(t *testing.T) {
	apiGC, errGC := NewGatewayClient(ma.AccID, ma.SecKey)
	if errGC != nil {
		t.Error(errGC)
	}

	var WRONG_OP builder.OperationType = "WRONG_OP"

	_, _, err := determineAPIAction(apiGC.API.Uri, apiGC.API.Version, WRONG_OP)
	if err == nil {
		t.Error(err)
	}
}

func TestSendRequest(t *testing.T)  {
	gc, err := NewGatewayClient(ma.AccID, ma.SecKey)

	sms := gc.NewOp().SMS()
	sms.PaymentMethod.Pan = "5262482284416445"
	sms.PaymentMethod.ExpMmYy = "12/20"
	sms.PaymentMethod.Cvv = "403"
	sms.Money.Amount = 300
	sms.Money.Currency = "EUR"

	newReq, err := gc.NewRequest(builder.SMS, sms)
	if err != nil {
		t.Error(err)
	}

	resp, respErr := gc.SendRequest(newReq)
	if respErr != nil {
		t.Error(respErr)
	}

	if resp == nil {
		t.Error("Parsed response is empty")
	}

}
