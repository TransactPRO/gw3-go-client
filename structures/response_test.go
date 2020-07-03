package structures

import (
	"net/http"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSuccessful(t *testing.T) {
	examples := map[int]bool{
		200: true,
		201: true,
		302: true,
		402: true,
		404: false,
		401: false,
		403: false,
		500: false,
	}

	for code, expected := range examples {
		t.Run(strconv.Itoa(code), func(t *testing.T) {
			rawResp := &http.Response{StatusCode: code}
			gwResp := NewGatewayResponse(rawResp, nil)

			assert.Equal(t, expected, gwResp.Successful())
		})
	}
}

func TestParseJSONErrors(t *testing.T) {
	t.Run("nil", func(t *testing.T) {
		var instance *GatewayResponse
		err := instance.ParseJSON(nil)
		assert.EqualError(t, err, "cannot unmarshal JSON response: response is nil")
	})

	t.Run("corrupted payload", func(t *testing.T) {
		instance := &GatewayResponse{Payload: []byte("{")}
		err := instance.ParseJSON(nil)
		assert.EqualError(t, err, "cannot unmarshal JSON response: unexpected end of JSON input")
	})
}
