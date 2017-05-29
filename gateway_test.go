package tprogateway

import (
	"fmt"
	"math/rand"
	"testing"
	"time"

	"bitbucket.transactpro.lv/tls/gw3-go-client/operations"
)

// caa, Merchant authorization configuration
var (
	// Merchant auth structure
	caa *testCorrectAccAuth
	// Account config
	cac *testCorrectAccConfig
)

type testCorrectAccAuth struct {
	AccID  int
	SecKey string
}

type testCorrectAccConfig struct {
	TerminalMID string
}

func init() {
	caa = &testCorrectAccAuth{
		AccID: 40, SecKey: "Xh1PvJQiTtFCU0uxqwBE9pzI8eks3mSa"}
	cac = &testCorrectAccConfig{
		TerminalMID: "590c699593ac4"}
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

func TestNewGatewayClientRedefineDefaultAPISettings(t *testing.T) {
	apiGC, err := NewGatewayClient(caa.AccID, caa.SecKey)
	if err != nil {
		t.Error(err)
	}

	apiGC.API.BaseURI = "https://proxy.payment-tpro.co.uk"
	if apiGC.API.BaseURI == dAPIBaseURI {
		t.Error("GatewayClient API uri not changed")
	}

	apiGC.API.Version = "1.0"
	if apiGC.API.BaseURI == dAPIVersion {
		t.Error("GatewayClient API Version not changed")
	}
}

func TestNewOperation(t *testing.T) {
	gc, err := NewGatewayClient(caa.AccID, caa.SecKey)
	if err != nil {
		t.Error(err)
	}

	sms := gc.OperationBuilder().NewSms()
	sms.PaymentMethod.Pan = "5262482284416445"
	sms.PaymentMethod.ExpMmYy = "12/20"
	sms.PaymentMethod.Cvv = "123"
	sms.Money.Amount = 300
	sms.Money.Currency = "EUR"

	sms.System.UserIP = "xxx.0.0.1"
	sms.System.XForwardedFor = "xxx.66.33.12"

	if sms.System.UserIP != "xxx.0.0.1" && sms.System.XForwardedFor != "xxx.66.33.12" {
		t.Error("SMS data system structure not changed")
	}
}

func TestNewRequest(t *testing.T) {
	gc, err := NewGatewayClient(caa.AccID, caa.SecKey)
	if err != nil {
		t.Error(err)
	}

	newOp := gc.OperationBuilder()
	sms := newOp.NewSms()
	sms.PaymentMethod.Pan = "5262482284416445"
	sms.PaymentMethod.ExpMmYy = "12/20"
	sms.PaymentMethod.Cvv = "403"
	sms.Money.Amount = 300
	sms.Money.Currency = "EUR"

	resp, err := gc.NewRequest(sms)
	if err != nil {
		t.Error(err)
	}

	if resp == nil {
		t.Error("HTTP NewRequest response is empty.")
	}
}

func TestSendRequest(t *testing.T) {
	correctGc, errGc := NewGatewayClient(caa.AccID, caa.SecKey)
	if errGc != nil {
		t.Error(errGc)
	}

	// Create some random values for our request
	newSource := rand.NewSource(time.Now().UnixNano())
	newRand := rand.New(newSource)

	sms := correctGc.OperationBuilder().NewSms()
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

	resp, reqErr := correctGc.NewRequest(sms)
	if reqErr != nil {
		t.Error(reqErr)
	}

	if resp == nil {
		t.Error("Parsed response is empty")
	}
}

// @TODO cover more code
func TestSendRequestSMSWithParse(t *testing.T) {
	correctGc, errGc := NewGatewayClient(caa.AccID, caa.SecKey)
	if errGc != nil {
		t.Error(errGc)
	}

	// Create some random values for our request
	newSource := rand.NewSource(time.Now().UnixNano())
	newRand := rand.New(newSource)

	sms := correctGc.OperationBuilder().NewSms()
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

	resp, reqErr := correctGc.NewRequest(sms)
	if reqErr != nil {
		t.Error(reqErr)
	}

	if resp == nil {
		t.Error("Parsed response is empty")
	}

	parsedRes, parseErr := correctGc.ParseResponse(resp, operations.SMS)
	if parseErr != nil {
		t.Error(parseErr)
	}

	if parsedRes == nil {
		t.Error("Parsed responses of SMS is empty, problem in method ParseResponse ")
	}

	fmt.Println(fmt.Sprintf("%+v", parsedRes))
}
