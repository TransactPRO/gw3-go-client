package exploring

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/TransactPRO/gw3-go-client/structures"
	"github.com/stretchr/testify/assert"
)

func TestParseHistoryResponse(t *testing.T) {
	expectedDate1, _ := time.Parse("2006-01-02 15:04:05", "2020-06-09 09:56:53")
	expectedDate2, _ := time.Parse("2006-01-02 15:04:05", "2020-06-09 09:57:53")

	body := "{\"transactions\":[{\"error\":{\"code\":400,\"message\":\"Failed to fetch data for transaction with gateway id: " +
		"a2975c68-e235-40a4-87a9-987824c20000\"},\"gateway-transaction-id\":\"a2975c68-e235-40a4-87a9-987824c20000\"}," +
		"{\"gateway-transaction-id\":\"a2975c68-e235-40a4-87a9-987824c2090a\",\"history\":[{\"date-updated\":\"2020-06-09 09:56:53\"," +
		"\"status-code-new\":2,\"status-code-old\":1,\"status-text-new\":\"SENT TO BANK\",\"status-text-old\":\"INIT\"}," +
		"{\"date-updated\":\"2020-06-09 09:57:53\",\"status-code-new\":7,\"status-code-old\":2,\"status-text-new\":\"SUCCESS\"," +
		"\"status-text-old\":\"SENT TO BANK\"}]}]}"

	response := structures.NewGatewayResponse(nil, []byte(body))
	parsedResponse, err := NewHistoryAssembly().ParseResponse(response)
	assert.NoError(t, err)
	assert.Equal(t, 2, len(parsedResponse.Transactions))

	tr1 := parsedResponse.Transactions[0]
	assert.Equal(t, "a2975c68-e235-40a4-87a9-987824c20000", tr1.GatewayTransactionID)
	assert.Equal(t, structures.ErrorCode(400), tr1.Error.Code)
	assert.Equal(t, "Failed to fetch data for transaction with gateway id: a2975c68-e235-40a4-87a9-987824c20000", tr1.Error.Message)

	tr2 := parsedResponse.Transactions[1]
	assert.Equal(t, "a2975c68-e235-40a4-87a9-987824c2090a", tr2.GatewayTransactionID)
	assert.Equal(t, 2, len(tr2.History))

	event1 := tr2.History[0]
	assert.Equal(t, structures.Time(expectedDate1), event1.DateUpdated)
	assert.Equal(t, structures.StatusInit, event1.StatusCodeOld)
	assert.Equal(t, structures.StatusSent2Bank, event1.StatusCodeNew)
	assert.Equal(t, "INIT", event1.StatusTextOld)
	assert.Equal(t, "SENT TO BANK", event1.StatusTextNew)

	event2 := tr2.History[1]
	assert.Equal(t, structures.Time(expectedDate2), event2.DateUpdated)
	assert.Equal(t, structures.StatusSent2Bank, event2.StatusCodeOld)
	assert.Equal(t, structures.StatusSuccess, event2.StatusCodeNew)
	assert.Equal(t, "SENT TO BANK", event2.StatusTextOld)
	assert.Equal(t, "SUCCESS", event2.StatusTextNew)
}

