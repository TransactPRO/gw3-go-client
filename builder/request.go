package builder

import "bitbucket.transactpro.lv/tls/gw3-go-client/structures"

type RequestData struct {
	auth structures.AuthData 	`json:"auth-data"`
	data interface{} 		`json:"data"`
}

type RequestBuilder struct {
	RequestData
}

type RequestDataBuilder interface {
	setMerchantAuthData(Auth structures.AuthData)
	setPayloadData(Data struct{})
}

func (rb *RequestBuilder) SetMerchantAuthData(auth structures.AuthData) {
	rb.auth = auth
}

func (rb *RequestBuilder) SetPayloadData(data interface{}) {
	rb.data = data
}