package builders

import "bitbucket.transactpro.lv/tls/gw3-go-client/structures"

type RequestData struct {
	Auth structures.AuthData 		`json:"auth-data"`
	Data interface{} 	`json:"data"`
}

type RequestBuilder struct {}

type RequestDataBuilder interface {
	setMerchantAuthData(Auth structures.AuthData)
	setPayloadData(Data struct{})
}

func (rd *RequestData) setMerchantAuthData(auth structures.AuthData) {
	rd.Auth = auth
}

func (rd *RequestData) setPayloadData(data struct{}) {
	rd.Data = data
}