// Package resume defines domain data models, repository interfaces, and services for resume data.
package resume

import (
	"context"
	"errors"
	"testing"
	"time"
)

func assertExpectedError(t *testing.T, err, target error) {
	t.Helper()
	if !errors.Is(err, target) {
		t.Fatalf("expected error %v, got %v", target, err)
	}
}

func assertRepoError(t *testing.T, ctx context.Context, repo *MemoryRepository, expected error) {
	t.Helper()
	_, err1 := repo.GetLatestDetails(ctx)
	assertExpectedError(t, err1, expected)
	_, err2 := repo.GetLatestEducations(ctx)
	assertExpectedError(t, err2, expected)
	_, err3 := repo.GetLatestExperiences(ctx)
	assertExpectedError(t, err3, expected)
}

func assertRepoErrorMore(t *testing.T, ctx context.Context, repo *MemoryRepository, expected error) {
	t.Helper()
	_, err1 := repo.GetLatestLinks(ctx)
	assertExpectedError(t, err1, expected)
	_, err2 := repo.GetLatestNames(ctx)
	assertExpectedError(t, err2, expected)
	_, err3 := repo.GetLatestTitles(ctx)
	assertExpectedError(t, err3, expected)
}

func verifyAllRepoErrors(t *testing.T, ctx context.Context, repo *MemoryRepository, expected error) {
	t.Helper()
	assertRepoError(t, ctx, repo, expected)
	assertRepoErrorMore(t, ctx, repo, expected)
}

func TestMemoryRepository_CanceledContext(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	verifyAllRepoErrors(t, ctx, NewMemoryRepository(), context.Canceled)
}

func TestMemoryRepository_DeadlineExceeded(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Nanosecond)
	defer cancel()
	time.Sleep(time.Millisecond)
	verifyAllRepoErrors(t, ctx, NewMemoryRepository(), context.DeadlineExceeded)
}

func TestMemoryRepository_CustomError(t *testing.T) {
	customErr := errors.New("custom memory repo error")
	repo := &MemoryRepository{Err: customErr}
	verifyAllRepoErrors(t, context.Background(), repo, customErr)
}

func testPositiveFirst(t *testing.T, ctx context.Context, repo *MemoryRepository) {
	t.Helper()
	_, err1 := repo.GetLatestDetails(ctx)
	assertNoError(t, err1)
	_, err2 := repo.GetLatestEducations(ctx)
	assertNoError(t, err2)
	_, err3 := repo.GetLatestExperiences(ctx)
	assertNoError(t, err3)
}

func testPositiveSecond(t *testing.T, ctx context.Context, repo *MemoryRepository) {
	t.Helper()
	_, err1 := repo.GetLatestLinks(ctx)
	assertNoError(t, err1)
	_, err2 := repo.GetLatestNames(ctx)
	assertNoError(t, err2)
	_, err3 := repo.GetLatestTitles(ctx)
	assertNoError(t, err3)
}

func TestMemoryRepository_Positive(t *testing.T) {
	repo := &MemoryRepository{
		Details:    &Details{Email: "me@example.com"},
		Education:  &Education{Degree: "B.Sc."},
		Experience: []*ExperienceItem{{Company: "Acme"}},
		Links:      []*LinkItem{{Label: "Web"}},
		Name:       "Jane Doe",
		Titles:     []string{"Lead Engineer"},
	}
	ctx := context.Background()
	testPositiveFirst(t, ctx, repo)
	testPositiveSecond(t, ctx, repo)
}
