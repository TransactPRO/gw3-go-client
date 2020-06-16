package structures

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRequestDigest(t *testing.T) {
	expected := "Digest username=bc501eda-e2a1-4e63-9a1e-7a7f6ff4813b, uri=\"/v3.0/sms\", algorithm=SHA-256, " +
		"cnonce=\"MTU5MTYyNTA2MzqydV+lpoF4ZtfSAifxoUretZdAzGaZa97iRogrQ8K/yg==\", qop=auth-int, " +
		"response=\"a21df219fd9bb2efb71554eb9ebb47f6a7a61769a289f9ab4fcbe41d7544e28d\""

	body := "{\"auth-data\":{},\"data\":{\"command-data\":{},\"general-data\":{\"customer-data\":{\"email\":\"test@test.domain\",\"birth-date\":\"01/00\"," +
		"\"phone\":\"123456789\",\"billing-address\":{\"country\":\"FR\",\"state\":\"FR\",\"city\":\"Chalon-sur-Saône\",\"street\":\"Rue Garibaldi\"," +
		"\"house\":\"10\",\"flat\":\"10\",\"zip\":\"71100\"},\"shipping-address\":{\"country\":\"FR\",\"state\":\"FR\",\"city\":\"Chalon-sur-Saône\"," +
		"\"street\":\"Rue Garibaldi\",\"house\":\"10\",\"flat\":\"10\",\"zip\":\"71100\"}},\"order-data\":{\"merchant-transaction-id\":\"\"," +
		"\"order-id\":\"Order ID\",\"order-description\":\"Payment\",\"merchant-side-url\":\"https://domain.com/custom-url/\"," +
		"\"merchant-referring-name\":\"Test payment\",\"custom3d-return-url\":\"https://domain.com\"}},\"payment-method-data\":" +
		"{\"pan\":\"4111111111111111\",\"exp-mm-yy\":\"09/31\",\"cvv\":\"123\",\"cardholder-name\":\"John Doe\"},\"money-data\":" +
		"{\"amount\":100,\"currency\":\"EUR\"},\"system\":{\"user-ip\":\"127.0.0.1\"}}}"

	instance, err := NewRequestDigest(
		"bc501eda-e2a1-4e63-9a1e-7a7f6ff4813b",
		"agHJSthpTPfKEORLDynBuIl07i4sYVmw",
		"/v3.0/sms",
		[]byte(body),
	)

	assert.NoError(t, err)
	assert.Equal(t, AlgorithmSHA256, instance.Algorithm)
	assert.Equal(t, QopAuthInt, instance.QOP)
	assert.Equal(t, 43, len(instance.Cnonce))

	instance.Cnonce, _ = base64.StdEncoding.DecodeString("MTU5MTYyNTA2MzqydV+lpoF4ZtfSAifxoUretZdAzGaZa97iRogrQ8K/yg==")
	actual, err2 := instance.CreateHeader()
	assert.NoError(t, err2)

	assert.Equal(t, expected, actual)
}