func TestParseLimitsResponse(t *testing.T) {
	body := "{\"childs\":[{\"childs\":[{\"childs\":[{\"counters\":[{\"counter-type\":\"TR_SUCCESS_AMOUNT\",\"currency\":\"EUR\"," +
		"\"limit\":5000000,\"payment-method-subtype\":\"all\",\"payment-method-type\":\"all\",\"value\":28410}," +
		"{\"counter-type\":\"TR_SUCCESS_COUNT\",\"currency\":\"EUR\",\"limit\":20000,\"payment-method-subtype\":\"all\"," +
		"\"payment-method-type\":\"all\",\"value\":992}],\"acq-terminal-id\":\"5800978\",\"title\":\"Test T1\",\"type\":\"terminal\"}]," +
		"\"counters\":[{\"counter-type\":\"TR_SUCCESS_AMOUNT\",\"currency\":\"EUR\",\"limit\":5000000,\"payment-method-subtype\":\"all\"," +
		"\"payment-method-type\":\"all\",\"value\":2400}],\"title\":\"Test TG\",\"type\":\"terminal-group\"}],\"counters\":" +
		"[{\"counter-type\":\"TR_SUCCESS_AMOUNT\",\"currency\":\"EUR\",\"limit\":5000000,\"payment-method-subtype\":\"all\"," +
		"\"payment-method-type\":\"all\",\"value\":2400}],\"title\":\"Test ACC\",\"type\":\"account\"}],\"counters\":" +
		"[{\"counter-type\":\"TR_SUCCESS_AMOUNT\",\"currency\":\"EUR\",\"limit\":5000000,\"payment-method-subtype\":\"all\"," +
		"\"payment-method-type\":\"all\",\"value\":2400}],\"title\":\"Test M\",\"type\":\"merchant\"}"

	response := structures.NewGatewayResponse(nil, []byte(body))
	merchant, err := NewLimitsAssembly().ParseResponse(response)
	assert.NoError(t, err)

	assert.Equal(t, "merchant", merchant.Type)
	assert.Equal(t, "Test M", merchant.Title)
	assert.Nil(t, merchant.AcqTerminalID)
	assert.Equal(t, 1, len(merchant.Children))
	assert.Equal(t, 1, len(merchant.Limits))
	assert.Equal(t, "TR_SUCCESS_AMOUNT", merchant.Limits[0].CounterType)
	assert.Equal(t, "EUR", merchant.Limits[0].Currency)
	assert.Equal(t, json.Number("5000000"), merchant.Limits[0].Limit)
	assert.Equal(t, json.Number("2400"), merchant.Limits[0].Value)
	assert.Equal(t, "all", merchant.Limits[0].PaymentMethodType)
	assert.Equal(t, "all", merchant.Limits[0].PaymentMethodSubtype)

	account := merchant.Children[0]
	assert.Equal(t, "account", account.Type)
	assert.Equal(t, "Test ACC", account.Title)
	assert.Nil(t, account.AcqTerminalID)
	assert.Equal(t, 1, len(account.Children))
	assert.Equal(t, 1, len(account.Limits))
	assert.Equal(t, "TR_SUCCESS_AMOUNT", account.Limits[0].CounterType)
	assert.Equal(t, "EUR", account.Limits[0].Currency)
	assert.Equal(t, json.Number("5000000"), account.Limits[0].Limit)
	assert.Equal(t, json.Number("2400"), account.Limits[0].Value)
	assert.Equal(t, "all", account.Limits[0].PaymentMethodType)
	assert.Equal(t, "all", account.Limits[0].PaymentMethodSubtype)

	terminalGroup := account.Children[0]
	assert.Equal(t, "terminal-group", terminalGroup.Type)
	assert.Equal(t, "Test TG", terminalGroup.Title)
	assert.Nil(t, terminalGroup.AcqTerminalID)
	assert.Equal(t, 1, len(terminalGroup.Children))
	assert.Equal(t, 1, len(terminalGroup.Limits))
	assert.Equal(t, "TR_SUCCESS_AMOUNT", terminalGroup.Limits[0].CounterType)
	assert.Equal(t, "EUR", terminalGroup.Limits[0].Currency)
	assert.Equal(t, json.Number("5000000"), terminalGroup.Limits[0].Limit)
	assert.Equal(t, json.Number("2400"), terminalGroup.Limits[0].Value)
	assert.Equal(t, "all", terminalGroup.Limits[0].PaymentMethodType)
	assert.Equal(t, "all", terminalGroup.Limits[0].PaymentMethodSubtype)

	terminal := terminalGroup.Children[0]
	assert.Equal(t, "terminal", terminal.Type)
	assert.Equal(t, "Test T1", terminal.Title)
	assert.NotNil(t, terminal.AcqTerminalID)
	assert.Equal(t, "5800978", *terminal.AcqTerminalID)
	assert.Nil(t, terminal.Children)
	assert.Equal(t, 2, len(terminal.Limits))

	assert.Equal(t, "TR_SUCCESS_AMOUNT", terminal.Limits[0].CounterType)
	assert.Equal(t, "EUR", terminal.Limits[0].Currency)
	assert.Equal(t, json.Number("5000000"), terminal.Limits[0].Limit)
	assert.Equal(t, json.Number("28410"), terminal.Limits[0].Value)
	assert.Equal(t, "all", terminal.Limits[0].PaymentMethodType)
	assert.Equal(t, "all", terminal.Limits[0].PaymentMethodSubtype)

	assert.Equal(t, "TR_SUCCESS_COUNT", terminal.Limits[1].CounterType)
	assert.Equal(t, "EUR", terminal.Limits[1].Currency)
	assert.Equal(t, json.Number("20000"), terminal.Limits[1].Limit)
	assert.Equal(t, json.Number("992"), terminal.Limits[1].Value)
	assert.Equal(t, "all", terminal.Limits[1].PaymentMethodType)
	assert.Equal(t, "all", terminal.Limits[1].PaymentMethodSubtype)
}

