package builder

import "bitbucket.transactpro.lv/tls/gw3-go-client/structures"

// RequestData combined request body payload
type RequestData struct {
	Auth structures.AuthData `json:"auth-data"`
	Data interface{}         `json:"data"`
}

// RequestBuilder provides request payload structure
type RequestBuilder struct {
	RequestData
}

// RequestDataBuilder contains interfaces to correctly assign whole request data
type RequestDataBuilder interface {
	setMerchantAuthData(Auth structures.AuthData)
	setPayloadData(Data struct{})
}

// SetMerchantAuthData method allows correctly append merchant authorization data structure to request data
func (rb *RequestBuilder) SetMerchantAuthData(auth structures.AuthData) {
	rb.Auth = auth
}

// SetPayloadData method allows correctly append operation data structure to request data
func (rb *RequestBuilder) SetPayloadData(data interface{}) {
	rb.Data = data
}
