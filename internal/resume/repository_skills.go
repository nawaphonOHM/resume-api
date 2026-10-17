// Package resume defines domain data models, repository interfaces, and services for resume data.
package resume

import (
	"context"
)

// GetLatestSkills retrieves the latest skills document from MongoDB.
func (r *MongoRepository) GetLatestSkills(ctx context.Context) ([]string, error) {
	res, err := queryLatest[[]string](ctx, r.client, collectionSkill)
	if err != nil {
		return nil, err
	}
	return *res, nil
}

// GetLatestSkills returns the in-memory skills or ErrNotFound.
func (m *MemoryRepository) GetLatestSkills(ctx context.Context) ([]string, error) {
	if err := m.checkErr(ctx); err != nil {
		return nil, err
	}
	if m.Skills == nil {
		return nil, ErrNotFound
	}
	return m.Skills, nil
}