func TestParseRecurringTransactionsResponse(t *testing.T) {
	expectedDateFinished, _ := time.Parse("2006-01-02 15:04:05", "2020-06-09 09:56:53")

	body := "{\"transactions\":[{\"error\":{\"code\":400,\"message\":\"Failed to fetch data for transaction with gateway id: " +
		"9e09bad0-5704-4b78-bf6a-c612f0101900\"},\"gateway-transaction-id\":\"9e09bad0-5704-4b78-bf6a-c612f0101900\"}," +
		"{\"gateway-transaction-id\":\"9e09bad0-5704-4b78-bf6a-c612f010192a\",\"recurrents\":[{\"account-guid\":" +
		"\"bc501eda-e2a1-4e63-9a1e-7a7f6ff4813b\",\"account-id\":108,\"acq-terminal-id\":\"5800978\",\"acq-transaction-id\":" +
		"\"7435540948424227\",\"amount\":100,\"approval-code\":\"4773442\",\"cardholder-name\":\"John Doe\",\"currency\":\"EUR\"," +
		"\"date-finished\":\"2020-06-09 09:56:53\",\"eci-sli\":\"464\",\"gateway-transaction-id\":\"a2975c68-e235-40a4-87a9-987824c2090a\"," +
		"\"merchant-transaction-id\":\"52a9990bad03e15417c70ef11a8103e1\",\"status-code\":7,\"status-code-general\":13," +
		"\"status-text\":\"SUCCESS\",\"status-text-general\":\"REFUND SUCCESS\"}]}]}"

	response := structures.NewGatewayResponse(nil, []byte(body))
	parsedResponse, err := NewRecurrentsAssembly().ParseResponse(response)
	assert.NoError(t, err)
	assert.Equal(t, 2, len(parsedResponse.Transactions))

	tr1 := parsedResponse.Transactions[0]
	assert.Equal(t, "9e09bad0-5704-4b78-bf6a-c612f0101900", tr1.GatewayTransactionID)
	assert.Equal(t, structures.ErrorCode(400), tr1.Error.Code)
	assert.Equal(t, "Failed to fetch data for transaction with gateway id: 9e09bad0-5704-4b78-bf6a-c612f0101900", tr1.Error.Message)

	tr2 := parsedResponse.Transactions[1]
	assert.Equal(t, "9e09bad0-5704-4b78-bf6a-c612f010192a", tr2.GatewayTransactionID)
	assert.Equal(t, 1, len(tr2.Subsequent))

	info1 := tr2.Subsequent[0]
	assert.Equal(t, "bc501eda-e2a1-4e63-9a1e-7a7f6ff4813b", info1.AccountGUID)
	assert.Equal(t, "5800978", info1.AcqTerminalID)
	assert.Equal(t, "7435540948424227", info1.AcqTransactionID)
	assert.Equal(t, json.Number("100"), info1.Amount)
	assert.Equal(t, "4773442", info1.ApprovalCode)
	assert.Equal(t, "John Doe", info1.CardholderName)
	assert.Equal(t, "EUR", info1.Currency)
	assert.Equal(t, structures.Time(expectedDateFinished), info1.DateFinished)
	assert.Equal(t, "464", info1.EciSli)
	assert.Equal(t, "a2975c68-e235-40a4-87a9-987824c2090a", info1.GatewayTransactionID)
	assert.Equal(t, "52a9990bad03e15417c70ef11a8103e1", info1.MerchantTransactionID)
	assert.Equal(t, structures.StatusSuccess, info1.StatusCode)
	assert.Equal(t, structures.StatusRefundSuccess, info1.StatusCodeGeneral)
	assert.Equal(t, "SUCCESS", info1.StatusText)
	assert.Equal(t, "REFUND SUCCESS", info1.StatusTextGeneral)
}

