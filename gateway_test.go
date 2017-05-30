package tprogateway

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"strconv"
	"testing"
	"time"

	"bitbucket.transactpro.lv/tls/gw3-go-client/structures"
)

// @TODO Mock\Stub responses from gateway when calling HTTP requests\response

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
		return
	}
}

func TestNewGatewayClientIncorrectAccountID(t *testing.T) {
	_, err := NewGatewayClient(0, caa.SecKey)
	if err == nil {
		t.Error(err)
		return
	}
}

func TestNewGatewayClientIncorrectSecretKey(t *testing.T) {
	_, err := NewGatewayClient(caa.AccID, "")
	if err == nil {
		t.Error(err)
		return
	}
}

func TestNewGatewayClientRedefineDefaultAPISettings(t *testing.T) {
	apiGC, err := NewGatewayClient(caa.AccID, caa.SecKey)
	if err != nil {
		t.Error(err)
		return
	}

	apiGC.API.BaseURI = "https://proxy.payment-tpro.co.uk"
	if apiGC.API.BaseURI == dAPIBaseURI {
		t.Error("GatewayClient API uri not changed")
		return
	}

	apiGC.API.Version = "1.0"
	if apiGC.API.BaseURI == dAPIVersion {
		t.Error("GatewayClient API Version not changed")
		return
	}
}

func TestNewOperation(t *testing.T) {
	gc, err := NewGatewayClient(caa.AccID, caa.SecKey)
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

	if sms.System.UserIP != "xxx.0.0.1" && sms.System.XForwardedFor != "xxx.66.33.12" {
		t.Error("SMS data system structure not changed")
		return
	}
}

func TestNewRequest(t *testing.T) {
	gc, err := NewGatewayClient(caa.AccID, caa.SecKey)
	if err != nil {
		t.Error(err)
		return
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
		return
	}

	if resp == nil {
		t.Error("HTTP NewRequest response is empty.")
		return
	}
}

func TestSendRequest(t *testing.T) {
	correctGc, errGc := NewGatewayClient(caa.AccID, caa.SecKey)
	if errGc != nil {
		t.Error(errGc)
		return
	}

	// Create some random values for our request
	newSource := rand.NewSource(time.Now().UnixNano())
	newRand := rand.New(newSource)

	sms := correctGc.OperationBuilder().NewSms()
	sms.CommandData.FormID = strconv.Itoa(newRand.Intn(100500))
	sms.CommandData.TerminalMID = cac.TerminalMID
	sms.GeneralData.OrderData.MerchantTransactionID = fmt.Sprintf("TestTranID:%d", newRand.Intn(rand.Int()))
	sms.GeneralData.OrderData.OrderDescription = "Gopher Gufer ordering goods"
	sms.GeneralData.OrderData.OrderID = fmt.Sprintf("TestOrderID:%d", newRand.Intn(rand.Int()))
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
		return
	}

	if resp == nil {
		t.Error("Parsed response is empty")
		return
	}

	// @TODO Debug print
	fmt.Println(fmt.Sprintf("RAW %+v", resp))
}

func TestSendRequestSMSWithParse(t *testing.T) {
	// HA
	//correctGc, errGc := NewGatewayClient(28, "94dYyLTjVGM2aXSh6Aq1QcCHnPNev38p")
	correctGc, errGc := NewGatewayClient(caa.AccID, caa.SecKey)
	if errGc != nil {
		t.Error(errGc)
		return
	}

	//correctGc.API.BaseURI = "http://uriel.ha.fpngw3.env"

	// Create some random values for our request
	newSource := rand.NewSource(time.Now().UnixNano())
	newRand := rand.New(newSource)

	sms := correctGc.OperationBuilder().NewSms()
	sms.CommandData.FormID = strconv.Itoa(newRand.Intn(100500))
	sms.CommandData.TerminalMID = "30"
	sms.GeneralData.OrderData.MerchantTransactionID = fmt.Sprintf("TestTranID:%d", newRand.Intn(rand.Int()))
	sms.GeneralData.OrderData.OrderDescription = "Gopher Gufer ordering goods"
	sms.GeneralData.OrderData.OrderID = fmt.Sprintf("TestOrderID:%d", newRand.Intn(rand.Int()))
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
		return
	}

	if resp == nil {
		t.Error("Parsed response is empty")
		return
	}
	// @TODO Debug print
	fmt.Println(fmt.Sprintf("RAW %+v", resp))

	parsedRes, parseErr := correctGc.ParseResponse(resp, structures.SMS)
	if parseErr != nil {
		t.Error(parseErr)
		return
	}

	if parsedRes == nil {
		t.Error("Parsed responses of SMS is empty, problem in method ParseResponse ")
		return
	}

	// @TODO Debug print
	fmt.Println(fmt.Sprintf("%+v", parsedRes))
}

