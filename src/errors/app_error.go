package errors

import (
	"fmt"
	"net/http"
)

// AppError represents a structured application error
type AppError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Details string `json:"details,omitempty"`
}

// Error implements the error interface
func (e *AppError) Error() string {
	if e.Details != "" {
		return fmt.Sprintf("%s: %s", e.Message, e.Details)
	}
	return e.Message
}

// HTTPStatusCode returns the HTTP status code for this error
func (e *AppError) HTTPStatusCode() int {
	return e.Code
}

// NewAppError creates a new AppError
func NewAppError(code int, message, details string) *AppError {
	return &AppError{
		Code:    code,
		Message: message,
		Details: details,
	}
}

// NewBadRequestError creates a 400 Bad Request error
func NewBadRequestError(message, details string) *AppError {
	return NewAppError(http.StatusBadRequest, message, details)
}

// NewNotFoundError creates a 404 Not Found error
func NewNotFoundError(message, details string) *AppError {
	return NewAppError(http.StatusNotFound, message, details)
}

// NewInternalServerError creates a 500 Internal Server Error
func NewInternalServerError(message, details string) *AppError {
	return NewAppError(http.StatusInternalServerError, message, details)
}

// NewServiceUnavailableError creates a 503 Service Unavailable error
func NewServiceUnavailableError(message, details string) *AppError {
	return NewAppError(http.StatusServiceUnavailable, message, details)
}
