//go:build testcontainers

// Package resume defines domain data models, repository interfaces, and services for resume data.
package resume

import (
	"context"
	"testing"
)

func TestMongoRepository_Integration_Names(t *testing.T) {
	h := newTestMongoHelper(t)
	ctx := context.Background()
	_, err := h.repo.GetLatestNames(ctx)
	assertNotFound(t, err)

	h.insert("name", 1, "Old Name")
	h.insert("name", 2, "Latest Name")

	res, err := h.repo.GetLatestNames(ctx)
	assertNoError(t, err)
	if res != "Latest Name" {
		t.Fatalf("expected 'Latest Name', got: %s", res)
	}
}

func verifyTitleValues(t *testing.T, res []string) {
	t.Helper()
	if res[0] != "Senior Engineer" {
		t.Fatalf("unexpected title 0: %s", res[0])
	}
	if res[1] != "Architect" {
		t.Fatalf("unexpected title 1: %s", res[1])
	}
}

func seedTitles(h *testMongoHelper) {
	h.insert("title", 1, []string{"Junior Dev"})
	h.insert("title", 2, []string{"Senior Engineer", "Architect"})
}

func verifyTitles(t *testing.T, res []string) {
	t.Helper()
	if len(res) != 2 {
		t.Fatalf("unexpected title count: %d", len(res))
	}
	verifyTitleValues(t, res)
}

func TestMongoRepository_Integration_Titles(t *testing.T) {
	h := newTestMongoHelper(t)
	ctx := context.Background()
	_, err := h.repo.GetLatestTitles(ctx)
	assertNotFound(t, err)

	seedTitles(h)
	res, err := h.repo.GetLatestTitles(ctx)
	assertNoError(t, err)
	verifyTitles(t, res)
}
