package _test

import (
	"testing"
	"voucher-redeem-api/src/Commons/exceptions"
)

func TestNewNotFoundError(t *testing.T) {
	notFoundError := exceptions.NewNotFoundError("not found!")

	if notFoundError.StatusCode != 404 {
		t.Errorf("Expected status code 404, got %d", notFoundError.StatusCode)
	}

	if notFoundError.Message != "not found!" {
		t.Errorf("Expected message 'not found!', got %s", notFoundError.Error())
	}
}
