// Package tprogateway provide ability to make requests to Transact Pro Gateway API v3.
package tprogateway

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"bitbucket.transactpro.lv/tls/gw3-go-client/operations"
	"bitbucket.transactpro.lv/tls/gw3-go-client/structures"
)

// @TODO Add logger with mode enabled\disabled, debug
// @TODO cover more code with test

// Default API settings
const (
	dAPIBaseURI = "https://api.sandbox.transactpro.io"
	dAPIVersion = "3.0"
)

type (
	// confAPI, endpoint config to Transact Pro system
	confAPI struct {
		// BaseURI typical host name with scheme. Example: http://some.host
		BaseURI string
		// Version is prefix in route path for url. Example: 42.1
		Version string
	}

	// AuthData merchant authorization structure fields used in operaion request
	authData struct {
		// Transact Pro Account ID
		AccountID int `json:"account-id"`
		// Transact Pro Merchant Password
		SecretKey string `json:"secret-key"`
	}

	// GatewayClient represents REST API client
	GatewayClient struct {
		API        *confAPI
		auth       *authData
		httpClient http.Client
	}

	// GenericRequest describes general request data structure
	GenericRequest struct {
		Auth interface{} `json:"auth-data"`
		Data interface{} `json:"data"`
	}
)

// NewGatewayClient new instance of prepared gateway client structure
func NewGatewayClient(AccountID int, SecretKey string) (*GatewayClient, error) {
	if AccountID == 0 {
		return nil, errors.New("Account ID can not be 0, please use the given ID from Transact Pro")
	}

	if SecretKey == "" {
		return nil, errors.New("Secret key can't be empty. It's needed for merchant authorization")
	}

	return &GatewayClient{
		API: &confAPI{
			BaseURI: dAPIBaseURI, Version: dAPIVersion},
		auth: &authData{
			AccountID: AccountID, SecretKey: SecretKey},
	}, nil
}

// OperationBuilder method, returns builder for needed operation, like SMS, Reversal, even exploring transactions such as Refund ExploringHistory
func (gc *GatewayClient) OperationBuilder() *operations.Builder {
	return &operations.Builder{}
}

// NewRequest method, send HTTP request to Transact Pro API
func (gc *GatewayClient) NewRequest(opData structures.OperationRequestInterface) (*http.Response, error) {
	// Build whole payload structure with nested data bundles
	rawReqData := &GenericRequest{}
	rawReqData.Auth = *gc.auth
	rawReqData.Data = opData

	// Get prepared structure of json byte array
	bufPayload, bufErr := prepareJSONPayload(rawReqData)
	if bufErr != nil {
		return nil, bufErr
	}

	// Get combined URL path for request to API
	url, errURLPath := determineURL(gc, opData.GetOperationType())
	if errURLPath != nil {
		return nil, errURLPath
	}

	// Build correct HTTP request
	newReq, reqErr := buildHTTPRequest(opData.GetHTTPMethod(), url, bufPayload)
	if reqErr != nil {
		return nil, reqErr
	}

	// Send HTTP request object
	resp, respErr := gc.httpClient.Do(newReq)
	if respErr != nil {
		return nil, respErr
	}

	return resp, nil
}

// prepareJSONPayload, validates\combines AuthData and Data struct to one big structure and converts to json(Marshal) to buffer
func prepareJSONPayload(rawReq *GenericRequest) (*bytes.Buffer, error) {
	// When payload ready, convert it to Json format
	bReqData, err := json.Marshal(&rawReq)
	if err != nil {
		return nil, err
	}

	// @TODO Debug print marshal body
	fmt.Println("RAW Marshal BODY " + string(bReqData))

	// Write json object to buffer
	buffer := bytes.NewBuffer(bReqData)

	return buffer, nil
}

