// Package resume defines domain data models, repository interfaces, and services for resume data.
package resume

import (
	"context"
	"errors"
	"testing"

	"github.com/nawaphonOHM/whatever/pkg/mongodb"
)

type mockCursor struct {
	err       error
	decodeErr error
	next      bool
	closed    bool
}

func (m *mockCursor) Next(context.Context) bool { return m.next }
func (m *mockCursor) Err() error                { return m.err }
func (m *mockCursor) Close(context.Context) error {
	m.closed = true
	return nil
}
func (m *mockCursor) Decode(v any) error {
	if m.decodeErr != nil {
		return m.decodeErr
	}
	if d, ok := v.(*Document[string]); ok {
		*d = Document[string]{Value: "test-name"}
	}
	return nil
}

func TestNewRepositories(t *testing.T) {
	if NewMongoRepository(nil) == nil {
		t.Fatal("expected non-nil MongoRepository")
	}
	if NewMemoryRepository() == nil {
		t.Fatal("expected non-nil MemoryRepository")
	}
}

func assertNilClient(t *testing.T, err error) {
	t.Helper()
	if !errors.Is(err, mongodb.ErrNilClient) {
		t.Fatalf("expected ErrNilClient, got: %v", err)
	}
}

func TestMongoRepository_NilClient(t *testing.T) {
	ctx := context.Background()
	r := NewMongoRepository(nil)
	_, err := r.GetLatestDetails(ctx)
	assertNilClient(t, err)
	_, err = r.GetLatestEducation(ctx)
	assertNilClient(t, err)
	_, err = r.GetLatestExperience(ctx)
	assertNilClient(t, err)
	_, err = r.GetLatestLinks(ctx)
	assertNilClient(t, err)
}

func TestMongoRepository_NilClientMore(t *testing.T) {
	ctx := context.Background()
	r := NewMongoRepository(nil)
	_, err := r.GetLatestName(ctx)
	assertNilClient(t, err)
	_, err = r.GetLatestTitles(ctx)
	assertNilClient(t, err)
}
