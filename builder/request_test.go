package builder

import (
	"testing"
	"bitbucket.transactpro.lv/tls/gw3-go-client/structures"
)

func TestGetRequestBuilder(t *testing.T) {
	var rd RequestBuilder
	var tAuth structures.AuthData
	var tData interface{}

	if rd.auth != tAuth {
		t.Error("RequestBuilder's -> RequestBuilder struct Auth isn't type of structures.AuthData")
	}

	if rd.data != tData {
		t.Error("RequestBuilder's -> RequestBuilder struct Data isn't type interface")
	}
}

func TestSetMerchantAuthData(t *testing.T) {
	var rd RequestBuilder
	tAuth := structures.AuthData{AccountID:42, SecretKey:"Secret"}

	rd.SetMerchantAuthData(tAuth)
	if rd.auth.AccountID == 0 && rd.auth.SecretKey == "" {
		t.Error("RequestBuilder's method setMerchantAuthData didn't set AccounID and SecretKey value")
	}
}

func TestSetPayloadData(t *testing.T) {
	var rd RequestBuilder
	data := struct {
		SomeTransactionData string
	}{SomeTransactionData:"Transaction data"}


	rd.SetPayloadData(data)
	if rd.data != data {
		t.Error("RequestBuilder's method SetPayloadData didn't set new interface data object")
	}
}