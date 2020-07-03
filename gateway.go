// Package tprogateway provide ability to make requests to Transact Pro Gateway API v3.
package tprogateway

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/TransactPRO/gw3-go-client/operations"
	"github.com/TransactPRO/gw3-go-client/structures"
)

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

	// AuthData merchant authorization structure fields used in operation request
	authData struct {
		ObjectGUID string `json:"-"`
		SecretKey  string `json:"-"`
		SessionID  string `json:"session-id,omitempty"`
	}

	// GatewayClient represents REST API client
	GatewayClient struct {
		API        *confAPI
		Auth       *authData
		HTTPClient http.Client
	}

	// GenericRequest describes general request data structure
	GenericRequest struct {
		Auth       *authData   `json:"auth-data,omitempty"`
		Data       interface{} `json:"data,omitempty"`
		FilterData interface{} `json:"filter-data,omitempty"`
	}
)

// NewGatewayClient creates new instance of prepared gateway client structure
func NewGatewayClient(ObjectGUID, SecretKey string) (*GatewayClient, error) {
	if ObjectGUID == "" {
		return nil, errors.New("GUID can't be empty. It's required for merchant authorization")
	}

	if SecretKey == "" {
		return nil, errors.New("secret key can't be empty. It's required for merchant authorization")
	}

	return &GatewayClient{
		API:  &confAPI{BaseURI: dAPIBaseURI, Version: dAPIVersion},
		Auth: &authData{ObjectGUID: ObjectGUID, SecretKey: SecretKey},
	}, nil
}

// NewGatewayClientForSession creates new instance of prepared gateway client structure
// Should be used when active session is available
func NewGatewayClientForSession(ObjectGUID, SecretKey, SessionID string) (response *GatewayClient, err error) {
	if SessionID == "" {
		return nil, errors.New("SessionID can't be empty. Session authorization means non-empty session")
	}

	if response, err = NewGatewayClient(ObjectGUID, SecretKey); err == nil {
		response.Auth.SessionID = SessionID
	}

	return
}

// OperationBuilder method, returns builder for needed operation, like SMS, Reversal, even exploring transactions such as Refund ExploringHistory
func (gc *GatewayClient) OperationBuilder() *operations.Builder {
	return &operations.Builder{}
}

// NewRequest method, send HTTP request to Transact Pro API
// GatewayResponse may be non-nil in case of error if a response payload was read
// but some validation after failed (like digest verification)
func (gc *GatewayClient) NewRequest(opData structures.OperationRequestInterface) (*structures.GatewayResponse, error) {
	// Build whole payload structure with nested data bundles
	rawReqData := &GenericRequest{}
	rawReqData.Auth = gc.Auth
	if opData.GetOperationType() == structures.Report {
		rawReqData.FilterData = opData
	} else {
		rawReqData.Data = opData
	}

	// Get prepared structure of json byte array
	var bufPayload *bytes.Buffer
	if opData.GetHTTPMethod() != http.MethodGet {
		var bufErr error
		if bufPayload, bufErr = prepareJSONPayload(rawReqData); bufErr != nil {
			return nil, bufErr
		}
	} else {
		bufPayload = bytes.NewBuffer(nil)
	}

	// Get combined URL path for request to API
	requestURL, errURLPath := determineURL(gc, opData.GetOperationType())
	if errURLPath != nil {
		return nil, errURLPath
	}

	// Build correct HTTP request
	newReq, reqDigest, reqErr := buildHTTPRequest(rawReqData.Auth, opData.GetHTTPMethod(), requestURL, bufPayload)
	if reqErr != nil {
		return nil, reqErr
	}

	// Send HTTP request object
	resp, respErr := gc.HTTPClient.Do(newReq)
	if respErr != nil {
		return nil, respErr
	}
	defer func() { _ = resp.Body.Close() }()

	content, payloadErr := ioutil.ReadAll(resp.Body)
	if payloadErr != nil {
		return nil, payloadErr
	}

	gwResponse := structures.NewGatewayResponse(resp, content)
	if gwResponse.Successful() {
		var digestErr error
		gwResponse.Digest, digestErr = structures.NewResponseDigest(resp.Header.Get("Authorization"))
		if digestErr != nil {
			return gwResponse, digestErr
		}

		gwResponse.Digest.OriginalURI = reqDigest.URI
		gwResponse.Digest.OriginalCnonce = reqDigest.Cnonce
		gwResponse.Digest.Body = gwResponse.Payload
		digestErr = gwResponse.Digest.Verify(rawReqData.Auth.ObjectGUID, rawReqData.Auth.SecretKey)
		if digestErr != nil {
			return gwResponse, digestErr
		}
	}

	return gwResponse, nil
}

// prepareJSONPayload, validates\combines AuthData and Data struct to one big structure and converts to json(Marshal) to buffer
func prepareJSONPayload(rawReq *GenericRequest) (*bytes.Buffer, error) {
	// When payload ready, convert it to Json format
	bReqData, err := json.Marshal(&rawReq)
	if err != nil {
		return nil, err
	}

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
		return "", errors.New("gateway client's URL is empty in, API settings")
	}

	if gc.API.Version == "" {
		return "", errors.New("gateway client's Version is empty in, API settings")
	}

	// Try to get operation type from request data
	if opType == "" {
		return "", errors.New("operation type is empty. Problem in operation builder")
	}

	// AS example must be like: http://url.pay.com/v55.0/sms
	if strings.HasPrefix(string(opType), "http") {
		completeURL = string(opType)
	} else if opType == structures.Report {
		completeURL = fmt.Sprintf("%s/%s", gc.API.BaseURI, opType)
	} else {
		completeURL = fmt.Sprintf("%s/v%s/%s", gc.API.BaseURI, gc.API.Version, opType)
	}

	return completeURL, nil
}

// buildHTTPRequest, accepts prepared body for HTTP
// Builds NewRequest from http package
func buildHTTPRequest(auth *authData, method, requestURL string, payload *bytes.Buffer) (*http.Request, *structures.RequestDigest, error) {
	var err error

	var parsedURL *url.URL
	if parsedURL, err = url.Parse(requestURL); err != nil {
		return nil, nil, fmt.Errorf("incorrect URL: %s", err)
	}

	var requestDigest *structures.RequestDigest
	if requestDigest, err = structures.NewRequestDigest(auth.ObjectGUID, auth.SecretKey, parsedURL.Path, payload.Bytes()); err != nil {
		return nil, nil, err
	}

	var digest string
	if digest, err = requestDigest.CreateHeader(); err != nil {
		return nil, nil, err
	}

	// Build whole HTTP request with payload data
	var newReq *http.Request
	if newReq, err = http.NewRequest(method, requestURL, payload); err != nil {
		return nil, nil, err
	}

	// Set default headers for new request
	newReq.Header.Set("Authorization", digest)
	if method != http.MethodGet {
		newReq.Header.Set("Content-type", "application/json")
	}

	return newReq, requestDigest, nil
}
