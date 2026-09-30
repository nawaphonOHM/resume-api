// Package resume defines domain data models, repository interfaces, and services for resume data.
package resume

import (
	"errors"
	"net/http"
	"testing"
)

func TestHandler_GetLinks(t *testing.T) {
	repo := &MemoryRepository{Links: []*LinkItem{{Label: "LinkedIn"}}}
	h := NewHandler(NewService(repo))
	assertStatus(t, h.GetLinks(nil), http.StatusOK)
	repo.Links = nil
	assertStatus(t, h.GetLinks(nil), http.StatusNotFound)
}

func TestHandler_GetNames(t *testing.T) {
	repo := &MemoryRepository{Name: "John Doe"}
	h := NewHandler(NewService(repo))
	assertStatus(t, h.GetNames(nil), http.StatusOK)
	repo.Name = ""
	assertStatus(t, h.GetNames(nil), http.StatusNotFound)
}

func TestHandler_GetTitles(t *testing.T) {
	repo := &MemoryRepository{Titles: []string{"Architect"}}
	h := NewHandler(NewService(repo))
	assertStatus(t, h.GetTitles(nil), http.StatusOK)
	repo.Titles = nil
	assertStatus(t, h.GetTitles(nil), http.StatusNotFound)
}

func TestHandler_InternalError(t *testing.T) {
	repo := &MemoryRepository{Err: errors.New("db failure")}
	h := NewHandler(NewService(repo))
	assertStatus(t, h.GetDetails(nil), http.StatusInternalServerError)
}
