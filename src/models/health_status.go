package models

import (
	"time"
)

// HealthStatus represents the health status of the service
type HealthStatus struct {
	// Status is the current health status of the service (required)
	// Valid values: "healthy", "unhealthy"
	Status string `json:"status" validate:"required,oneof=healthy unhealthy"`

	// Timestamp is the ISO8601 timestamp when the health check was performed (required)
	Timestamp time.Time `json:"timestamp" validate:"required"`

	// Version is the version of the service (optional)
	Version string `json:"version,omitempty"`

	// Details contains additional details about the health status (optional)
	Details map[string]interface{} `json:"details,omitempty"`
}
