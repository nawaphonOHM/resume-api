// Package resume defines domain data models, repository interfaces, and services for resume data.
package resume

import (
	"context"
)

// GetSkills retrieves the latest skills list.
func (s *resumeService) GetSkills(ctx context.Context) ([]string, error) {
	res, err := s.repo.GetLatestSkills(ctx)
	if err != nil {
		return nil, err
	}
	if len(res) == 0 {
		return nil, ErrNotFound
	}
	return res, nil
}
