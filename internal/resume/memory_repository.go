// Package resume defines domain data models, repository interfaces, and services for resume data.
package resume

import (
	"context"
)

// MemoryRepository provides an in-memory mock implementation of Repository for testing.
type MemoryRepository struct {
	Details    *Details
	Education  *Education
	Err        error
	Name       string
	Links      []*LinkItem
	Titles     []string
	Experience []*ExperienceItem
}

// NewMemoryRepository creates an empty in-memory repository.
func NewMemoryRepository() *MemoryRepository {
	return &MemoryRepository{}
}

func (m *MemoryRepository) checkErr(ctx context.Context) error {
	if err := ctx.Err(); err != nil {
		return err
	}
	return m.Err
}

// GetLatestDetails returns the in-memory details or ErrNotFound.
func (m *MemoryRepository) GetLatestDetails(ctx context.Context) (*Details, error) {
	if err := m.checkErr(ctx); err != nil {
		return nil, err
	}
	if m.Details == nil {
		return nil, ErrNotFound
	}
	return m.Details, nil
}

// GetLatestEducation returns the in-memory education or ErrNotFound.
func (m *MemoryRepository) GetLatestEducation(ctx context.Context) (*Education, error) {
	if err := m.checkErr(ctx); err != nil {
		return nil, err
	}
	if m.Education == nil {
		return nil, ErrNotFound
	}
	return m.Education, nil
}

// GetLatestExperience returns the in-memory experience or ErrNotFound.
func (m *MemoryRepository) GetLatestExperience(ctx context.Context) ([]*ExperienceItem, error) {
	if err := m.checkErr(ctx); err != nil {
		return nil, err
	}
	if m.Experience == nil {
		return nil, ErrNotFound
	}
	return m.Experience, nil
}

// GetLatestLinks returns the in-memory links or ErrNotFound.
func (m *MemoryRepository) GetLatestLinks(ctx context.Context) ([]*LinkItem, error) {
	if err := m.checkErr(ctx); err != nil {
		return nil, err
	}
	if m.Links == nil {
		return nil, ErrNotFound
	}
	return m.Links, nil
}

// GetLatestName returns the in-memory name or ErrNotFound.
func (m *MemoryRepository) GetLatestName(ctx context.Context) (string, error) {
	if err := m.checkErr(ctx); err != nil {
		return "", err
	}
	if m.Name == "" {
		return "", ErrNotFound
	}
	return m.Name, nil
}

// GetLatestTitles returns the in-memory titles or ErrNotFound.
func (m *MemoryRepository) GetLatestTitles(ctx context.Context) ([]string, error) {
	if err := m.checkErr(ctx); err != nil {
		return nil, err
	}
	if m.Titles == nil {
		return nil, ErrNotFound
	}
	return m.Titles, nil
}
