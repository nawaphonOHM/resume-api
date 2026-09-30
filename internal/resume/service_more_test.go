// Package resume defines domain data models, repository interfaces, and services for resume data.
package resume

import (
	"context"
	"testing"
)

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
	res, err := svc.GetName(context.Background())
	assertNoError(t, err)
	if res != "John Doe" {
		t.Fatalf("unexpected name: %s", res)
	}
	repo.Name = ""
	_, err = svc.GetName(context.Background())
	assertNotFound(t, err)
}

func TestService_GetTitles(t *testing.T) {
	repo := &MemoryRepository{Titles: []string{"Engineer"}}
	svc := NewService(repo)
	res, err := svc.GetTitles(context.Background())
	assertNoError(t, err)
	if len(res) == 0 {
		t.Fatal("expected title items")
	}
	repo.Titles = nil
	_, err = svc.GetTitles(context.Background())
	assertNotFound(t, err)
}
