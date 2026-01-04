package yandexgpt

import "fmt"

type YandexGPTError struct {
	Message string
	Err     error
}

func (e *YandexGPTError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("%s: %v", e.Message, e.Err)
	}
	return e.Message
}

func (e *YandexGPTError) Unwrap() error {
	return e.Err
}

type AuthenticationError struct {
	YandexGPTError
}

func NewAuthenticationError(message string, err error) *AuthenticationError {
	return &AuthenticationError{
		YandexGPTError: YandexGPTError{
			Message: message,
			Err:     err,
		},
	}
}

type APIError struct {
	YandexGPTError
	StatusCode int
}

func NewAPIError(message string, statusCode int, err error) *APIError {
	return &APIError{
		YandexGPTError: YandexGPTError{
			Message: message,
			Err:     err,
		},
		StatusCode: statusCode,
	}
}

func (e *APIError) Error() string {
	if e.StatusCode > 0 {
		return fmt.Sprintf("API error (code %d): %s", e.StatusCode, e.YandexGPTError.Error())
	}
	return e.YandexGPTError.Error()
}
