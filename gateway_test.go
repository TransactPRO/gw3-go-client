package tprogateway

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewGatewayClient(t *testing.T) {
	_, err := NewGatewayClient("3383e58e-9cde-4ffa-85cf-81cd25b2423e", "SecKey")
	assert.NoError(t, err)
}

func TestNewGatewayClientIncorrectData(t *testing.T) {
	_, err := NewGatewayClient("", "SecKey")
	assert.EqualError(t, err, "GUID can't be empty. It's required for merchant authorization")

	_, err2 := NewGatewayClient("3383e58e-9cde-4ffa-85cf-81cd25b2423e", "")
	assert.EqualError(t, err2, "secret key can't be empty. It's required for merchant authorization")
}

func TestNewGatewayClientForSession(t *testing.T) {
	_, err := NewGatewayClientForSession("3383e58e-9cde-4ffa-85cf-81cd25b2423e", "SecKey", "some-session-id")
	assert.NoError(t, err)
}

func TestNewGatewayClientForSessionIncorrectData(t *testing.T) {
	_, err := NewGatewayClientForSession("", "SecKey", "some-session-id")
	assert.EqualError(t, err, "GUID can't be empty. It's required for merchant authorization")

	_, err2 := NewGatewayClientForSession("3383e58e-9cde-4ffa-85cf-81cd25b2423e", "", "some-session-id")
	assert.EqualError(t, err2, "secret key can't be empty. It's required for merchant authorization")

	_, err3 := NewGatewayClientForSession("3383e58e-9cde-4ffa-85cf-81cd25b2423e", "SecKey", "")
	assert.EqualError(t, err3, "SessionID can't be empty. Session authorization means non-empty session")
}