func TestSendRequestUnauthorizedSMSWithParse(t *testing.T) {
	gc, err := NewGatewayClient(caa.AccID, "SomeWrongKeyOrOldKeyHere")
	if err != nil {
		t.Error(err)
		return
	}

	newOp := gc.OperationBuilder()
	sms := newOp.NewSms()
	sms.PaymentMethod.Pan = "5262482284416445"
	sms.PaymentMethod.ExpMmYy = "12/20"
	sms.PaymentMethod.Cvv = "403"
	sms.Money.Amount = 6699
	sms.Money.Currency = "EUR"

	resp, err := gc.NewRequest(sms)
	if err != nil {
		t.Error(err)
		return
	}

	if resp == nil {
		t.Error("HTTP NewRequest response is empty.")
		return
	}

	defer resp.Body.Close()

	body, bodyErr := ioutil.ReadAll(resp.Body)
	if bodyErr != nil {
		t.Error(fmt.Sprintf("Failed to read received body: %s ", bodyErr.Error()))
		return
	}

	var gwResp structures.UnauthorizedResponse

	parseErr := json.Unmarshal(body, &gwResp)
	if parseErr != nil {
		if bodyErr != nil {
			t.Error(fmt.Sprintf("Failed to unmarshal received body: %s ", bodyErr.Error()))
			return
		}
		t.Error("Failed to unmarshal received body, body error unkown")
		return
	}

	if gwResp.Msg != "Unauthorized" && gwResp.Status != 401 {
		t.Error("Incorect parse of unauthorized response from Transact Pro gateway")
		return
	}

	// @TODO Debug print
	fmt.Println(fmt.Sprintf("%+v", gwResp))

}

func TestSendRequestDMS(t *testing.T) {
	correctGc, errGc := NewGatewayClient(caa.AccID, caa.SecKey)
	if errGc != nil {
		t.Error(errGc)
		return
	}

	// Create some random values for our request
	newSource := rand.NewSource(time.Now().UnixNano())
	newRand := rand.New(newSource)
	tranAmount := newRand.Intn(500)

	opBuild := correctGc.OperationBuilder()

	holdDMS := opBuild.NewHoldDMS()
	holdDMS.GeneralData.OrderData.OrderDescription = "Gopher Gufer DO HOLD DMS"
	holdDMS.GeneralData.OrderData.OrderID = fmt.Sprintf("TestOrderID:%d", newRand.Intn(rand.Int()))
	holdDMS.GeneralData.CustomerData.BillingAddress.City = "Riga"

	holdDMS.PaymentMethod.Pan = "5262482284416445"
	holdDMS.PaymentMethod.ExpMmYy = "12/20"
	holdDMS.PaymentMethod.Cvv = "403"

	holdDMS.Money.Amount = tranAmount
	holdDMS.Money.Currency = "EUR"

	holdDMS.System.UserIP = "127.0.0.1"
	holdDMS.System.XForwardedFor = "127.0.0.1"

	respHold, respHoldErr := correctGc.NewRequest(holdDMS)
	if respHoldErr != nil {
		t.Error(respHoldErr)
		return
	}

	if respHold == nil {
		t.Error("DMS Hold parsed response is empty")
		return
	}

	parsedHoldRes, parseErr := correctGc.ParseResponse(respHold, structures.DMSHold)
	if parseErr != nil {
		t.Error(parseErr)
		return
	}

	// @TODO Debug print
	fmt.Println(fmt.Sprintf("%+v", parsedHoldRes))

	chargeDMS := opBuild.NewChargeDMS()
	chargeDMS.CommandData.GWTransactionID = parsedHoldRes.(structures.TransactionResponse).GateWay.GatewayTransactionID
	chargeDMS.Money.Amount = tranAmount
	chargeDMS.System.UserIP = "127.0.0.1"
	chargeDMS.System.XForwardedFor = "127.0.0.1"

	respCharge, respChargeErr := correctGc.NewRequest(chargeDMS)
	if respChargeErr != nil {
		t.Error(respChargeErr)
		return
	}

	if respCharge == nil {
		t.Error("DMS Charge parsed response is empty")
		return
	}

	parsedCharge, parseChargeErr := correctGc.ParseResponse(respCharge, structures.DMSCharge)
	if parseChargeErr != nil {
		t.Error(parseChargeErr)
		return
	}

	// @TODO Debug print
	fmt.Println(fmt.Sprintf("%+v", parsedCharge))

	cancel := opBuild.NewCancel()
	cancel.CommandData.GWTransactionID = parsedHoldRes.(structures.TransactionResponse).GateWay.GatewayTransactionID
	cancel.System.UserIP = "127.0.0.1"
	cancel.System.XForwardedFor = "127.0.0.1"

	respCancel, respCancelErr := correctGc.NewRequest(cancel)
	if respCancelErr != nil {
		t.Error(respCancelErr)
		return
	}

	if respCancel == nil {
		t.Error("Cancel parsed response is empty")
		return
	}

	parsedCancel, parseCancelErr := correctGc.ParseResponse(respCancel, structures.CANCEL)
	if parseCancelErr != nil {
		t.Error(parseCancelErr)
		return
	}

	// @TODO Debug print
	fmt.Println(fmt.Sprintf("%+v", parsedCancel))
}
