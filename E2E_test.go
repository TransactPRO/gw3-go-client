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

// testAuth is a pointer to GatewayClient authData structure with merchant cridtetionals
var testAuth *authData

func init() {
	// testAuth can be changed to your testing merchant auth data
	testAuth = &authData{
		AccountID: 40,
		SecretKey: "Xh1PvJQiTtFCU0uxqwBE9pzI8eks3mSa",
	}
}

func TestSendRequest(t *testing.T) {
	gc, gcErr := NewGatewayClient(testAuth.AccountID, testAuth.SecretKey)
	if gcErr != nil {
		t.Error(gcErr)
		return
	}

	// Create some random values for our request
	newSource := rand.NewSource(time.Now().UnixNano())
	newRand := rand.New(newSource)

	sms := gc.OperationBuilder().NewSms()
	sms.CommandData.FormID = strconv.Itoa(newRand.Intn(100500))
	sms.GeneralData.OrderData.MerchantTransactionID = fmt.Sprintf("TestTranID:%d", newRand.Intn(rand.Int()))
	sms.GeneralData.OrderData.OrderDescription = "Gopher Gufer ordering goods"
	sms.GeneralData.OrderData.OrderID = fmt.Sprintf("TestOrderID:%d", newRand.Intn(rand.Int()))
	sms.GeneralData.CustomerData.Email = "some@email.com"
	sms.GeneralData.CustomerData.BillingAddress.City = "Riga"
	sms.PaymentMethod.Pan = "5262482284416445"
	sms.PaymentMethod.ExpMmYy = "12/20"
	sms.PaymentMethod.Cvv = "403"
	sms.Money.Amount = newRand.Intn(11)
	sms.Money.Currency = "EUR"
	sms.System.UserIP = "127.0.0.1"
	sms.System.XForwardedFor = "127.0.0.1"

	resp, reqErr := gc.NewRequest(sms)
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
	gc, gcErr := NewGatewayClient(testAuth.AccountID, testAuth.SecretKey)
	if gcErr != nil {
		t.Error(gcErr)
		return
	}

	// Create some random values for our request
	newSource := rand.NewSource(time.Now().UnixNano())
	newRand := rand.New(newSource)

	sms := gc.OperationBuilder().NewSms()
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
	sms.Money.Amount = newRand.Intn(5)
	sms.Money.Currency = "EUR"
	sms.System.UserIP = "127.0.0.1"
	sms.System.XForwardedFor = "127.0.0.1"

	resp, reqErr := gc.NewRequest(sms)
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

	parsedRes, parseErr := gc.ParseResponse(resp, structures.SMS)
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
	gc, gcErr := NewGatewayClient(testAuth.AccountID, "ImBroken")
	if gcErr != nil {
		t.Error(gcErr)
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
		t.Errorf("Failed to read received body: %s ", bodyErr.Error)
		return
	}

	var gwResp structures.UnauthorizedResponse

	parseErr := json.Unmarshal(body, &gwResp)
	if parseErr != nil {
		if bodyErr != nil {
			t.Errorf("Failed to unmarshal received body: %s ", bodyErr.Error)
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
	gc, gcErr := NewGatewayClient(testAuth.AccountID, testAuth.SecretKey)
	if gcErr != nil {
		t.Error(gcErr)
		return
	}

	// Create some random values for our request
	newSource := rand.NewSource(time.Now().UnixNano())
	newRand := rand.New(newSource)
	tranAmount := newRand.Intn(12)

	opBuild := gc.OperationBuilder()

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

	respHold, respHoldErr := gc.NewRequest(holdDMS)
	if respHoldErr != nil {
		t.Error(respHoldErr)
		return
	}

	if respHold == nil {
		t.Error("DMS Hold parsed response is empty")
		return
	}

	parsedHoldRes, parseErr := gc.ParseResponse(respHold, structures.DMSHold)
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

	respCharge, respChargeErr := gc.NewRequest(chargeDMS)
	if respChargeErr != nil {
		t.Error(respChargeErr)
		return
	}

	if respCharge == nil {
		t.Error("DMS Charge parsed response is empty")
		return
	}

	parsedCharge, parseChargeErr := gc.ParseResponse(respCharge, structures.DMSCharge)
	if parseChargeErr != nil {
		t.Error(parseChargeErr)
		return
	}

	// @TODO Debug print
	fmt.Println(fmt.Sprintf("%+v", parsedCharge))
}

func TestSendRequestDMSCancel(t *testing.T) {
	gc, gcErr := NewGatewayClient(testAuth.AccountID, testAuth.SecretKey)
	if gcErr != nil {
		t.Error(gcErr)
		return
	}

	// Create some random values for our request
	newSource := rand.NewSource(time.Now().UnixNano())
	newRand := rand.New(newSource)
	tranAmount := newRand.Intn(10)

	opBuild := gc.OperationBuilder()

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

	respHold, respHoldErr := gc.NewRequest(holdDMS)
	if respHoldErr != nil {
		t.Error(respHoldErr)
		return
	}

	if respHold == nil {
		t.Error("DMS Hold parsed response is empty")
		return
	}

	parsedHoldRes, parseErr := gc.ParseResponse(respHold, structures.DMSHold)
	if parseErr != nil {
		t.Error(parseErr)
		return
	}

	// @TODO Debug print
	fmt.Println(fmt.Sprintf("%+v", parsedHoldRes))

	cancel := opBuild.NewCancel()
	cancel.CommandData.GWTransactionID = parsedHoldRes.(structures.TransactionResponse).GateWay.GatewayTransactionID

	respCancel, respCancelErr := gc.NewRequest(cancel)
	if respCancelErr != nil {
		t.Error(respCancelErr)
		return
	}

	if respCancel == nil {
		t.Error("Cancel parsed response is empty")
		return
	}

	parsedCancel, parseCancelErr := gc.ParseResponse(respCancel, structures.CANCEL)
	if parseCancelErr != nil {
		t.Error(parseCancelErr)
		return
	}

	// @TODO Debug print
	fmt.Println(fmt.Sprintf("%+v", parsedCancel))
}

func TestSendRequestGetStatus(t *testing.T) {
	gc, gcErr := NewGatewayClient(testAuth.AccountID, testAuth.SecretKey)
	if gcErr != nil {
		t.Error(gcErr)
		return
	}

	// Create some random values for our request
	newSource := rand.NewSource(time.Now().UnixNano())
	newRand := rand.New(newSource)

	opBuilder := gc.OperationBuilder()
	sms := opBuilder.NewSms()
	sms.GeneralData.OrderData.OrderDescription = "Gopher Gufer will get transaction status"
	sms.PaymentMethod.Pan = "5262482284416445"
	sms.PaymentMethod.ExpMmYy = "12/20"
	sms.PaymentMethod.Cvv = "403"
	sms.Money.Amount = newRand.Intn(10)
	sms.Money.Currency = "EUR"
	sms.System.UserIP = "127.0.0.1"
	sms.System.XForwardedFor = "127.0.0.1"

	resp, reqErr := gc.NewRequest(sms)
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

	parsedRes, parseErr := gc.ParseResponse(resp, structures.SMS)
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

	gwIds := make([]string, 1)
	gwIds[0] = parsedRes.(structures.TransactionResponse).GateWay.GatewayTransactionID

	getStatus := opBuilder.NewGetStatus()
	getStatus.CommandData.GWTransactionIDs = gwIds
	getStatus.System.UserIP = "127.0.0.1"
	getStatus.System.XForwardedFor = "127.0.0.1"

	respGetStatus, reqGetStatusErr := gc.NewRequest(getStatus)
	if reqGetStatusErr != nil {
		t.Error(reqGetStatusErr)
		return
	}

	// @TODO Debug print
	fmt.Println(fmt.Sprintf("%+v", respGetStatus))

	parsedGetStatus, parsedGetStatusErr := gc.ParseResponse(respGetStatus, structures.Status)
	if parsedGetStatusErr != nil {
		t.Error(parsedGetStatusErr)
		return
	}

	// @TODO Debug print
	fmt.Println(fmt.Sprintf("%+v", parsedGetStatus))

	tranSatus := parsedGetStatus.([]structures.ExploringResponse)

	for _, tranD := range tranSatus {
		// @TODO Debug print
		fmt.Println(tranD.GatewayTransactionID)
		for _, tranS := range tranD.Status {
			// @TODO Debug print
			fmt.Println(fmt.Sprintf("StatusCodeGeneral: %d  : %s", tranS.StatusCodeGeneral, tranS.StatusTextGeneral))
			fmt.Println(fmt.Sprintf("StatusCode: %d  : %s", tranS.StatusCode, tranS.StatusText))
		}
	}
}
