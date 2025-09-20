package api

import (
	"context"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

// ResumeHandler handles resume-related HTTP requests
type ResumeHandler struct {
	logger *zap.Logger
}

// NewResumeHandler creates a new ResumeHandler
func NewResumeHandler(logger *zap.Logger) *ResumeHandler {
	return &ResumeHandler{
		logger: logger,
	}
}

// RegisterRoutes registers the resume routes
func (h *ResumeHandler) RegisterRoutes(api huma.API) {
	huma.Register(api, huma.Operation{
		OperationID: "generate-resume",
		Method:      http.MethodPost,
		Path:        "/api/profiles/{profileId}/resume",
		Summary:     "Generate resume PDF",
		Description: "Generate a PDF resume from the profile data",
	}, h.generateResume)
}

// ResumeTemplate represents a LaTeX template for resume generation
type ResumeTemplate struct {
	ID          string `json:"id" format:"uuid" doc:"Template ID"`
	Name        string `json:"name" doc:"Name of the template"`
	Description string `json:"description" doc:"Description of the template"`
	IsDefault   bool   `json:"is_default" doc:"Whether this is the default template"`
	CreatedAt   string `json:"created_at" format:"date-time" doc:"When the template was created"`
	UpdatedAt   string `json:"updated_at" format:"date-time" doc:"When the template was last updated"`
}

// GeneratedResume represents a generated resume PDF
type GeneratedResume struct {
	ID          string `json:"id" format:"uuid" doc:"Resume ID"`
	ProfileID   string `json:"profile_id" format:"uuid" doc:"Profile ID"`
	TemplateID  string `json:"template_id" format:"uuid" doc:"Template ID"`
	GeneratedAt string `json:"generated_at" format:"date-time" doc:"When the resume was generated"`
	FileSize    int    `json:"file_size" doc:"Size of the PDF in bytes"`
}

// GenerateResumeRequest represents the request for the generate resume endpoint
type GenerateResumeRequest struct {
	ProfileID string `path:"profileId" format:"uuid" doc:"Profile ID"`
	Body      struct {
		TemplateID string `json:"template_id,omitempty" format:"uuid" doc:"Template ID to use for generation (optional)"`
	} `json:"body"`
}

// GenerateResumeResponse represents the response for the generate resume endpoint
type GenerateResumeResponse struct {
	Body struct {
		JobID   string `json:"job_id" format:"uuid" doc:"Job ID for tracking generation progress"`
		Status  string `json:"status" enum:"queued,processing,completed,failed" doc:"Status of the generation job"`
		Message string `json:"message" doc:"Status message"`
	} `json:"body"`
}

// generateResume handles the POST /api/profiles/{profileId}/resume endpoint
func (h *ResumeHandler) generateResume(ctx context.Context, input *GenerateResumeRequest) (*GenerateResumeResponse, error) {
	// In a real implementation, this would:
	// 1. Validate that the user has access to this profile
	// 2. Call the resume service to start the generation process
	// 3. Return a job ID for tracking progress
	
	// For now, we'll just return a mock response
	h.logger.Info("Generating resume for profile", zap.String("profile_id", input.ProfileID))
	
	return &GenerateResumeResponse{
		Body: struct {
			JobID   string `json:"job_id" format:"uuid" doc:"Job ID for tracking generation progress"`
			Status  string `json:"status" enum:"queued,processing,completed,failed" doc:"Status of the generation job"`
			Message string `json:"message" doc:"Status message"`
		}{
			JobID:   uuid.New().String(),
			Status:  "queued",
			Message: "Resume generation started",
		},
	}, nil
}