func TestParseRefundsResponse(t *testing.T) {
	expectedDateFinished1, _ := time.Parse("2006-01-02 15:04:05", "2020-06-09 10:18:15")
	expectedDateFinished2, _ := time.Parse("2006-01-02 15:04:05", "2020-06-09 10:18:22")

	body := "{\"transactions\":[{\"gateway-transaction-id\":\"a2975c68-e235-40a4-87a9-987824c2090a\",\"refunds\":" +
		"[{\"account-guid\":\"bc501eda-e2a1-4e63-9a1e-7a7f6ff4813b\",\"account-id\":108,\"acq-terminal-id\":\"5800978\"," +
		"\"acq-transaction-id\":\"1128894405863338\",\"amount\":10,\"approval-code\":\"1299034\",\"cardholder-name\":\"John Doe\"," +
		"\"currency\":\"EUR\",\"date-finished\":\"2020-06-09 10:18:15\",\"eci-sli\":\"960\",\"gateway-transaction-id\":" +
		"\"508fd8b9-3f78-486b-812b-2756f44e1bc6\",\"merchant-transaction-id\":\"aaa1\",\"status-code\":13,\"status-code-general\":11," +
		"\"status-text\":\"REFUND SUCCESS\",\"status-text-general\":\"REFUND FAILED\"},{\"account-guid\":" +
		"\"bc501eda-e2a1-4e63-9a1e-7a7f6ff4813b\",\"account-id\":108,\"acq-terminal-id\":\"5800978\",\"acq-transaction-id\":" +
		"\"0508080614087693\",\"amount\":20,\"approval-code\":\"7117603\",\"cardholder-name\":\"John Doe\",\"currency\":\"EUR\"," +
		"\"date-finished\":\"2020-06-09 10:18:22\",\"eci-sli\":\"690\",\"gateway-transaction-id\":\"191228b8-fd2d-47c8-8ff7-d28ba799cdb4\"," +
		"\"merchant-transaction-id\":\"\",\"status-code\":13,\"status-code-general\":13,\"status-text\":\"REFUND SUCCESS\"," +
		"\"status-text-general\":\"REFUND SUCCESS\"}]},{\"error\":{\"code\":400,\"message\":" +
		"\"Failed to fetch data for transaction with gateway id: a2975c68-e235-40a4-87a9-987824c20900\"}," +
		"\"gateway-transaction-id\":\"a2975c68-e235-40a4-87a9-987824c20900\"}]}"

	response := structures.NewGatewayResponse(nil, []byte(body))
	parsedResponse, err := NewRefundsAssembly().ParseResponse(response)
	assert.NoError(t, err)
	assert.Equal(t, 2, len(parsedResponse.Transactions))

	tr1 := parsedResponse.Transactions[0]
	assert.Equal(t, "a2975c68-e235-40a4-87a9-987824c2090a", tr1.GatewayTransactionID)
	assert.Equal(t, 2, len(tr1.Refunds))

	refund1 := tr1.Refunds[0]
	assert.Equal(t, "bc501eda-e2a1-4e63-9a1e-7a7f6ff4813b", refund1.AccountGUID)
	assert.Equal(t, "5800978", refund1.AcqTerminalID)
	assert.Equal(t, "1128894405863338", refund1.AcqTransactionID)
	assert.Equal(t, json.Number("10"), refund1.Amount)
	assert.Equal(t, "1299034", refund1.ApprovalCode)
	assert.Equal(t, "John Doe", refund1.CardholderName)
	assert.Equal(t, "EUR", refund1.Currency)
	assert.Equal(t, structures.Time(expectedDateFinished1), refund1.DateFinished)
	assert.Equal(t, "960", refund1.EciSli)
	assert.Equal(t, "508fd8b9-3f78-486b-812b-2756f44e1bc6", refund1.GatewayTransactionID)
	assert.Equal(t, "aaa1", refund1.MerchantTransactionID)
	assert.Equal(t, structures.StatusRefundSuccess, refund1.StatusCode)
	assert.Equal(t, structures.StatusRefundFailed, refund1.StatusCodeGeneral)
	assert.Equal(t, "REFUND SUCCESS", refund1.StatusText)
	assert.Equal(t, "REFUND FAILED", refund1.StatusTextGeneral)

	refund2 := tr1.Refunds[1]
	assert.Equal(t, "bc501eda-e2a1-4e63-9a1e-7a7f6ff4813b", refund2.AccountGUID)
	assert.Equal(t, "5800978", refund2.AcqTerminalID)
	assert.Equal(t, "0508080614087693", refund2.AcqTransactionID)
	assert.Equal(t, json.Number("20"), refund2.Amount)
	assert.Equal(t, "7117603", refund2.ApprovalCode)
	assert.Equal(t, "John Doe", refund2.CardholderName)
	assert.Equal(t, "EUR", refund2.Currency)
	assert.Equal(t, structures.Time(expectedDateFinished2), refund2.DateFinished)
	assert.Equal(t, "690", refund2.EciSli)
	assert.Equal(t, "191228b8-fd2d-47c8-8ff7-d28ba799cdb4", refund2.GatewayTransactionID)
	assert.Equal(t, "", refund2.MerchantTransactionID)
	assert.Equal(t, structures.StatusRefundSuccess, refund2.StatusCode)
	assert.Equal(t, structures.StatusRefundSuccess, refund2.StatusCodeGeneral)
	assert.Equal(t, "REFUND SUCCESS", refund2.StatusText)
	assert.Equal(t, "REFUND SUCCESS", refund2.StatusTextGeneral)

	tr2 := parsedResponse.Transactions[1]
	assert.Equal(t, "a2975c68-e235-40a4-87a9-987824c20900", tr2.GatewayTransactionID)
	assert.Equal(t, structures.ErrorCode(400), tr2.Error.Code)
	assert.Equal(t, "Failed to fetch data for transaction with gateway id: a2975c68-e235-40a4-87a9-987824c20900", tr2.Error.Message)
}

