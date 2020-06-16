package helpers

import (
	"net/http"
	"net/url"
	"strconv"
	"testing"

	"github.com/TransactPRO/gw3-go-client/structures"
	"github.com/stretchr/testify/assert"
)

func TestNewRetrieveFormAssemblyNoURL(t *testing.T) {
	examples := []*structures.TransactionResponse{
		nil,
		{},
	}

	for i, input := range examples {
		t.Run("#"+strconv.Itoa(i), func(t *testing.T) {
			_, err := NewRetrieveFormAssembly(input)
			assert.EqualError(t, err, "response doesn't contain link to an HTML form")
		})
	}
}

func TestNewRetrieveFormAssemblySuccessful(t *testing.T) {
	redirectURL, _ := url.Parse("https://api.url/a4345be5b8a1af9773b8b0642b49ff26")

	input := &structures.TransactionResponse{
		Gateway: structures.Gateway{
			RedirectURL: (*structures.URL)(redirectURL),
		},
	}

	instance, err := NewRetrieveFormAssembly(input)
	assert.NoError(t, err)
	assert.Equal(t, http.MethodGet, instance.GetHTTPMethod())
	assert.Equal(t, structures.OperationType("https://api.url/a4345be5b8a1af9773b8b0642b49ff26"), instance.GetOperationType())
}
