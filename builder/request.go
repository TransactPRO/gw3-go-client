package builder

import "bitbucket.transactpro.lv/tls/gw3-go-client/structures"

type RequestData struct {
	Auth structures.AuthData 	`json:"auth-data"`
	Data interface{} 		`json:"data"`
}

type RequestBuilder struct {
	RequestData
}

type RequestDataBuilder interface {
	setMerchantAuthData(Auth structures.AuthData)
	setPayloadData(Data struct{})
}

func (rb *RequestBuilder) SetMerchantAuthData(auth structures.AuthData) {
	rb.Auth = auth
}

func (rb *RequestBuilder) SetPayloadData(data interface{}) {
	rb.Data = data
}