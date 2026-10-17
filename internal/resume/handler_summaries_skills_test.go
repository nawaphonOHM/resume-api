// Package resume defines domain data models, repository interfaces, and services for resume data.
package resume

import (
	"errors"
	"net/http"
	"testing"

	"github.com/nawaphonOHM/whatever/v3/pkg/rest"
)

func TestHandler_GetSummaries(t *testing.T) {
	expected := []string{"Experienced backend engineer", "Go specialist"}
	repo := &MemoryRepository{Summaries: expected}
	h := NewHandler(NewService(repo))
	assertJSONData(t, h.GetSummaries(nil), expected)
	repo.Summaries = nil
	assertStatus(t, h.GetSummaries(nil), http.StatusNotFound)
}

func TestHandler_GetSummaries_InternalError(t *testing.T) {
	repo := &MemoryRepository{Err: errors.New("summaries db error")}
	h := NewHandler(NewService(repo))
	assertStatus(t, h.GetSummaries(nil), http.StatusInternalServerError)
}

func TestHandler_GetSkills(t *testing.T) {
	expected := []string{"Go", "MongoDB", "Distributed Systems"}
	repo := &MemoryRepository{Skills: expected}
	h := NewHandler(NewService(repo))
	assertJSONData(t, h.GetSkills(nil), expected)
	repo.Skills = nil
	assertStatus(t, h.GetSkills(nil), http.StatusNotFound)
}

func TestHandler_GetSkills_InternalError(t *testing.T) {
	repo := &MemoryRepository{Err: errors.New("skills db error")}
	h := NewHandler(NewService(repo))
	assertStatus(t, h.GetSkills(nil), http.StatusInternalServerError)
}

func findRouteAPI(apis []*rest.ExportableAPI, path string) *rest.ExportableAPI {
	for _, api := range apis {
		if string(api.Path) == path {
			return api
		}
	}
	return nil
}

func verifyRoutePayload(t *testing.T, apis []*rest.ExportableAPI, path string, expected []string) {
	t.Helper()
	api := findRouteAPI(apis, path)
	if api == nil {
		t.Fatalf("missing %s route", path)
	}
	assertJSONData(t, api.Handler(nil), expected)
}

func TestHandler_Routes_SummariesAndSkills(t *testing.T) {
	sums := []string{"Experienced backend engineer"}
	skills := []string{"Go", "MongoDB"}
	repo := &MemoryRepository{Summaries: sums, Skills: skills}
	reg := Routes(NewHandler(NewService(repo)))
	verifyRoutePayload(t, reg.Apis, "/summaries", sums)
	verifyRoutePayload(t, reg.Apis, "/skills", skills)
}
