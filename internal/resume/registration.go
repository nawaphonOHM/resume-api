// Package resume defines domain data models, repository interfaces, and services for resume data.
package resume

import (
	"github.com/nawaphonOHM/whatever/pkg/rest"
)

// Routes returns the route registration for resume API version 1.
func Routes(h *Handler) *rest.RRestAPIRegistration {
	return &rest.RRestAPIRegistration{
		Prefix:  "",
		Version: 1,
		Apis: []*rest.ExportableAPI{
			{Path: "/details", Method: rest.GET, Handler: h.GetDetails},
			{Path: "/education", Method: rest.GET, Handler: h.GetEducation},
			{Path: "/experience", Method: rest.GET, Handler: h.GetExperience},
			{Path: "/links", Method: rest.GET, Handler: h.GetLinks},
			{Path: "/names", Method: rest.GET, Handler: h.GetNames},
			{Path: "/titles", Method: rest.GET, Handler: h.GetTitles},
		},
	}
}
