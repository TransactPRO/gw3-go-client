package tprogateway

import (
	"testing"
)

func TestNewGateway(t *testing.T) {
	_, err := NewGateway(22, "rg342QZSUaWzKHoCc5slyMGdAITk9LfR")
	if err != nil {
		t.Error(err)
	}
}

func TestNewGatewayIncorrectAccountID(t *testing.T) {
	_, err := NewGateway(0, "rg342QZSUaWzKHoCc5slyMGdAITk9LfR")
	if err == nil {
		t.Error(err)
	}
}

func TestNewGatewayIncorrectSecretKey(t *testing.T) {
	_, err := NewGateway(11, "")
	if err == nil {
		t.Error(err)
	}
}