func TestResponseDigestParseErrors(t *testing.T) {
	nonce := base64.StdEncoding.EncodeToString([]byte("1:q"))
	noTsNonce := base64.StdEncoding.EncodeToString([]byte("qqq"))
	wrongTsNonce := base64.StdEncoding.EncodeToString([]byte("qqq:www"))

	// format: [authorization header value: expected error]
	examples := []struct{ digest, expectedError string }{
		{"", "authorization header is missing"},
		{
			fmt.Sprintf("Digest uri=b, algorithm=SHA-256, cnonce=%s, snonce=%s, qop=auth-int, response=e", nonce, nonce),
			"digest mismatch: empty value for username",
		},
		{
			fmt.Sprintf("Digest username=a, algorithm=SHA-256, cnonce=%s, snonce=%s, qop=auth-int, response=e", nonce, nonce),
			"digest mismatch: empty value for uri",
		},
		{
			fmt.Sprintf("Digest username=a, uri=b, cnonce=%s, snonce=%s, qop=auth-int, response=e", nonce, nonce),
			"digest mismatch: empty value for algorithm",
		},
		{
			fmt.Sprintf("Digest username=a, uri=b, algorithm=SHA-256, snonce=%s, qop=auth-int, response=e", nonce),
			"digest mismatch: empty value for cnonce",
		},
		{
			fmt.Sprintf("Digest username=a, uri=b, algorithm=SHA-256, cnonce=%s, qop=auth-int, response=e", nonce),
			"digest mismatch: empty value for snonce",
		},
		{
			fmt.Sprintf("Digest username=a, uri=b, algorithm=SHA-256, cnonce=%s, snonce=%s, response=e", nonce, nonce),
			"digest mismatch: empty value for qop",
		},
		{
			fmt.Sprintf("Digest username=a, uri=b, algorithm=SHA-256, cnonce=%s, snonce=%s, qop=auth", nonce, nonce),
			"digest mismatch: empty value for response",
		},
		{
			fmt.Sprintf("Digest username=a, uri=b, algorithm=SHA-256, cnonce=%s, snonce=%s, qop=aaa, response=x", nonce, nonce),
			"digest mismatch: format error: unknown QOP value aaa",
		},
		{
			fmt.Sprintf("Digest username=a, uri=b, algorithm=aaa, cnonce=%s, snonce=%s, qop=auth, response=x", nonce, nonce),
			"digest mismatch: format error: unknown algorithm aaa",
		},
		{
			fmt.Sprintf("Digest username=a, uri=b, algorithm=SHA-256, cnonce=%s, snonce=%s, qop=auth, response=x", "123", nonce),
			"digest mismatch: corrupted value for cnonce (illegal base64 data at input byte 0)",
		},
		{
			fmt.Sprintf("Digest username=a, uri=b, algorithm=SHA-256, cnonce=%s, snonce=%s, qop=auth, response=x", nonce, "123"),
			"digest mismatch: corrupted value for snonce (illegal base64 data at input byte 0)",
		},
		{
			fmt.Sprintf("Digest username=a, uri=b, algorithm=SHA-256, cnonce=%s, snonce=%s, qop=auth, response=x", nonce, noTsNonce),
			"digest mismatch: corrupted value for snonce (missing timestamp)",
		},
		{
			fmt.Sprintf("Digest username=a, uri=b, algorithm=SHA-256, cnonce=%s, snonce=%s, qop=auth, response=x", nonce, wrongTsNonce),
			"digest mismatch: corrupted value for snonce (unexpected timestamp value: strconv.Atoi: parsing \"qqq\": invalid syntax)",
		},
	}

	for _, testCase := range examples {
		t.Run(testCase.expectedError, func(t *testing.T) {
			_, err := NewResponseDigest(testCase.digest)
			assert.EqualError(t, err, testCase.expectedError)
		})
	}
}

func TestResponseDigestParseSuccessful(t *testing.T) {
	expectedCnonce, _ := base64.StdEncoding.DecodeString("MTU5MTYyNTA2MzqydV+lpoF4ZtfSAifxoUretZdAzGaZa97iRogrQ8K/yg==")
	expectedSnonce, _ := base64.StdEncoding.DecodeString("MTU5MTYyNDgwNzoUte6YsXIJmUo1EsA4yrYDCVbPrvCrEtqGq6CHTMhImg==")

	headerValue := "Digest username=bc501eda-e2a1-4e63-9a1e-7a7f6ff4813b, uri=\"/v3.0/sms\", algorithm=SHA-256, " +
		"cnonce=\"MTU5MTYyNTA2MzqydV+lpoF4ZtfSAifxoUretZdAzGaZa97iRogrQ8K/yg==\", " +
		"snonce=\"MTU5MTYyNDgwNzoUte6YsXIJmUo1EsA4yrYDCVbPrvCrEtqGq6CHTMhImg==\", qop=auth-int, " +
		"response=\"a21df219fd9bb2efb71554eb9ebb47f6a7a61769a289f9ab4fcbe41d7544e28d\""

	digest, err := NewResponseDigest(headerValue)
	assert.NoError(t, err)

	assert.Equal(t, "bc501eda-e2a1-4e63-9a1e-7a7f6ff4813b", digest.Username)
	assert.Equal(t, "/v3.0/sms", digest.URI)
	assert.Equal(t, AlgorithmSHA256, digest.Algorithm)
	assert.Equal(t, QopAuthInt, digest.QOP)
	assert.Equal(t, "a21df219fd9bb2efb71554eb9ebb47f6a7a61769a289f9ab4fcbe41d7544e28d", digest.Response)
	assert.Equal(t, 1591624807, digest.Timestamp)
	assert.Equal(t, expectedCnonce, digest.Cnonce)
	assert.Equal(t, expectedSnonce, digest.Snonce)
}

