// Package resume defines domain data models, repository interfaces, and services for resume data.
package resume

import (
	"context"

	"github.com/nawaphonOHM/whatever/v3/pkg/rest"
)

// GetSummaries handles GET /v1/summaries.
func (h *Handler) GetSummaries(rest.Context) rest.Response {
	data, err := h.service.GetSummaries(context.Background())
	if err != nil {
		return mapError(err)
	}
	return rest.OK(data)
}
