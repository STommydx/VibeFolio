package handlers

import (
	"fmt"
	"net/http"

	"github.com/a-h/templ"
	"go.uber.org/zap"

	"github.com/stommydx/VibeFolio/backend/src/ui/pages"
)

// ProfileHandlers handles profile-related HTTP requests for UI
type ProfileHandlers struct {
	logger *zap.Logger
}

// NewProfileHandlers creates new profile handlers
func NewProfileHandlers(logger *zap.Logger) *ProfileHandlers {
	return &ProfileHandlers{
		logger: logger,
	}
}

// ListProfilesHandler handles the GET /profiles endpoint for UI
func (h *ProfileHandlers) ListProfilesHandler(w http.ResponseWriter, r *http.Request) {
	// In a real implementation, this would:
	// 1. Get the authenticated user
	// 2. Fetch profiles from the service
	// 3. Render the profile list page
	
	// For now, we'll just render a mock page
	h.logger.Info("Rendering profile list page")
	
	// Mock profiles
	profiles := []pages.Profile{
		{ID: "1", Name: "John Doe"},
		{ID: "2", Name: "Jane Smith"},
	}
	
	component := pages.ProfileListPage(profiles)
	
	// Set content type
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	
	// Render the component
	if err := component.Render(r.Context(), w); err != nil {
		http.Error(w, "Failed to render page", http.StatusInternalServerError)
		return
	}
}

// GetProfileHandler handles the GET /profiles/{profileId} endpoint for UI
func (h *ProfileHandlers) GetProfileHandler(w http.ResponseWriter, r *http.Request) {
	// In a real implementation, this would:
	// 1. Get the authenticated user
	// 2. Fetch the profile from the service
	// 3. Fetch the profile sections from the service
	// 4. Render the profile detail page
	
	// For now, we'll just render a mock page
	h.logger.Info("Rendering profile detail page")
	
	// Mock profile
	profile := pages.ProfileDetail{
		ID:       "1",
		Name:     "John Doe",
		Title:    "Software Engineer",
		Summary:  "Experienced software developer with 5 years of experience",
		IsPublic: true,
	}
	
	// Mock sections
	sections := []pages.ProfileSection{
		{
			ID:      "1",
			Type:    "summary",
			Title:   "Professional Summary",
			Content: "Highly skilled software engineer...",
		},
		{
			ID:      "2",
			Type:    "experience",
			Title:   "Work Experience",
			Content: "Software Engineer at Tech Corp...",
		},
	}
	
	component := pages.ProfileDetailPage(profile, sections)
	
	// Set content type
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	
	// Render the component
	if err := component.Render(r.Context(), w); err != nil {
		http.Error(w, "Failed to render page", http.StatusInternalServerError)
		return
	}
}

// EditProfileHandler handles the GET /profiles/{profileId}/edit endpoint for UI
func (h *ProfileHandlers) EditProfileHandler(w http.ResponseWriter, r *http.Request) {
	// In a real implementation, this would:
	// 1. Get the authenticated user
	// 2. Fetch the profile from the service
	// 3. Render the profile edit page
	
	// For now, we'll just render a mock page
	h.logger.Info("Rendering profile edit page")
	
	// Mock profile
	profile := pages.ProfileDetail{
		ID:       "1",
		Name:     "John Doe",
		Title:    "Software Engineer",
		Summary:  "Experienced software developer with 5 years of experience",
		IsPublic: true,
	}
	
	component := pages.ProfileEditPage(profile)
	
	// Set content type
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	
	// Render the component
	if err := component.Render(r.Context(), w); err != nil {
		http.Error(w, "Failed to render page", http.StatusInternalServerError)
		return
	}
}

// UpdateProfileHandler handles the PUT /profiles/{profileId} endpoint for UI
func (h *ProfileHandlers) UpdateProfileHandler(w http.ResponseWriter, r *http.Request) {
	// In a real implementation, this would:
	// 1. Get the authenticated user
	// 2. Parse form data
	// 3. Update the profile using the service
	// 4. Redirect to the profile detail page
	
	// For now, we'll just redirect back to the profile page
	h.logger.Info("Updating profile")
	
	// Redirect to profile page
	profileID := "1" // This would be extracted from the URL in a real implementation
	http.Redirect(w, r, "/profiles/"+profileID, http.StatusSeeOther)
}