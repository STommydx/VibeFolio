package controllers

import (
	"context"
	"net/http"

	"github.com/STommydx/VibeFolio/src/models"
	"github.com/STommydx/VibeFolio/src/services"
	"github.com/danielgtaylor/huma/v2"
)

// HealthController handles health check requests
type HealthController struct {
	healthService *services.HealthService
}

// NewHealthController creates a new HealthController instance
func NewHealthController(healthService *services.HealthService) *HealthController {
	return &HealthController{
		healthService: healthService,
	}
}

// RegisterRoutes registers the health check routes with the Huma API
func (h *HealthController) RegisterRoutes(api huma.API) {
	huma.Register(api, huma.Operation{
		OperationID: "get-health",
		Summary:     "Health check endpoint",
		Method:      http.MethodGet,
		Path:        "/healthz",
		Description: "Returns the health status of the service",
	}, h.GetHealth)
}

// GetHealthInput represents the input for the health check endpoint
type GetHealthInput struct{}

// GetHealthOutput represents the output for the health check endpoint
type GetHealthOutput struct {
	Body *models.HealthStatus
}

// GetHealth handles GET /healthz requests
func (h *HealthController) GetHealth(ctx context.Context, input *GetHealthInput) (*GetHealthOutput, error) {
	healthStatus := h.healthService.GetHealthStatus()

	output := &GetHealthOutput{
		Body: healthStatus,
	}

	// Set the appropriate status code based on health status
	if healthStatus.Status == "unhealthy" {
		return output, huma.NewError(http.StatusServiceUnavailable, "Service is unhealthy")
	}

	return output, nil
}