// determineURL the full URL address to send request to Transact Pro API
func determineURL(gc *GatewayClient, opType structures.OperationType) (string, error) {
	// Complete URL for request
	var completeURL string

	// Validate API config, base URL and version of API
	if gc.API.BaseURI == "" {
		return "", errors.New("Gateway client's URL is empty in, API settings")
	}

	if gc.API.Version == "" {
		return "", errors.New("Gateway client's Version is empty in, API settings")
	}

	// Try to get operation type from request data
	if opType == "" {
		return "", errors.New("Operation type is empty. Problem in operation builder")
	}

	// AS example must be like: http://url.pay.com/v55.0/sms
	completeURL = fmt.Sprintf("%s/v%s/%s", gc.API.BaseURI, gc.API.Version, opType)

	// @TODO Debug print URL
	fmt.Println("REQUEST URL " + completeURL)

	return completeURL, nil
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

// ParseResponse method maps response to structure for given operation type
func (gc *GatewayClient) ParseResponse(resp *http.Response, opType structures.OperationType) (interface{}, error) {
	// @TODO Split response parser to additional packg
	parsedResp, parseErr := parseResponse(resp, opType)
	if parseErr != nil {
		return nil, parseErr
	}

	return parsedResp, nil
}

// parseResponse, parsing response to structure
func parseResponse(resp *http.Response, opType structures.OperationType) (interface{}, error) {
	defer resp.Body.Close()

	// Empty response body
	var responseBody interface{}

	body, bodyErr := ioutil.ReadAll(resp.Body)
	if bodyErr != nil {
		return nil, fmt.Errorf("Failed to read received body: %s ", bodyErr.Error())
	}

	// @TODO Map unauthorized response before map to any transactions

	// Determine operation response structure and parse it
	switch opType {
	case structures.ExploringStatus:
		var gwResp []structures.ExploringStatusResponse

		// Try parse response to transactions default structure
		// @TODO Debug print marshal body
		fmt.Println("RAW HTTP BODY " + string(body))

		parseErr := json.Unmarshal(body, &gwResp)
		if parseErr != nil {
			if bodyErr != nil {
				return nil, fmt.Errorf("Failed to unmarshal received http body: %s ", bodyErr.Error())
			}

			return nil, fmt.Errorf("Failed to unmarshal received http body, http body error unkown, unmarshal error: %s", parseErr)
		}

		// Assign parsed response structure to interface
		responseBody = gwResp
	case structures.ExploringResult:
		var gwResp []structures.ExploringResultResponse

		// Try parse response to transactions default structure
		// @TODO Debug print marshal body
		fmt.Println("RAW HTTP BODY " + string(body))

		parseErr := json.Unmarshal(body, &gwResp)
		if parseErr != nil {
			if bodyErr != nil {
				return nil, fmt.Errorf("Failed to unmarshal received http body: %s ", bodyErr.Error())
			}

			return nil, fmt.Errorf("Failed to unmarshal received http body, http body error unkown, unmarshal error: %s", parseErr)
		}

		// Assign parsed response structure to interface
		responseBody = gwResp
	case structures.ExploringHistory:
		var gwResp []structures.ExploringHistoryResponse

		// Try parse response to transactions default structure
		// @TODO Debug print marshal body
		fmt.Println("RAW HTTP BODY " + string(body))

		parseErr := json.Unmarshal(body, &gwResp)
		if parseErr != nil {
			if bodyErr != nil {
				return nil, fmt.Errorf("Failed to unmarshal received http body: %s ", bodyErr.Error())
			}

			return nil, fmt.Errorf("Failed to unmarshal received http body, http body error unkown, unmarshal error: %s", parseErr)
		}

		// Assign parsed response structure to interface
		responseBody = gwResp
	case structures.ExploringRecurrents:
		var gwResp []structures.ExploringRecurrentsResponse

		// Try parse response to transactions default structure
		// @TODO Debug print marshal body
		fmt.Println("RAW HTTP BODY " + string(body))

		parseErr := json.Unmarshal(body, &gwResp)
		if parseErr != nil {
			if bodyErr != nil {
				return nil, fmt.Errorf("Failed to unmarshal received http body: %s ", bodyErr.Error())
			}

			return nil, fmt.Errorf("Failed to unmarshal received http body, http body error unkown, unmarshal error: %s", parseErr)
		}

		// Assign parsed response structure to interface
		responseBody = gwResp
	case structures.ExploringRefunds:
		var gwResp []structures.ExploringRefundResponse

		// Try parse response to transactions default structure
		// @TODO Debug print marshal body
		fmt.Println("RAW HTTP BODY " + string(body))

		parseErr := json.Unmarshal(body, &gwResp)
		if parseErr != nil {
			if bodyErr != nil {
				return nil, fmt.Errorf("Failed to unmarshal received http body: %s ", bodyErr.Error())
			}

			return nil, fmt.Errorf("Failed to unmarshal received http body, http body error unkown, unmarshal error: %s", parseErr)
		}

		// Assign parsed response structure to interface
		responseBody = gwResp
	default:
		var gwResp structures.TransactionResponse

		// Try parse response to transactions default structure
		// @TODO Debug print marshal body
		fmt.Println("RAW HTTP BODY " + string(body))
		parseErr := json.Unmarshal(body, &gwResp)
		if parseErr != nil {
			if bodyErr != nil {
				return nil, fmt.Errorf("Failed to unmarshal received http body: %s ", bodyErr.Error())
			}

			return nil, fmt.Errorf("Failed to unmarshal received http body, http body error unkown, unmarshal error: %s", parseErr)
		}

		// Assign parsed response structure to interface
		responseBody = gwResp
	}

	// @TODO Think about return type, mb instead of using interface better to use strict structure types
	return responseBody, nil
}
