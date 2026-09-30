// Package resume defines domain data models, repository interfaces, and services for resume data.
package resume

import (
	"context"
)

// Service defines the business logic interface for resume operations.
type Service interface {
	// GetDetails retrieves the latest personal details.
	GetDetails(context.Context) (*Details, error)
	// GetEducation retrieves the latest education background.
	GetEducation(context.Context) (*Education, error)
	// GetExperience retrieves the latest work experiences.
	GetExperience(context.Context) ([]*ExperienceItem, error)
	// GetLinks retrieves the latest external links.
	GetLinks(context.Context) ([]*LinkItem, error)
	// GetName retrieves the latest display name.
	GetName(context.Context) (string, error)
	// GetTitles retrieves the latest list of titles.
	GetTitles(context.Context) ([]string, error)
}

type resumeService struct {
	repo Repository
}

// NewService creates a new resume Service instance.
func NewService(repo Repository) Service {
	return &resumeService{repo: repo}
}

func (s *resumeService) GetDetails(ctx context.Context) (*Details, error) {
	return s.repo.GetLatestDetails(ctx)
}

func (s *resumeService) GetEducation(ctx context.Context) (*Education, error) {
	return s.repo.GetLatestEducation(ctx)
}

func (s *resumeService) GetExperience(ctx context.Context) ([]*ExperienceItem, error) {
	return s.repo.GetLatestExperience(ctx)
}

func (s *resumeService) GetLinks(ctx context.Context) ([]*LinkItem, error) {
	return s.repo.GetLatestLinks(ctx)
}

func (s *resumeService) GetName(ctx context.Context) (string, error) {
	return s.repo.GetLatestName(ctx)
}

func (s *resumeService) GetTitles(ctx context.Context) ([]string, error) {
	return s.repo.GetLatestTitles(ctx)
}
