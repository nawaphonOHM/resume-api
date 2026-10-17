// Package resume defines domain data models, repository interfaces, and services for resume data.
package resume

import (
	"context"
)

// GetLatestSummaries retrieves the latest summaries document from MongoDB.
func (r *MongoRepository) GetLatestSummaries(ctx context.Context) ([]string, error) {
	res, err := queryLatest[[]string](ctx, r.client, collectionSummary)
	if err != nil {
		return nil, err
	}
	return *res, nil
}

// GetLatestSummaries returns the in-memory summaries or ErrNotFound.
func (m *MemoryRepository) GetLatestSummaries(ctx context.Context) ([]string, error) {
	if err := m.checkErr(ctx); err != nil {
		return nil, err
	}
	if m.Summaries == nil {
		return nil, ErrNotFound
	}
	return m.Summaries, nil
}
