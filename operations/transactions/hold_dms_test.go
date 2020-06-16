package transactions

import (
	"net/url"
	"testing"

	"github.com/TransactPRO/gw3-go-client/structures"
	"github.com/stretchr/testify/assert"
)

func TestParseHoldDMSResponseSuccessfulRedirect(t *testing.T) {
	expectedURL, _ := url.Parse("https://api.url/a4345be5b8a1af9773b8b0642b49ff26")

	body := "{\"acquirer-details\": {},\"error\": {},\"gw\": {\"gateway-transaction-id\": \"965ffd17-1874-48d0-89f3-f2c2f06bf749\"," +
		"\"redirect-url\": \"https://api.url/a4345be5b8a1af9773b8b0642b49ff26\",\"status-code\": 30,\"status-text\": \"INSIDE FORM URL SENT\"}}"

	response := structures.NewGatewayResponse(nil, []byte(body))
	parsedResponse, err := NewHoldDMSAssembly().ParseResponse(response)
	assert.NoError(t, err)

	assert.NotNil(t, parsedResponse.Gateway)
	assert.Equal(t, "965ffd17-1874-48d0-89f3-f2c2f06bf749", parsedResponse.Gateway.GatewayTransactionID)
	assert.NotNil(t, parsedResponse.Gateway.RedirectURL)
	assert.Equal(t, structures.URL(*expectedURL), *parsedResponse.Gateway.RedirectURL)
	assert.Equal(t, structures.StatusCardFormURLSent, parsedResponse.Gateway.StatusCode)
	assert.Equal(t, "INSIDE FORM URL SENT", parsedResponse.Gateway.StatusText)
}
