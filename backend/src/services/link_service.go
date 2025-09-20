package services

import (
	"context"
	"time"

	"github.com/google/uuid"
	"go.uber.org/zap"
)

// LinkService handles external link-related business logic
type LinkService struct {
	logger *zap.Logger
}

// ExternalLink represents an external link associated with a profile
type ExternalLink struct {
	ID        string    `json:"id"`
	ProfileID string    `json:"profile_id"`
	URL       string    `json:"url"`
	Title     string    `json:"title"`
	AddedAt   time.Time `json:"added_at"`
}

// NewLinkService creates a new LinkService
func NewLinkService(logger *zap.Logger) *LinkService {
	return &LinkService{
		logger: logger,
	}
}

// CreateLink creates a new external link
func (s *LinkService) CreateLink(ctx context.Context, profileID, url, title string) (*ExternalLink, error) {
	// In a real implementation, this would interact with the database
	// For now, we'll just return a mock external link
	link := &ExternalLink{
		ID:        uuid.New().String(),
		ProfileID: profileID,
		URL:       url,
		Title:     title,
		AddedAt:   time.Now(),
	}

	s.logger.Info("Created external link", zap.String("link_id", link.ID))
	return link, nil
}

// GetLink retrieves an external link by ID
func (s *LinkService) GetLink(ctx context.Context, linkID string) (*ExternalLink, error) {
	// In a real implementation, this would fetch from the database
	// For now, we'll just return a mock external link
	link := &ExternalLink{
		ID:        linkID,
		ProfileID: "mock_profile_id",
		URL:       "https://example.com",
		Title:     "Example Project",
		AddedAt:   time.Now().Add(-24 * time.Hour),
	}

	s.logger.Info("Retrieved external link", zap.String("link_id", link.ID))
	return link, nil
}

// UpdateLink updates an external link
func (s *LinkService) UpdateLink(ctx context.Context, linkID, url, title string) (*ExternalLink, error) {
	// In a real implementation, this would update the database
	// For now, we'll just return a mock external link
	link := &ExternalLink{
		ID:        linkID,
		ProfileID: "mock_profile_id",
		URL:       url,
		Title:     title,
		AddedAt:   time.Now().Add(-24 * time.Hour),
	}

	s.logger.Info("Updated external link", zap.String("link_id", link.ID))
	return link, nil
}

// DeleteLink deletes an external link
func (s *LinkService) DeleteLink(ctx context.Context, linkID string) error {
	// In a real implementation, this would delete from the database
	s.logger.Info("Deleted external link", zap.String("link_id", linkID))
	return nil
}

// ListLinks lists all external links for a profile
func (s *LinkService) ListLinks(ctx context.Context, profileID string) ([]*ExternalLink, error) {
	// In a real implementation, this would fetch from the database
	// For now, we'll just return mock external links
	links := []*ExternalLink{
		{
			ID:        uuid.New().String(),
			ProfileID: profileID,
			URL:       "https://github.com/example",
			Title:     "GitHub Profile",
			AddedAt:   time.Now().Add(-24 * time.Hour),
		},
		{
			ID:        uuid.New().String(),
			ProfileID: profileID,
			URL:       "https://linkedin.com/in/example",
			Title:     "LinkedIn Profile",
			AddedAt:   time.Now().Add(-48 * time.Hour),
		},
	}

	s.logger.Info("Listed external links", zap.String("profile_id", profileID), zap.Int("count", len(links)))
	return links, nil
}