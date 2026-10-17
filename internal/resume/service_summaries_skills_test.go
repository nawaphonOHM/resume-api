// Package resume defines domain data models, repository interfaces, and services for resume data.
package resume

import (
	"context"
	"errors"
	"testing"
)

func TestService_GetSummaries(t *testing.T) {
	repo := &MemoryRepository{Summaries: []string{"Experienced backend engineer"}}
	svc := NewService(repo)
	res, err := svc.GetSummaries(context.Background())
	assertNoError(t, err)
	if len(res) != 1 || res[0] != "Experienced backend engineer" {
		t.Fatalf("unexpected summaries: %v", res)
	}
}

func TestService_GetSummaries_NotFound(t *testing.T) {
	repo := &MemoryRepository{Summaries: make([]string, 0)}
	svc := NewService(repo)
	_, err := svc.GetSummaries(context.Background())
	assertNotFound(t, err)
	repo.Summaries = nil
	_, err = svc.GetSummaries(context.Background())
	assertNotFound(t, err)
}

func TestService_GetSummaries_Error(t *testing.T) {
	dummyErr := errors.New("query failed")
	repo := &MemoryRepository{Err: dummyErr}
	svc := NewService(repo)
	_, err := svc.GetSummaries(context.Background())
	if !errors.Is(err, dummyErr) {
		t.Fatalf("expected dummyErr, got %v", err)
	}
}

func TestService_GetSkills(t *testing.T) {
	repo := &MemoryRepository{Skills: []string{"Go", "MongoDB"}}
	svc := NewService(repo)
	res, err := svc.GetSkills(context.Background())
	assertNoError(t, err)
	if len(res) != 2 || res[0] != "Go" {
		t.Fatalf("unexpected skills: %v", res)
	}
}

func TestService_GetSkills_NotFound(t *testing.T) {
	repo := &MemoryRepository{Skills: make([]string, 0)}
	svc := NewService(repo)
	_, err := svc.GetSkills(context.Background())
	assertNotFound(t, err)
	repo.Skills = nil
	_, err = svc.GetSkills(context.Background())
	assertNotFound(t, err)
}

func TestService_GetSkills_Error(t *testing.T) {
	dummyErr := errors.New("query failed")
	repo := &MemoryRepository{Err: dummyErr}
	svc := NewService(repo)
	_, err := svc.GetSkills(context.Background())
	if !errors.Is(err, dummyErr) {
		t.Fatalf("expected dummyErr, got %v", err)
	}
}
