package tprogateway

import (
	"errors"
	"net/http"
	"bitbucket.transactpro.lv/tls/gw3-go-client/request"
)

// Default API settings
const (
	dAPIUri = "http://uriel.sk.fpngw3.env"
	dAPIVersion = "3.0"
)

type (
	// API, Transact PRO gateway endpoint and version
	API struct {
		APIUri     string
		APIVersion string
	}

	// GatewayClient
	// Transact Pro Gateway integration client
	// Tt represents a REST API Client
	GatewayClient struct {
		client      *http.Client
		APIConfig   *API
		AuthData    *request.AuthData
	}
)

// NewGateway, new instance of prepared gateway client
func NewGateway(AccountID int, SecretKey string) (*GatewayClient, error) {
	if AccountID == 0 {
		return nil, errors.New("Account ID can not be 0, please use the given ID from Transact Pro.")
	}

	if SecretKey == "" {
		return nil, errors.New("Secret key can't be empty. It's needed for merchant authorization.")
	}

	return &GatewayClient {
		client: &http.Client{},
		// @TODO Add way to change API config: yml, json, env
		APIConfig: &API{
			dAPIUri,
			dAPIVersion,
		},
		AuthData: &request.AuthData{AccountID: AccountID, SecretKey:SecretKey},
	}, nil
}