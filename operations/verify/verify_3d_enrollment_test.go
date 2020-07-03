package verify

import (
	"testing"

	"github.com/TransactPRO/gw3-go-client/structures"
	"github.com/stretchr/testify/assert"
)

func TestParseEnrollmentResponse(t *testing.T) {
	examples := map[string]structures.Enrollment{
		"{\"enrollment\":\"y\"}":           structures.EnrollmentYes,
		"{\"enrollment\":\"n\"}":           structures.EnrollmentNo,
		"{\"enrollment\":\"abracadabra\"}": structures.EnrollmentUnknown,
	}

	for input, expectedOutput := range examples {
		t.Run(expectedOutput.String(), func(t *testing.T) {
			response := structures.NewGatewayResponse(nil, []byte(input))
			parsedResponse, err := NewVerify3dEnrollmentAssembly().ParseResponse(response)
			assert.NoError(t, err)

			assert.Equal(t, expectedOutput, parsedResponse.Enrollment)
		})
	}
}
