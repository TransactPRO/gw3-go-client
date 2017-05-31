package tprogateway

import (
	"testing"

	"bitbucket.transactpro.lv/tls/gw3-go-client/operations/transactions"
)

// @TODO Mock\Stub responses from gateway when calling HTTP requests\response

func TestNewGatewayClient(t *testing.T) {
	_, err := NewGatewayClient(42, "SecKey")
	if err != nil {
		t.Error(err)
		return
	}
}

func TestNewGatewayClientIncorrectAccountID(t *testing.T) {
	_, err := NewGatewayClient(0, "SecKey")
	if err == nil {
		t.Error(err)
		return
	}
}

func TestNewGatewayClientIncorrectSecretKey(t *testing.T) {
	_, err := NewGatewayClient(42, "")
	if err == nil {
		t.Error(err)
		return
	}
}

func TestNewGatewayClientRedefineDefaultAPISettings(t *testing.T) {
	gc, err := NewGatewayClient(42, "SecKey")
	if err != nil {
		t.Error(err)
		return
	}

	gc.API.BaseURI = "https://proxy.payment-tpro.co.uk"
	if gc.API.BaseURI == dAPIBaseURI {
		t.Error("GatewayClient API uri not changed")
		return
	}

	gc.API.Version = "1.0"
	if gc.API.BaseURI == dAPIVersion {
		t.Error("GatewayClient API Version not changed")
		return
	}
}

func TestNewOperation(t *testing.T) {
	gc, err := NewGatewayClient(42, "SecKey")
	if err != nil {
		t.Error(err)
		return
	}

	sms := gc.OperationBuilder().NewSms()
	sms.PaymentMethod.Pan = "5262482284416445"
	sms.PaymentMethod.ExpMmYy = "12/20"
	sms.PaymentMethod.Cvv = "123"
	sms.Money.Amount = 300
	sms.Money.Currency = "EUR"

	sms.System.UserIP = "xxx.0.0.1"
	sms.System.XForwardedFor = "xxx.66.33.12"

	if sms == transactions.NewSMSAssembly() {
		t.Error("SMS data system structure not changed")
		return
	}
}
