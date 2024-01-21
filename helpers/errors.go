package helpers

import (
	"fmt"
	"net/http"
)

type HTTPError struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func NewHTTPError(status int, message string) *HTTPError {
	return &HTTPError{
		Status:  status,
		Message: message,
	}
}

func (e *HTTPError) Error() string {
	return fmt.Sprintf("HTTP error: %d - %s", e.Status, e.Message)
}

func BadRequest(message string) *HTTPError {
	return NewHTTPError(http.StatusBadRequest, message)
}

func InternalServerError(message string) *HTTPError {
	return NewHTTPError(http.StatusInternalServerError, message)
}
