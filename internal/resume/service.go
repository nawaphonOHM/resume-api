// Package resume defines domain data models, repository interfaces, and services for resume data.
package resume

import (
	"context"
)

// Service defines the business logic interface for resume operations.
type Service interface {
	// GetDetails retrieves the latest personal details.
	GetDetails(context.Context) (*Details, error)
	// GetEducations retrieves the latest education background.
	GetEducations(context.Context) (*Education, error)
	// GetExperiences retrieves the latest work experiences.
	GetExperiences(context.Context) ([]*ExperienceItem, error)
	// GetLinks retrieves the latest external links.
	GetLinks(context.Context) ([]*LinkItem, error)
	// GetNames retrieves the latest display name.
	GetNames(context.Context) (string, error)
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

func (s *resumeService) GetEducations(ctx context.Context) (*Education, error) {
	return s.repo.GetLatestEducations(ctx)
}

func (s *resumeService) GetExperiences(ctx context.Context) ([]*ExperienceItem, error) {
	return s.repo.GetLatestExperiences(ctx)
}

func (s *resumeService) GetLinks(ctx context.Context) ([]*LinkItem, error) {
	return s.repo.GetLatestLinks(ctx)
}

func (s *resumeService) GetNames(ctx context.Context) (string, error) {
	return s.repo.GetLatestNames(ctx)
}

func (s *resumeService) GetTitles(ctx context.Context) ([]string, error) {
	return s.repo.GetLatestTitles(ctx)
}
