// Package resume defines domain data models, repository interfaces, and services for resume data.
package resume

import (
	"context"
	"errors"

	"github.com/nawaphonOHM/whatever/pkg/rest"
)

// Handler handles REST requests for resume data.
type Handler struct {
	service Service
}

// NewHandler creates a new Handler instance.
func NewHandler(service Service) *Handler {
	return &Handler{service: service}
}

func mapError(err error) rest.Response {
	if errors.Is(err, ErrNotFound) {
		return rest.NotFound("NOT_FOUND", "resource not found")
	}
	return rest.InternalServerError("INTERNAL_SERVER_ERROR", "internal server error")
}

// GetDetails handles GET /v1/details.
func (h *Handler) GetDetails(rest.Context) rest.Response {
	data, err := h.service.GetDetails(context.Background())
	if err != nil {
		return mapError(err)
	}
	return rest.OK(data)
}

// GetEducation handles GET /v1/education.
func (h *Handler) GetEducation(rest.Context) rest.Response {
	data, err := h.service.GetEducation(context.Background())
	if err != nil {
		return mapError(err)
	}
	return rest.OK(data)
}

// GetExperience handles GET /v1/experience.
func (h *Handler) GetExperience(rest.Context) rest.Response {
	data, err := h.service.GetExperience(context.Background())
	if err != nil {
		return mapError(err)
	}
	return rest.OK(data)
}

// GetLinks handles GET /v1/links.
func (h *Handler) GetLinks(rest.Context) rest.Response {
	data, err := h.service.GetLinks(context.Background())
	if err != nil {
		return mapError(err)
	}
	return rest.OK(data)
}

// GetNames handles GET /v1/names.
func (h *Handler) GetNames(rest.Context) rest.Response {
	data, err := h.service.GetName(context.Background())
	if err != nil {
		return mapError(err)
	}
	return rest.OK(data)
}

// GetTitles handles GET /v1/titles.
func (h *Handler) GetTitles(rest.Context) rest.Response {
	data, err := h.service.GetTitles(context.Background())
	if err != nil {
		return mapError(err)
	}
	return rest.OK(data)
}
