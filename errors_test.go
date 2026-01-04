package yandexgpt

import (
	"errors"
	"testing"
)

func TestAuthenticationError(t *testing.T) {
	baseErr := errors.New("base error")
	authErr := NewAuthenticationError("authentication failed", baseErr)

	if authErr == nil {
		t.Fatal("Expected error, got nil")
	}

	if authErr.Message != "authentication failed" {
		t.Errorf("Expected message 'authentication failed', got '%s'", authErr.Message)
	}

	if !errors.Is(authErr, baseErr) {
		t.Error("Expected error to wrap base error")
	}

	errorString := authErr.Error()
	if errorString == "" {
		t.Error("Expected non-empty error string")
	}
}

func TestAPIError(t *testing.T) {
	baseErr := errors.New("base error")
	apiErr := NewAPIError("API request failed", 400, baseErr)

	if apiErr == nil {
		t.Fatal("Expected error, got nil")
	}

	if apiErr.Message != "API request failed" {
		t.Errorf("Expected message 'API request failed', got '%s'", apiErr.Message)
	}

	if apiErr.StatusCode != 400 {
		t.Errorf("Expected status code 400, got %d", apiErr.StatusCode)
	}

	if !errors.Is(apiErr, baseErr) {
		t.Error("Expected error to wrap base error")
	}

	errorString := apiErr.Error()
	if errorString == "" {
		t.Error("Expected non-empty error string")
	}
}

func TestYandexGPTError(t *testing.T) {
	err := &YandexGPTError{
		Message: "test error",
		Err:     nil,
	}

	if err.Error() != "test error" {
		t.Errorf("Expected 'test error', got '%s'", err.Error())
	}

	baseErr := errors.New("base error")
	err = &YandexGPTError{
		Message: "test error",
		Err:     baseErr,
	}

	if err.Unwrap() != baseErr {
		t.Error("Expected Unwrap to return base error")
	}
}