func TestParseResultResponse(t *testing.T) {
	expectedDateCreated, _ := time.Parse("2006-01-02 15:04:05", "2020-06-10 08:37:22")
	expectedDateFinished, _ := time.Parse("2006-01-02 15:04:05", "2020-06-10 08:37:23")

	body := "{\"transactions\":[{\"date-created\":\"2020-06-10 08:37:22\",\"date-finished\":\"2020-06-10 08:37:23\"," +
		"\"gateway-transaction-id\":\"b552fe8c-0fe3-4982-b2d6-9c37fa96dc58\",\"result-data\":{\"acquirer-details\":" +
		"{\"eci-sli\":\"736\",\"result-code\":\"000\",\"status-description\":\"Approved\",\"status-text\":\"Approved\"," +
		"\"terminal-mid\":\"5800978\",\"transaction-id\":\"8225174463086463\"},\"error\":{},\"gw\":" +
		"{\"gateway-transaction-id\":\"b552fe8c-0fe3-4982-b2d6-9c37fa96dc58\",\"original-gateway-transaction-id\":" +
		"\"096a93f4-c4d9-4b46-bbe9-22e30031f2d2\",\"parent-gateway-transaction-id\":\"096a93f4-c4d9-4b46-bbe9-22e30031f2d2\"," +
		"\"status-code\":15,\"status-text\":\"CANCELLED\"}}},{\"error\":{\"code\":400,\"message\":" +
		"\"Failed to get transaction result for transaction with gateway id: 965ffd17-1874-48d0-89f3-f2c2f06bf749\"}," +
		"\"gateway-transaction-id\":\"965ffd17-1874-48d0-89f3-f2c2f06bf749\"}]}"

	response := structures.NewGatewayResponse(nil, []byte(body))
	parsedResponse, err := NewResultAssembly().ParseResponse(response)
	assert.NoError(t, err)
	assert.Equal(t, 2, len(parsedResponse.Transactions))

	tr1 := parsedResponse.Transactions[0]
	assert.Equal(t, "b552fe8c-0fe3-4982-b2d6-9c37fa96dc58", tr1.GatewayTransactionID)
	assert.Equal(t, structures.Time(expectedDateCreated), tr1.DateCreated)
	assert.Equal(t, structures.Time(expectedDateFinished), tr1.DateFinished)

	assert.NotNil(t, tr1.ResultData.AcquirerDetails)
	assert.Equal(t, "", tr1.ResultData.AcquirerDetails.DynamicDescriptor)
	assert.Equal(t, "736", tr1.ResultData.AcquirerDetails.EciSli)
	assert.Equal(t, "000", tr1.ResultData.AcquirerDetails.ResultCode)
	assert.Equal(t, "Approved", tr1.ResultData.AcquirerDetails.StatusDescription)
	assert.Equal(t, "Approved", tr1.ResultData.AcquirerDetails.StatusText)
	assert.Equal(t, "5800978", tr1.ResultData.AcquirerDetails.TerminalID)
	assert.Equal(t, "8225174463086463", tr1.ResultData.AcquirerDetails.TransactionID)

	assert.NotNil(t, tr1.ResultData.Gateway)
	assert.Equal(t, "b552fe8c-0fe3-4982-b2d6-9c37fa96dc58", tr1.ResultData.Gateway.GatewayTransactionID)
	assert.NotNil(t, tr1.ResultData.Gateway.OriginalGatewayTransactionID)
	assert.Equal(t, "096a93f4-c4d9-4b46-bbe9-22e30031f2d2", *tr1.ResultData.Gateway.OriginalGatewayTransactionID)
	assert.NotNil(t, tr1.ResultData.Gateway.ParentGatewayTransactionID)
	assert.Equal(t, "096a93f4-c4d9-4b46-bbe9-22e30031f2d2", *tr1.ResultData.Gateway.ParentGatewayTransactionID)
	assert.Equal(t, structures.StatusDmsCanceled, tr1.ResultData.Gateway.StatusCode)
	assert.Equal(t, "CANCELLED", tr1.ResultData.Gateway.StatusText)

	tr2 := parsedResponse.Transactions[1]
	assert.Equal(t, "965ffd17-1874-48d0-89f3-f2c2f06bf749", tr2.GatewayTransactionID)
	assert.Equal(t, structures.ErrorCode(400), tr2.Error.Code)
	assert.Equal(t, "Failed to get transaction result for transaction with gateway id: 965ffd17-1874-48d0-89f3-f2c2f06bf749", tr2.Error.Message)
}