func TestResponseDigestVerifyErrors(t *testing.T) {
	body := "{\"acquirer-details\":{},\"error\":{},\"gw\":{\"gateway-transaction-id\":\"37b88436-b69c-45f3-ad26-b945153ad9a8\"," +
		"\"redirect-url\":\"http://api.local/4f1f647d10e8296a2ed4d21e3639f1ee\",\"status-code\":30,\"status-text\":" +
		"\"INSIDE FORM URL SENT\"},\"warnings\":[\"Soon counters will be exceeded for the merchant\",\"Soon counters will be exceeded " +
		"for the account\"]}"

	responseHeader := "Digest username=bc501eda-e2a1-4e63-9a1e-7a7f6ff4813b, uri=\"/v3.0/sms\", algorithm=SHA-256, " +
		"cnonce=\"MTU5MTg2NjU3Mzo38zMeHvu4qcbhR8X158atP/BB4dDb5DbOMRT656yS7Q==\", " +
		"snonce=\"MTU5MTg2NjU3MzpvnttqUse7hfrkUHtPS8tWE1jl0D0G/DgMmEFwbk5/jw==\", qop=auth-int, " +
		"response=\"624478f45d33bbadc7cf0ae9b34462efd7b9736111f295e6330fe0bc3b20acda\""

	validCnonce, _ := base64.StdEncoding.DecodeString("MTU5MTg2NjU3Mzo38zMeHvu4qcbhR8X158atP/BB4dDb5DbOMRT656yS7Q==")
	invalidCnonce, _ := base64.StdEncoding.DecodeString("MTU5MTg2NjU3MzpvnttqUse7hfrkUHtPS8tWE1jl0D0G/DgMmEFwbk5/jw==")

	examples := []struct {
		guid       string
		origURI    string
		origCnonce []byte
		error      string
	}{
		{"wrong-guid", "/v3.0/sms", validCnonce, "digest mismatch: username mismatch"},
		{"bc501eda-e2a1-4e63-9a1e-7a7f6ff4813b", "http://another.local", validCnonce, "digest mismatch: uri mismatch"},
		{"bc501eda-e2a1-4e63-9a1e-7a7f6ff4813b", "/v3.0/sms", invalidCnonce, "digest mismatch: cnonce mismatch"},
		{"bc501eda-e2a1-4e63-9a1e-7a7f6ff4813b", "/v3.0/sms", validCnonce, "digest mismatch"},
	}

	for _, testCase := range examples {
		t.Run(testCase.error, func(t *testing.T) {
			responseDigest, err := NewResponseDigest(responseHeader)
			assert.NoError(t, err)

			responseDigest.OriginalURI = testCase.origURI
			responseDigest.OriginalCnonce = testCase.origCnonce
			responseDigest.Body = []byte(body)

			verifyErr := responseDigest.Verify(testCase.guid, "something wrong")
			assert.EqualError(t, verifyErr, testCase.error)
		})
	}
}

func TestResponseDigestVerifySuccessFullChecks(t *testing.T) {
	body := "{\"acquirer-details\":{},\"error\":{},\"gw\":{\"gateway-transaction-id\":\"37b88436-b69c-45f3-ad26-b945153ad9a8\"," +
		"\"redirect-url\":\"http://api.local/4f1f647d10e8296a2ed4d21e3639f1ee\",\"status-code\":30,\"status-text\":" +
		"\"INSIDE FORM URL SENT\"},\"warnings\":[\"Soon counters will be exceeded for the merchant\",\"Soon counters will be exceeded " +
		"for the account\"]}"

	responseHeader := "Digest username=bc501eda-e2a1-4e63-9a1e-7a7f6ff4813b, uri=\"/v3.0/sms\", algorithm=SHA-256, " +
		"cnonce=\"MTU5MTg2NjU3Mzo38zMeHvu4qcbhR8X158atP/BB4dDb5DbOMRT656yS7Q==\", " +
		"snonce=\"MTU5MTg2NjU3MzpvnttqUse7hfrkUHtPS8tWE1jl0D0G/DgMmEFwbk5/jw==\", qop=auth-int, " +
		"response=\"dda7026eebbeeee19fda191fd951d470b2064e3e1bc416365835abc775352552\""

	responseDigest, err := NewResponseDigest(responseHeader)
	assert.NoError(t, err)

	responseDigest.OriginalURI = "/v3.0/sms"
	responseDigest.OriginalCnonce, _ = base64.StdEncoding.DecodeString("MTU5MTg2NjU3Mzo38zMeHvu4qcbhR8X158atP/BB4dDb5DbOMRT656yS7Q==")
	responseDigest.Body = []byte(body)

	verifyErr := responseDigest.Verify("bc501eda-e2a1-4e63-9a1e-7a7f6ff4813b", "tPMOogw7YBumh6RpXxi2nvGW0C9lJq3L")
	assert.NoError(t, verifyErr)
}

