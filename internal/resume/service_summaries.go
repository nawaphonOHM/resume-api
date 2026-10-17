// Package resume defines domain data models, repository interfaces, and services for resume data.
package resume

import (
	"context"
)

// GetSummaries retrieves the latest resume summaries.
func (s *resumeService) GetSummaries(ctx context.Context) ([]string, error) {
	res, err := s.repo.GetLatestSummaries(ctx)
	if err != nil {
		return nil, err
	}
	if len(res) == 0 {
		return nil, ErrNotFound
	}
	return res, nil
}
