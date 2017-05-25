package tprogateway

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
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
		AccID: 22, SecKey: "rg342QZSUaWzKHoCc5slyMGdAITk9LfR"}
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

	sms := gc.NewOperation().SMS()
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

	sms := gc.NewOperation().SMS()
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

	sms := correctGc.NewOperation().SMS()
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


// @TODO cover more code with test cases like old one

//func TestDetermineURLErrorBaseURI(t *testing.T) {
//	apiGC, errGC := NewGatewayClient(caa.AccID, caa.SecKey)
//	if errGC != nil {
//		t.Error(errGC)
//	}
//
//	var reqB builder.RequestBuilder
//
//	apiGC.API.BaseURI = ""
//	url, err := determineURL(apiGC, &reqB)
//	if err == nil {
//		t.Error(err)
//	}
//
//	if url == "" {
//		t.Error("URL returned empty")
//	}
//}
//
//func TestDetermineURLErrorVersion(t *testing.T) {
//	apiGC, errGC := NewGatewayClient(caa.AccID, caa.SecKey)
//	if errGC != nil {
//		t.Error(errGC)
//	}
//
//	var reqB builder.RequestBuilder
//
//	apiGC.API.Version = ""
//	url, err := determineURL(apiGC, &reqB)
//	if err == nil {
//		t.Error(err)
//	}
//
//	if url == "" {
//		t.Error("URL returned empty")
//	}
//}

//func TestDetermineURLErrorEmptyURL(t *testing.T) {
//	apiGC, errGC := NewGatewayClient(caa.AccID, caa.SecKey)
//	if errGC != nil {
//		t.Error(errGC)
//	}
//
//	var reqB builder.RequestBuilder
//
//	url, err := determineURL(apiGC, &reqB)
//	if err == nil {
//		t.Error(err)
//	}
//
//	if url != "" {
//		t.Error("URL formed with empty operation type. Problem in request builder")
//	}
//}

//
//func TestParseResponseError(t *testing.T) {
//	apiGC, errGC := NewGatewayClient(caa.AccID, caa.SecKey)
//	if errGC != nil {
//		t.Error(errGC)
//	}
//
//	apiGC.lastReqData.httpMethod = "POST"
//	apiGC.lastReqData.httpEndpoint = "http://unit.pay.com/v66.0/sms"
//
//	handler := func(w http.ResponseWriter, r *http.Request) {
//		io.WriteString(w, "TEST SMS UNIT")
//	}
//
//	req := httptest.NewRequest(apiGC.lastReqData.httpMethod, apiGC.lastReqData.httpEndpoint, nil)
//	w := httptest.NewRecorder()
//	handler(w, req)
//
//	tResp := w.Result()
//
//	_, err := parseResponse(apiGC, tResp)
//	if err == nil {
//		t.Error("GatewayClinet parse response didn't return error of parsing http response struct")
//	}
//}
//
//func TestParseResponseSMS(t *testing.T) {
//	apiGC, errGC := NewGatewayClient(caa.AccID, caa.SecKey)
//	if errGC != nil {
//		t.Error(errGC)
//	}
//
//	apiGC.lastReqData.operation = builder.SMS
//	apiGC.lastReqData.httpMethod = "POST"
//	apiGC.lastReqData.httpEndpoint = "http://unit.pay.com/v66.0/sms"
//
//	handler := func(w http.ResponseWriter, r *http.Request) {
//		io.WriteString(w, "TEST SMS UNIT")
//	}
//
//	req := httptest.NewRequest(apiGC.lastReqData.httpMethod, apiGC.lastReqData.httpEndpoint, nil)
//	w := httptest.NewRecorder()
//	handler(w, req)
//
//	tResp := w.Result()
//
//	_, err := parseResponse(apiGC, tResp)
//	if err == nil {
//		t.Error("GatewayClinet parse response didn't return error of parsing http response struct")
//	}
//}