func TestParseStatusResponse(t *testing.T) {
	body := "{\"transactions\":[{\"gateway-transaction-id\":\"cd7b8bdf-3c78-4540-95d0-68018d2aba97\",\"status\":" +
		"[{\"gateway-transaction-id\":\"cd7b8bdf-3c78-4540-95d0-68018d2aba97\",\"status-code\":7,\"status-code-general\":8," +
		"\"status-text\":\"SUCCESS\",\"status-text-general\":\"EXPIRED\"}]},{\"gateway-transaction-id\":\"37908991-789b-4d79-8c6a-f90ba0ce12b6\"," +
		"\"status\":[{\"gateway-transaction-id\":\"37908991-789b-4d79-8c6a-f90ba0ce12b6\",\"status-code\":8,\"status-code-general\":7," +
		"\"status-text\":\"EXPIRED\",\"status-text-general\":\"SUCCESS\"}]}," +
		"{\"error\":{\"code\":400,\"message\":\"Failed to fetch data for transaction with gateway id: 99900000-789b-4d79-8c6a-f90ba0ce12b0\"}," +
		"\"gateway-transaction-id\":\"99900000-789b-4d79-8c6a-f90ba0ce12b0\"}]}"

	response := structures.NewGatewayResponse(nil, []byte(body))
	parsedResponse, err := NewStatusAssembly().ParseResponse(response)
	assert.NoError(t, err)
	assert.Equal(t, 3, len(parsedResponse.Transactions))

	tr1 := parsedResponse.Transactions[0]
	assert.Equal(t, "cd7b8bdf-3c78-4540-95d0-68018d2aba97", tr1.GatewayTransactionID)
	assert.Equal(t, "cd7b8bdf-3c78-4540-95d0-68018d2aba97", tr1.Status[0].GatewayTransactionID)
	assert.Equal(t, structures.StatusSuccess, tr1.Status[0].StatusCode)
	assert.Equal(t, structures.StatusExpired, tr1.Status[0].StatusCodeGeneral)
	assert.Equal(t, "SUCCESS", tr1.Status[0].StatusText)
	assert.Equal(t, "EXPIRED", tr1.Status[0].StatusTextGeneral)

	tr2 := parsedResponse.Transactions[1]
	assert.Equal(t, "37908991-789b-4d79-8c6a-f90ba0ce12b6", tr2.GatewayTransactionID)
	assert.Equal(t, "37908991-789b-4d79-8c6a-f90ba0ce12b6", tr2.Status[0].GatewayTransactionID)
	assert.Equal(t, structures.StatusExpired, tr2.Status[0].StatusCode)
	assert.Equal(t, structures.StatusSuccess, tr2.Status[0].StatusCodeGeneral)
	assert.Equal(t, "EXPIRED", tr2.Status[0].StatusText)
	assert.Equal(t, "SUCCESS", tr2.Status[0].StatusTextGeneral)

	tr3 := parsedResponse.Transactions[2]
	assert.Equal(t, "99900000-789b-4d79-8c6a-f90ba0ce12b0", tr3.GatewayTransactionID)
	assert.Equal(t, structures.ErrorCode(400), tr3.Error.Code)
	assert.Equal(t, "Failed to fetch data for transaction with gateway id: 99900000-789b-4d79-8c6a-f90ba0ce12b0", tr3.Error.Message)
}
