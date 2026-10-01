// Package resume defines domain data models, repository interfaces, and services for resume data.
package resume

import (
	"errors"
	"testing"

	"github.com/nawaphonOHM/whatever/pkg/rest"
)

func assertNoError(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func assertNotFound(t *testing.T, err error) {
	t.Helper()
	if !errors.Is(err, ErrNotFound) {
		t.Fatalf("expected ErrNotFound, got: %v", err)
	}
}

func assertStatus(t *testing.T, resp rest.Response, expected int) {
	t.Helper()
	if resp.StatusCode() != expected {
		t.Fatalf("expected status %d, got %d", expected, resp.StatusCode())
	}
}
