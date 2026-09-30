// Package resume defines domain data models, repository interfaces, and services for resume data.
package resume

import (
	"context"
	"errors"
)

// ErrNotFound indicates that the requested resume resource does not exist.
var ErrNotFound = errors.New("resource not found")

// Repository defines data access methods for resume domain resources.
type Repository interface {
	// GetLatestDetails retrieves the latest details document.
	GetLatestDetails(context.Context) (*Details, error)
	// GetLatestEducation retrieves the latest education document.
	GetLatestEducation(context.Context) (*Education, error)
	// GetLatestExperience retrieves the latest experience document.
	GetLatestExperience(context.Context) ([]*ExperienceItem, error)
	// GetLatestLinks retrieves the latest links document.
	GetLatestLinks(context.Context) ([]*LinkItem, error)
	// GetLatestName retrieves the latest name document.
	GetLatestName(context.Context) (string, error)
	// GetLatestTitles retrieves the latest titles document.
	GetLatestTitles(context.Context) ([]string, error)
}
