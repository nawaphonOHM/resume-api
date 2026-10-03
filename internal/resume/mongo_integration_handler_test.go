//go:build testcontainers

// Package resume defines domain data models, repository interfaces, and services for resume data.
package resume

import (
	"net/http"
	"testing"
)

func seedIntegrationData(h *testMongoHelper) {
	h.insert("details", 1, Details{Email: "h@test.com"})
	h.insert("education", 1, Education{Institution: "Uni"})
	h.insert("experience", 1, []*ExperienceItem{{Company: "Corp"}})
	h.insert("links", 1, []*LinkItem{{Label: "Web"}})
	h.insert("name", 1, "Dev Name")
	h.insert("title", 1, []string{"Engineer"})
}

func verifyAllOK(t *testing.T, h *Handler) {
	t.Helper()
	assertStatus(t, h.GetDetails(nil), http.StatusOK)
	assertStatus(t, h.GetEducations(nil), http.StatusOK)
	assertStatus(t, h.GetExperiences(nil), http.StatusOK)
	assertStatus(t, h.GetLinks(nil), http.StatusOK)
	assertStatus(t, h.GetNames(nil), http.StatusOK)
	assertStatus(t, h.GetTitles(nil), http.StatusOK)
}

func verifyAllNotFound(t *testing.T, h *Handler) {
	t.Helper()
	assertStatus(t, h.GetDetails(nil), http.StatusNotFound)
	assertStatus(t, h.GetEducations(nil), http.StatusNotFound)
	assertStatus(t, h.GetExperiences(nil), http.StatusNotFound)
	assertStatus(t, h.GetLinks(nil), http.StatusNotFound)
	assertStatus(t, h.GetNames(nil), http.StatusNotFound)
	assertStatus(t, h.GetTitles(nil), http.StatusNotFound)
}

func TestHandler_Integration_Flow(t *testing.T) {
	helper := newTestMongoHelper(t)
	svc := NewService(helper.repo)
	h := NewHandler(svc)

	verifyAllNotFound(t, h)
	seedIntegrationData(helper)
	verifyAllOK(t, h)
}
