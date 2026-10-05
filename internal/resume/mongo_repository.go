// Package resume defines domain data models, repository interfaces, and services for resume data.
package resume

import (
	"context"

	"github.com/nawaphonOHM/whatever/v2/pkg/mongodb"
)

// MongoRepository implements Repository using MongoDB client.
type MongoRepository struct {
	client *mongodb.Client
}

// NewMongoRepository creates a new MongoRepository instance.
func NewMongoRepository(client *mongodb.Client) *MongoRepository {
	return &MongoRepository{client: client}
}

// GetLatestDetails retrieves the latest details document.
func (r *MongoRepository) GetLatestDetails(ctx context.Context) (*Details, error) {
	return queryLatest[Details](ctx, r.client, "details")
}

// GetLatestEducations retrieves the latest education document.
func (r *MongoRepository) GetLatestEducations(ctx context.Context) (*Education, error) {
	return queryLatest[Education](ctx, r.client, "education")
}

// GetLatestExperiences retrieves the latest experience document.
func (r *MongoRepository) GetLatestExperiences(ctx context.Context) ([]*ExperienceItem, error) {
	res, err := queryLatest[[]*ExperienceItem](ctx, r.client, "experience")
	if err != nil {
		return nil, err
	}
	return *res, nil
}

// GetLatestLinks retrieves the latest links document.
func (r *MongoRepository) GetLatestLinks(ctx context.Context) ([]*LinkItem, error) {
	res, err := queryLatest[[]*LinkItem](ctx, r.client, "links")
	if err != nil {
		return nil, err
	}
	return *res, nil
}

// GetLatestNames retrieves the latest name document.
func (r *MongoRepository) GetLatestNames(ctx context.Context) (string, error) {
	res, err := queryLatest[string](ctx, r.client, "name")
	if err != nil {
		return "", err
	}
	return *res, nil
}

// GetLatestTitles retrieves the latest titles document.
func (r *MongoRepository) GetLatestTitles(ctx context.Context) ([]string, error) {
	res, err := queryLatest[[]string](ctx, r.client, "title")
	if err != nil {
		return nil, err
	}
	return *res, nil
}
