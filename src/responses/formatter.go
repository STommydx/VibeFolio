package responses

import (
	"encoding/json"
	"net/http"

	"github.com/STommydx/VibeFolio/src/errors"
)

// JSONResponse formats a successful JSON response
func JSONResponse(w http.ResponseWriter, data interface{}, statusCode int) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	return json.NewEncoder(w).Encode(data)
}

// ErrorResponse formats an error response
func ErrorResponse(w http.ResponseWriter, appErr *errors.AppError) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(appErr.HTTPStatusCode())

	return json.NewEncoder(w).Encode(appErr)
}

// HealthResponse formats a health check response
func HealthResponse(w http.ResponseWriter, data interface{}) error {
	// For health checks, we'll use 200 for healthy and 503 for unhealthy
	// This logic would be implemented based on the health status
	return JSONResponse(w, data, http.StatusOK)
}
