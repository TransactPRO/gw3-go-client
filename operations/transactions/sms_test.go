package transactions

import (
	"net/url"
	"testing"

	"github.com/TransactPRO/gw3-go-client/structures"
	"github.com/stretchr/testify/assert"
)

func TestParseSMSResponseSuccessfulAPI(t *testing.T) {
	body := "{\"acquirer-details\":{\"dynamic-descriptor\":\"test\",\"eci-sli\":\"648\",\"result-code\":\"000\",\"status-description\":\"Approved\"," +
		"\"status-text\":\"Approved\",\"terminal-mid\":\"5800978\",\"transaction-id\":\"1899493845214315\"},\"error\":{}," +
		"\"gw\":{\"gateway-transaction-id\":\"8a9bed66-8412-494f-9866-2c26b5ceee62\",\"status-code\":7,\"status-text\":\"SUCCESS\"," +
		"\"original-gateway-transaction-id\":\"orig-aaa\",\"parent-gateway-transaction-id\":\"parent-aaa\"}," +
		"\"warnings\":[\"Soon counters will be exceeded for the merchant\",\"Soon counters will be exceeded for the account\"," +
		"\"Soon counters will be exceeded for the terminal group\",\"Soon counters will be exceeded for the terminal\"]}\n"

	response := structures.NewGatewayResponse(nil, []byte(body))
	parsedResponse, err := NewSMSAssembly().ParseResponse(response)
	assert.NoError(t, err)

	assert.NotNil(t, parsedResponse.AcquirerDetails)
	assert.Equal(t, "test", parsedResponse.AcquirerDetails.DynamicDescriptor)
	assert.Equal(t, "648", parsedResponse.AcquirerDetails.EciSli)
	assert.Equal(t, "000", parsedResponse.AcquirerDetails.ResultCode)
	assert.Equal(t, "Approved", parsedResponse.AcquirerDetails.StatusDescription)
	assert.Equal(t, "Approved", parsedResponse.AcquirerDetails.StatusText)
	assert.Equal(t, "5800978", parsedResponse.AcquirerDetails.TerminalID)
	assert.Equal(t, "1899493845214315", parsedResponse.AcquirerDetails.TransactionID)

	assert.NotNil(t, parsedResponse.Gateway)
	assert.Equal(t, "8a9bed66-8412-494f-9866-2c26b5ceee62", parsedResponse.Gateway.GatewayTransactionID)
	assert.NotNil(t, parsedResponse.Gateway.OriginalGatewayTransactionID)
	assert.Equal(t, "orig-aaa", *parsedResponse.Gateway.OriginalGatewayTransactionID)
	assert.NotNil(t, parsedResponse.Gateway.ParentGatewayTransactionID)
	assert.Equal(t, "parent-aaa", *parsedResponse.Gateway.ParentGatewayTransactionID)
	assert.Equal(t, structures.StatusSuccess, parsedResponse.Gateway.StatusCode)
	assert.Equal(t, "SUCCESS", parsedResponse.Gateway.StatusText)

	expectedWarnings := []string{
		"Soon counters will be exceeded for the merchant",
		"Soon counters will be exceeded for the account",
		"Soon counters will be exceeded for the terminal group",
		"Soon counters will be exceeded for the terminal",
	}
	assert.Equal(t, expectedWarnings, parsedResponse.Warnings)
}

func TestParseSMSResponseSuccessfulRedirect(t *testing.T) {
	expectedURL, _ := url.Parse("https://api.url/a4345be5b8a1af9773b8b0642b49ff26")

	body := "{\"acquirer-details\": {},\"error\": {},\"gw\": {\"gateway-transaction-id\": \"965ffd17-1874-48d0-89f3-f2c2f06bf749\"," +
		"\"redirect-url\": \"https://api.url/a4345be5b8a1af9773b8b0642b49ff26\",\"status-code\": 30,\"status-text\": \"INSIDE FORM URL SENT\"}}"

	response := structures.NewGatewayResponse(nil, []byte(body))
	parsedResponse, err := NewSMSAssembly().ParseResponse(response)
	assert.NoError(t, err)

	assert.NotNil(t, parsedResponse.Gateway)
	assert.Equal(t, "965ffd17-1874-48d0-89f3-f2c2f06bf749", parsedResponse.Gateway.GatewayTransactionID)
	assert.NotNil(t, parsedResponse.Gateway.RedirectURL)
	assert.Equal(t, structures.URL(*expectedURL), *parsedResponse.Gateway.RedirectURL)
	assert.Equal(t, structures.StatusCardFormURLSent, parsedResponse.Gateway.StatusCode)
	assert.Equal(t, "INSIDE FORM URL SENT", parsedResponse.Gateway.StatusText)
}

func TestParseSMSResponseError(t *testing.T) {
	body := "{\"acquirer-details\": {},\"error\": {\"code\": 1102,\"message\": \"Invalid pan number. Failed assertion that pan (false) == true\"}," +
		"\"gw\":{\"gateway-transaction-id\": \"33f17d34-3796-45e0-9bba-a771e9d3e504\",\"status-code\": 19,\"status-text\": \"BR VALIDATION FAILED\"}," +
		"\"warnings\": [\"Soon counters will be exceeded for the merchant\",\"Soon counters will be exceeded for the account\"]}"

	response := structures.NewGatewayResponse(nil, []byte(body))
	parsedResponse, err := NewSMSAssembly().ParseResponse(response)
	assert.NoError(t, err)

	assert.NotNil(t, parsedResponse.Error)
	assert.Equal(t, structures.EecCardBadNumber, parsedResponse.Error.Code)
	assert.Equal(t, "Invalid pan number. Failed assertion that pan (false) == true", parsedResponse.Error.Message)

	assert.NotNil(t, parsedResponse.Gateway)
	assert.Equal(t, "33f17d34-3796-45e0-9bba-a771e9d3e504", parsedResponse.Gateway.GatewayTransactionID)
	assert.Equal(t, structures.StatusBusinessRulesValidationFailed, parsedResponse.Gateway.StatusCode)
	assert.Equal(t, "BR VALIDATION FAILED", parsedResponse.Gateway.StatusText)

	expectedWarnings := []string{
		"Soon counters will be exceeded for the merchant",
		"Soon counters will be exceeded for the account",
	}
	assert.Equal(t, expectedWarnings, parsedResponse.Warnings)
}