func TestResponseDigestVerifySuccessMinimalChecks(t *testing.T) {
	body := "{\"acquirer-details\":{},\"error\":{},\"gw\":{\"gateway-transaction-id\":\"37b88436-b69c-45f3-ad26-b945153ad9a8\"," +
		"\"redirect-url\":\"http://api.local/4f1f647d10e8296a2ed4d21e3639f1ee\",\"status-code\":30,\"status-text\":" +
		"\"INSIDE FORM URL SENT\"},\"warnings\":[\"Soon counters will be exceeded for the merchant\",\"Soon counters will be exceeded " +
		"for the account\"]}"

	responseHeader := "Digest username=bc501eda-e2a1-4e63-9a1e-7a7f6ff4813b, uri=\"/v3.0/sms\", algorithm=SHA-256, " +
		"cnonce=\"MTU5MTg2NjU3Mzo38zMeHvu4qcbhR8X158atP/BB4dDb5DbOMRT656yS7Q==\", " +
		"snonce=\"MTU5MTg2NjU3MzpvnttqUse7hfrkUHtPS8tWE1jl0D0G/DgMmEFwbk5/jw==\", qop=auth-int, " +
		"response=\"dda7026eebbeeee19fda191fd951d470b2064e3e1bc416365835abc775352552\""

	responseDigest, err := NewResponseDigest(responseHeader)
	assert.NoError(t, err)

	responseDigest.Body = []byte(body)

	verifyErr := responseDigest.Verify("bc501eda-e2a1-4e63-9a1e-7a7f6ff4813b", "tPMOogw7YBumh6RpXxi2nvGW0C9lJq3L")
	assert.NoError(t, verifyErr)
}

func TestResponseDigestVerifyCallback(t *testing.T) {
	jsonFromPost := "{\"result-data\":{\"gw\":{\"gateway-transaction-id\":\"8d77f986-de7f-4d47-97ef-9de7f8561684\",\"status-code\":7,\"status-text\":\"SUCCESS\"}," +
		"\"error\":{},\"acquirer-details\":{\"eci-sli\":\"503\",\"terminal-mid\":\"3201210\",\"transaction-id\":\"7146311464333929\"," +
		"\"result-code\":\"000\",\"status-text\":\"Approved\",\"status-description\":\"Approved\"},\"warnings\":" +
		"[\"Soon counters will be exceeded for the merchant\",\"Soon counters will be exceeded for the account\"," +
		"\"Soon counters will be exceeded for the terminal group\",\"Soon counters will be exceeded for the terminal\"]}}"

	signFromPost := "Digest username=bc501eda-e2a1-4e63-9a1e-7a7f6ff4813b, uri=\"/v3.0/sms\", algorithm=SHA-256, " +
		"cnonce=\"MTU5MTg2OTQ3OTpbmPfGQxVAh5z7MdWnRjF1cavfwKyxiLVrX4p7IHNwWA==\", " +
		"snonce=\"MTU5MTg2OTQ4MTqfPxash/0hfNpI/gHuaoSiV+6PwVKYEawxchE0nxHTkA==\", qop=auth-int, " +
		"response=\"87bd753875e28da54dfcb5e61614e10a7120aba9a3f8bed0e6eaa9acb85aa9f9\""

	responseDigest, err := NewResponseDigest(signFromPost)
	assert.NoError(t, err)

	responseDigest.OriginalURI = "/v3.0/sms"
	responseDigest.OriginalCnonce, _ = base64.StdEncoding.DecodeString("MTU5MTg2OTQ3OTpbmPfGQxVAh5z7MdWnRjF1cavfwKyxiLVrX4p7IHNwWA==")
	responseDigest.Body = []byte(jsonFromPost)
	verifyErr := responseDigest.Verify("bc501eda-e2a1-4e63-9a1e-7a7f6ff4813b", "tPMOogw7YBumh6RpXxi2nvGW0C9lJq3L")
	assert.NoError(t, verifyErr)

	var parsedResult CallbackResult
	parseErr := json.Unmarshal(responseDigest.Body, &parsedResult)
	assert.NoError(t, parseErr)
	assert.Equal(t, "8d77f986-de7f-4d47-97ef-9de7f8561684", parsedResult.ResultData.Gateway.GatewayTransactionID)
}
