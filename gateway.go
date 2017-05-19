// Package tprogateway, provide ability to make requests to Transact Pro Gateway API v3.
package tprogateway

import (
	"errors"
	"encoding/json"
	"bytes"
	"net/http"
	"io/ioutil"
	"fmt"

	"bitbucket.transactpro.lv/tls/gw3-go-client/structures"
	"bitbucket.transactpro.lv/tls/gw3-go-client/builder"
)

// Default API settings
const (
	dAPIBaseUri = "http://uriel.sk.fpngw3.env"
	dAPIVersion = "3.0"
)

type (
	// confAPI, endpoint config to rich Transact Pro system
	confAPI struct {
		Uri 	string
		Version string
	}

	// GatewayClient, represents REST API client
	GatewayClient struct {
		API		*confAPI
		auth		*structures.AuthData
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
			Uri:dAPIBaseUri, Version: dAPIVersion},
		auth: &structures.AuthData{
			AccountID: AccountID, SecretKey:SecretKey},
	}, nil
}

// NewOp method, returns builder for needed operation, like SMS, Reversal, even exploring transaction such as Refund History
func (gc *GatewayClient) NewOp() *builder.OperationBuilder {
	return &builder.OperationBuilder{}
}

// NewRequest method, prepares whole HTTP request for Transact Pro API
func (gc *GatewayClient) NewRequest(opType builder.OperationType, opData interface{}) (*http.Request, error) {
	bufPayload, bufErr := prepareJsonPayload(*gc.auth, opData)
	if bufErr != nil {
		return nil, bufErr
	}

	// Get prepared URL path for API request
	httpMethod, urlPath, errUrlPath := determineAPIAction(gc.API.Uri, gc.API.Version, opType)
	if errUrlPath != nil {
		return nil, errUrlPath
	}

	newReq, reqErr := buildHTTPRequest(httpMethod, urlPath, bufPayload)
	if reqErr != nil {
		return nil, reqErr
	}

	return newReq, nil
}

// SendRequest method, sends prepared HTTP request to destination point and returns response from Transact Pro system
func (gc *GatewayClient) SendRequest(req *http.Request) (interface{}, error) {
	resp, respErr := gc.httpClient.Do(req)
	if respErr != nil {
		return nil, respErr
	}

	parsedResp, parseErr := parseResponse(resp)
	if parseErr != nil {
		return nil, parseErr
	}

	return parsedResp, nil
}

// prepareJsonPayload, validates\combines AuthData and Data struct to one big structure and converts to json(Marshal) to buffer
func prepareJsonPayload(pAuth structures.AuthData, pData interface{}) (*bytes.Buffer, error) {
	// Build whole payload structure with nested data bundles
	reqData := &builder.RequestData{}
	reqData.Auth = pAuth
	reqData.Data = pData

	// When payload ready, convert it to Json format
	bReqData, err := json.Marshal(&reqData)
	if err != nil {
		return nil, err
	}

	// Write json object to buffer
	buffer := bytes.NewBuffer(bReqData)

	return buffer, nil
}

// determineAPIAction, determiners needed HTTP action for request and builds destination URL path
// Return http method(string), endpoint api url(string), error or nil
func determineAPIAction(baseUri, version string, opType builder.OperationType) (string, string, error) {
	var httpMethod, apiEndpoint string

	// Validate API config, base URL and version of API
	if baseUri == "" {
		return httpMethod, apiEndpoint, errors.New("Gateway client's URL is empty in, API settings.")
	}

	if version == "" {
		return httpMethod, apiEndpoint, errors.New("Gateway client's Version is empty in, API settings.")
	}

	apiEndpoint = fmt.Sprintf("%s/v%s/%s", baseUri, version, opType)

	switch opType {
	case builder.SMS, builder.DMS_HOLD:
		httpMethod = "POST"
	default:
		return httpMethod, apiEndpoint, errors.New("Unknow operation type, can't determinets HTTP action.")
	}

	return httpMethod, apiEndpoint, nil
}

// buildHTTPRequest, accepts prepared body for HTTP
// Builds NewRequest from http package
func buildHTTPRequest(method, url string, payload *bytes.Buffer) (*http.Request, error) {
	// Build whole HTTP request with payload data
	newReq, err := http.NewRequest(method, url, payload)
	if err != nil {
		return nil, err
	}

	// Set default headers for new request
	newReq.Header.Set("Accept", "application/json")
	newReq.Header.Set("Content-type", "application/json")

	return newReq, nil
}

// parseResponse, parsing response to structure
func parseResponse(resp *http.Response) (interface{}, error){
	body, bodyErr := ioutil.ReadAll(resp.Body)
	if bodyErr != nil {
		return nil, errors.New(fmt.Sprintf("failed to read received body: %s ", bodyErr.Error()))
	}

	resp.Body.Close()

	// @TODO Refactor to dynamic structure determine
	var gwResp structures.ResponseSMS
	parseErr := json.Unmarshal(body, &gwResp)

	if parseErr != nil {
		return nil, errors.New(fmt.Sprintf("failed to read received body: %s ", bodyErr.Error()))
	}

	return gwResp, nil
}