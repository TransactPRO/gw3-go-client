// Package tprogateway, provide ability to make requests to Transact Pro Gateway API v3.
package tprogateway

import (
	"errors"
	"io"
	"encoding/json"
	"bytes"
	"net/http"
	"fmt"

	"bitbucket.transactpro.lv/tls/gw3-go-client/operations"
	"log"
)

// Default API settings
const (
	dAPIUri = "http://uriel.sk.fpngw3.env"
	dAPIVersion = "3.0"
)

type (
	// confAPI, endpoint config to rich Transact Pro system
	confAPI struct {
		Uri     string
		Version string
	}

	// GatewayClient, represents REST API client
	GatewayClient struct {
		API		*confAPI
		auth		*operations.AuthData
		httpClient	http.Client
	}
)

// NewGatewayClient, new instance of prepared gateway client structure
func NewGatewayClient(AccountID int, SecretKey string) (*GatewayClient, error) {
	if AccountID == 0 {
		return nil, errors.New("Account ID can not be 0, please use the given ID from Transact Pro.")
	}

	if SecretKey == "" {
		return nil, errors.New("Secret key can't be empty. It's needed for merchant authorization.")
	}

	return &GatewayClient {
		API:  &confAPI{
			Uri:dAPIUri, Version: dAPIVersion},
		auth: &operations.AuthData{
			AccountID: AccountID, SecretKey:SecretKey},
	}, nil
}

// buildAPIUrlPath, returns full URL path for request
func buildAPIUrlPath(gc *GatewayClient) (string, error) {
	if gc.API.Uri == "" {
		return "", errors.New("Gateway client's URL is empty in, API settings.")
	}

	if gc.API.Version == "" {
		return "", errors.New("Gateway client's Version is empty in, API settings.")
	}

	return fmt.Sprintf("%s/v%s/sms", gc.API.Uri, gc.API.Version), nil
}

// NewOp method, returns builder for needed operation, like SMS, Reversal, even exploring transaction such as Refund History
func (gc *GatewayClient) NewOp() *operations.OperationBuilder {
	return &operations.OperationBuilder{}
}

// NewRequest method, prepares whole HTTP request for Transact Pro API
func (gc *GatewayClient) NewRequest(operationData interface{}) (*http.Response, error) {
	// Build whole payload structure with nested data bundles
	reqData := &operations.RequestData{}
	reqData.Auth = *gc.auth
	reqData.Data = operationData
	if reqData == nil {
		return nil, errors.New("Payload data of request is empty.")
	}

	// When payload ready, convert it to Json format
	var buffer io.Reader
	var bReqData []byte

	bReqData, err := json.Marshal(&reqData)
	if err != nil {
		return nil, err
	}

	// Write json object to buffer
	buffer = bytes.NewBuffer(bReqData)

	// Get prepared URL path for API request
	urlPath, err := buildAPIUrlPath(gc)
	if err != nil {
		return nil, err
	}

	// Build whole HTTP request with payload data
	newReq, err := http.NewRequest("POST", urlPath, buffer)
	if err != nil {
		return nil, err
	}

	// Set default headers
	newReq.Header.Set("Accept", "application/json")
	// Default values for headers
	newReq.Header.Set("Content-type", "application/json")

	log.Println(newReq.Method, ": ", newReq.URL)

	// Send HTTP request to server
	resp, respErr := gc.httpClient.Do(newReq)
	if respErr != nil {
		return nil, respErr
	}

	return resp, nil
}