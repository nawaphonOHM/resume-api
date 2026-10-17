// Package resume defines domain data models, repository interfaces, and services for resume data.
package resume

import (
	"context"
	"net/http"
	"testing"
)

type emptyTitleRepository struct {
	Repository
}

func (emptyTitleRepository) GetLatestTitles(context.Context) (string, error) {
	return "", nil
}

func TestService_GetLinks(t *testing.T) {
	repo := &MemoryRepository{Links: []*LinkItem{{
		Label: "GitHub",
		Logo:  &Logo{Src: "github.png"},
	}}}
	svc := NewService(repo)
	res, err := svc.GetLinks(context.Background())
	assertNoError(t, err)
	if len(res) == 0 || res[0].Logo == nil {
		t.Fatal("expected links items with logo pointer")
	}
	repo.Links = nil
	_, err = svc.GetLinks(context.Background())
	assertNotFound(t, err)
}

func TestService_GetName(t *testing.T) {
	repo := &MemoryRepository{Name: "John Doe"}
	svc := NewService(repo)
	res, err := svc.GetNames(context.Background())
	assertNoError(t, err)
	if res != "John Doe" {
		t.Fatalf("unexpected name: %s", res)
	}
	repo.Name = ""
	_, err = svc.GetNames(context.Background())
	assertNotFound(t, err)
}

func TestService_GetTitles(t *testing.T) {
	repo := &MemoryRepository{Titles: "Engineer"}
	svc := NewService(repo)
	res, err := svc.GetTitles(context.Background())
	assertNoError(t, err)
	if res != "Engineer" {
		t.Fatalf("unexpected title: %s", res)
	}
	repo.Titles = ""
	_, err = svc.GetTitles(context.Background())
	assertNotFound(t, err)
}

func TestService_GetTitles_EmptyRepositoryValue(t *testing.T) {
	svc := NewService(emptyTitleRepository{})
	_, err := svc.GetTitles(context.Background())
	assertNotFound(t, err)
	assertStatus(t, NewHandler(svc).GetTitles(nil), http.StatusNotFound)
}
