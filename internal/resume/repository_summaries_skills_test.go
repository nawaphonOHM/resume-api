// Package resume defines domain data models, repository interfaces, and services for resume data.
package resume

import (
	"context"
	"errors"
	"testing"
	"time"
)

func TestMongoRepository_NilClient_SummariesAndSkills(t *testing.T) {
	ctx := context.Background()
	r := NewMongoRepository(nil)
	_, err := r.GetLatestSummaries(ctx)
	assertNilClient(t, err)
	_, err = r.GetLatestSkills(ctx)
	assertNilClient(t, err)
}

func TestMemoryRepository_Summaries_Positive(t *testing.T) {
	expected := []string{"Summary point 1", "Summary point 2"}
	repo := &MemoryRepository{Summaries: expected}
	res, err := repo.GetLatestSummaries(context.Background())
	assertNoError(t, err)
	if len(res) != len(expected) || res[0] != expected[0] {
		t.Fatalf("unexpected summaries: %v", res)
	}
}

func TestMemoryRepository_Skills_Positive(t *testing.T) {
	expected := []string{"Go", "MongoDB"}
	repo := &MemoryRepository{Skills: expected}
	res, err := repo.GetLatestSkills(context.Background())
	assertNoError(t, err)
	if len(res) != len(expected) || res[0] != expected[0] {
		t.Fatalf("unexpected skills: %v", res)
	}
}

func TestMemoryRepository_SummariesSkills_NotFound(t *testing.T) {
	repo := NewMemoryRepository()
	ctx := context.Background()
	_, err := repo.GetLatestSummaries(ctx)
	assertNotFound(t, err)
	_, err = repo.GetLatestSkills(ctx)
	assertNotFound(t, err)
}

func TestMemoryRepository_SummariesSkills_Errors(t *testing.T) {
	customErr := errors.New("custom error")
	repo := &MemoryRepository{Err: customErr}
	ctx := context.Background()
	_, err := repo.GetLatestSummaries(ctx)
	if !errors.Is(err, customErr) {
		t.Fatalf("expected customErr, got %v", err)
	}
	_, err = repo.GetLatestSkills(ctx)
	if !errors.Is(err, customErr) {
		t.Fatalf("expected customErr, got %v", err)
	}
}

func assertSummariesSkillsContext(t *testing.T, ctx context.Context, target error) {
	t.Helper()
	repo := NewMemoryRepository()
	_, err := repo.GetLatestSummaries(ctx)
	assertExpectedError(t, err, target)
	_, err = repo.GetLatestSkills(ctx)
	assertExpectedError(t, err, target)
}

func TestMemoryRepository_SummariesSkills_CanceledContext(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	assertSummariesSkillsContext(t, ctx, context.Canceled)
}

func TestMemoryRepository_SummariesSkills_DeadlineExceeded(t *testing.T) {
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(-time.Second))
	defer cancel()
	assertSummariesSkillsContext(t, ctx, context.DeadlineExceeded)
}
