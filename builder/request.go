package builder

import (
	"bitbucket.transactpro.lv/tls/gw3-go-client/structures"
)

type (
	// requestHTTPData contains HTTP request method and operationType to append in URL path
	requestHTTPData struct {
		// HTTP method
		method string
		// Operation type
		operationType OperationType
	}

	// RequestData combined request body payload
	RequestData struct {
		Auth structures.AuthData `json:"auth-data"`
		Data interface{}         `json:"data"`
	}

	// RequestBuilder provides request payload structure
	RequestBuilder struct {
		RequestData
	}
)

type RequestHTTPDataInterface interface {
	GetHTTPMethod() string
	GetOperationType() OperationType
}

// GetHTTPMethod return HTTP method which will be used for send request
func (rd *requestHTTPData) GetHTTPMethod() string {
	return rd.method
}

// GetOperationType return part of route path which will be used for send request
func (rd *requestHTTPData) GetOperationType() OperationType {
	return rd.operationType
}

// SetMerchantAuthData method allows correctly append merchant authorization data structure to request data
func (rb *RequestBuilder) SetMerchantAuthData(auth structures.AuthData) {
	rb.Auth = auth
}

// SetPayloadData method allows correctly append operation data structure to request data
func (rb *RequestBuilder) SetPayloadData(data interface{}) {
	rb.Data = data
}
