// Package resume defines domain data models, repository interfaces, and services for resume data.
package resume

import (
	"errors"
	"net/http"
	"testing"
)

func TestHandler_InternalErrors_Educations(t *testing.T) {
	repo := &MemoryRepository{Err: errors.New("edu error")}
	h := NewHandler(NewService(repo))
	assertStatus(t, h.GetEducations(nil), http.StatusInternalServerError)
}

func TestHandler_InternalErrors_Experiences(t *testing.T) {
	repo := &MemoryRepository{Err: errors.New("exp error")}
	h := NewHandler(NewService(repo))
	assertStatus(t, h.GetExperiences(nil), http.StatusInternalServerError)
}

func TestHandler_InternalErrors_Links(t *testing.T) {
	repo := &MemoryRepository{Err: errors.New("link error")}
	h := NewHandler(NewService(repo))
	assertStatus(t, h.GetLinks(nil), http.StatusInternalServerError)
}

func TestHandler_InternalErrors_Names(t *testing.T) {
	repo := &MemoryRepository{Err: errors.New("name error")}
	h := NewHandler(NewService(repo))
	assertStatus(t, h.GetNames(nil), http.StatusInternalServerError)
}

func TestHandler_InternalErrors_Titles(t *testing.T) {
	repo := &MemoryRepository{Err: errors.New("titles error")}
	h := NewHandler(NewService(repo))
	assertStatus(t, h.GetTitles(nil), http.StatusInternalServerError)
}
