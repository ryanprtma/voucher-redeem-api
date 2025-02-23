package _test

import (
	"testing"

	"voucher-redeem-api/src/Commons/exceptions"
)

func TestInvariantError(t *testing.T) {
	message := "an error occurs"

	invariantError := exceptions.NewInvariantError(message)

	if invariantError.StatusCode != 400 {
		t.Errorf("Expected status code 400, got: %d", invariantError.StatusCode)
	}

	if invariantError.Message != message {
		t.Errorf("Expected message '%s', got: %s", message, invariantError.Error())
	}
}
