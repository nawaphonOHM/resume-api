// Package resume defines domain data models, repository interfaces, and services for resume data.
package resume

import (
	"net/http"
	"testing"

	"github.com/nawaphonOHM/whatever/pkg/rest"
)

const expectedTotalRoutes = 6

func TestHandler_GetDetails(t *testing.T) {
	repo := &MemoryRepository{Details: &Details{Email: "test@example.com"}}
	h := NewHandler(NewService(repo))
	assertStatus(t, h.GetDetails(nil), http.StatusOK)
	repo.Details = nil
	assertStatus(t, h.GetDetails(nil), http.StatusNotFound)
}

func TestHandler_GetEducation(t *testing.T) {
	repo := &MemoryRepository{Education: &Education{Degree: "B.Sc."}}
	h := NewHandler(NewService(repo))
	assertStatus(t, h.GetEducation(nil), http.StatusOK)
	repo.Education = nil
	assertStatus(t, h.GetEducation(nil), http.StatusNotFound)
}

func TestHandler_GetExperience(t *testing.T) {
	repo := &MemoryRepository{Experience: []*ExperienceItem{{Company: "Corp"}}}
	h := NewHandler(NewService(repo))
	assertStatus(t, h.GetExperience(nil), http.StatusOK)
	repo.Experience = nil
	assertStatus(t, h.GetExperience(nil), http.StatusNotFound)
}

func verifyRouteMethod(t *testing.T, api *rest.ExportableAPI) {
	t.Helper()
	if api.Method != rest.GET {
		t.Fatalf("expected GET method, got %s", api.Method)
	}
}

func verifyAllRouteMethods(t *testing.T, apis []*rest.ExportableAPI) {
	t.Helper()
	for _, api := range apis {
		verifyRouteMethod(t, api)
	}
}

func TestHandler_Routes(t *testing.T) {
	repo := &MemoryRepository{}
	h := NewHandler(NewService(repo))
	reg := Routes(h)
	if reg == nil || len(reg.Apis) != expectedTotalRoutes {
		t.Fatalf("unexpected routes registration: %v", reg)
	}
	verifyAllRouteMethods(t, reg.Apis)
}
