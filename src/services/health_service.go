package services

import (
	"time"

	"github.com/STommydx/VibeFolio/src/models"
)

// HealthService provides health check functionality
type HealthService struct {
	version string
}

// NewHealthService creates a new HealthService instance
func NewHealthService(version string) *HealthService {
	return &HealthService{
		version: version,
	}
}

// GetHealthStatus returns the current health status of the service
func (h *HealthService) GetHealthStatus() *models.HealthStatus {
	return &models.HealthStatus{
		Status:    "healthy",
		Timestamp: time.Now().UTC(),
		Version:   h.version,
		Details: map[string]interface{}{
			"service": "VibeFolio Health Check",
		},
	}
}

// IsHealthy returns true if the service is healthy
func (h *HealthService) IsHealthy() bool {
	// In a real implementation, this would check various components
	// For now, we'll always return true
	return true
}
