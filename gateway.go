// Package tprogateway, provide ability to make requests to Transact Pro Gateway API v3.
package tprogateway

import (
	"errors"

	"bitbucket.transactpro.lv/tls/gw3-go-client/operations"
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
		operation	*operations.Operation
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