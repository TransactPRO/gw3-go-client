package operations


type RequestData struct {
	Auth AuthData 		`json:"auth-data"`
	Data interface{} 	`json:"data"`
}

type RequestBuilder struct {
}

type RequestDataBuilder interface {
	setMerchantAuthData(Auth AuthData)
	setPayloadData(Data struct{})
}

func (rd *RequestData) setMerchantAuthData(auth AuthData) {
	rd.Auth = auth
}

func (rd *RequestData) setPayloadData(data struct{}) {
	rd.Data = data
}