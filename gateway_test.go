package tprogateway

import (
	"testing"
	"bitbucket.transactpro.lv/tls/gw3-go-client/builder"
	"fmt"
	"math/rand"
	"time"
)

// caa, Merchant authorization configuration
var (
	// Merchant auth structure
	caa *testCorrectAccAuth
	// Accouunt config
	cac *testCorrectAccConfig
	// Correct instance of GatewayClient
	gc *GatewayClient
)

type testCorrectAccAuth struct {
	AccID  int
	SecKey string
}

type testCorrectAccConfig struct {
	TerminalMID string
}

func init()  {
	caa = &testCorrectAccAuth{
		AccID:22, SecKey:"rg342QZSUaWzKHoCc5slyMGdAITk9LfR"}
	cac = &testCorrectAccConfig{
		TerminalMID: "590c699593ac4"}
	gc, _ = NewGatewayClient(caa.AccID, caa.SecKey)
}

func TestNewGatewayClient(t *testing.T) {
	_, err := NewGatewayClient(caa.AccID, caa.SecKey)
	if err != nil {
		t.Error(err)
	}
}

func TestNewGatewayClientIncorrectAccountID(t *testing.T) {
	_, err := NewGatewayClient(0, caa.SecKey)
	if err == nil {
		t.Error(err)
	}
}

func TestNewGatewayClientIncorrectSecretKey(t *testing.T) {
	_, err := NewGatewayClient(caa.AccID, "")
	if err == nil {
		t.Error(err)
	}
}

func TestNewGatewayClientRedefineDefaultAPISettings(t *testing.T)  {
	apiGC, err := NewGatewayClient(caa.AccID, caa.SecKey)
	if err != nil {
		t.Error(err)
	}
	apiGC.API.BaseUri = "https://proxy.payment-tpro.co.uk"
	if apiGC.API.BaseUri == dAPIBaseUri {
		t.Error("API uri not changed")
	}

	apiGC.API.Version = "1.0"
	if apiGC.API.BaseUri == dAPIVersion {
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
	apiGC, errGC := NewGatewayClient(caa.AccID, caa.SecKey)
	if errGC != nil {
		t.Error(errGC)
	}

	apiGC.API.BaseUri = ""
	apiGC.lastReqData.operation = builder.SMS
	err := determineAPIAction(apiGC)
	if err == nil {
		t.Error(err)
	}
}

func TestDetermineAPIActionVersionError(t *testing.T) {
	apiGC, errGC := NewGatewayClient(caa.AccID, caa.SecKey)
	if errGC != nil {
		t.Error(errGC)
	}

	apiGC.API.Version = ""
	apiGC.lastReqData.operation = builder.SMS
	err := determineAPIAction(apiGC)
	if err == nil {
		t.Error(err)
	}
}

func TestDetermineAPIActionHttpMethodError(t *testing.T) {
	apiGC, errGC := NewGatewayClient(caa.AccID, caa.SecKey)
	if errGC != nil {
		t.Error(errGC)
	}

	var WRONG_OP builder.OperationType = "WRONG_OP"
	apiGC.lastReqData.operation = WRONG_OP

	err := determineAPIAction(apiGC)
	if err == nil {
		t.Error(err)
	}
}

func TestSendRequest(t *testing.T)  {
	correctGc, errGc := NewGatewayClient(caa.AccID, caa.SecKey)
	if errGc != nil {
		t.Error(errGc)
	}

	sms := correctGc.NewOp().SMS()
	newSource := rand.NewSource(time.Now().UnixNano())
	newRand := rand.New(newSource)
	sms.CommandData.FormID = fmt.Sprintf("%d", newRand.Intn(100500))
	sms.CommandData.TerminalMID = cac.TerminalMID
	sms.GeneralData.OrderData.MerchantTransactionID = fmt.Sprintf("TestTranID:%d", newRand.Intn(rand.Int()))
	sms.GeneralData.OrderData.OrderDescription = "Gopher Gufer ordering goods"
	sms.GeneralData.OrderData.OrderID = fmt.Sprintf("TestOrderID%d", newRand.Intn(rand.Int()))
	sms.GeneralData.CustomerData.Email = "some@email.com"
	sms.GeneralData.CustomerData.BillingAddress.City = "Riga"
	sms.PaymentMethod.Pan = "5262482284416445"
	sms.PaymentMethod.ExpMmYy = "12/20"
	sms.PaymentMethod.Cvv = "403"
	sms.Money.Amount = newRand.Intn(500)
	sms.Money.Currency = "EUR"
	sms.System.UserIP = "127.0.0.1"
	sms.System.XForwardedFor = "127.0.0.1"

	req, reqErr := correctGc.NewRequest(builder.SMS, sms)
	if reqErr != nil {
		t.Error(reqErr)
	}

	resp, respErr := correctGc.SendRequest(req)
	if respErr != nil {
		t.Error(respErr)
	}

	if resp == nil {
		t.Error("Parsed response is empty")
	}
}
