// Package resume defines domain data models, repository interfaces, and services for resume data.
package resume

import (
	"context"

	"github.com/nawaphonOHM/whatever/v3/pkg/rest"
)

// GetSkills handles GET /v1/skills.
func (h *Handler) GetSkills(rest.Context) rest.Response {
	data, err := h.service.GetSkills(context.Background())
	if err != nil {
		return mapError(err)
	}
	return rest.OK(data)
}
