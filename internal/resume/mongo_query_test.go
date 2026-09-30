// Package resume defines domain data models, repository interfaces, and services for resume data.
package resume

import (
	"context"
	"errors"
	"testing"
)

func assertName(t *testing.T, val *string) {
	t.Helper()
	if val == nil {
		t.Fatal("expected non-nil val")
	}
	if *val != "test-name" {
		t.Fatalf("unexpected val: %s", *val)
	}
}

func TestDecodeCursor_Success(t *testing.T) {
	ctx := context.Background()
	cur := &mockCursor{next: true}
	val, err := decodeCursor[string](ctx, cur)
	assertNoError(t, err)
	assertName(t, val)
	if !cur.closed {
		t.Fatal("expected cursor to be closed")
	}
}

func TestDecodeCursor_NotFound(t *testing.T) {
	ctx := context.Background()
	cur := &mockCursor{next: false}
	_, err := decodeCursor[string](ctx, cur)
	assertNotFound(t, err)
}

func TestDecodeCursor_CursorErr(t *testing.T) {
	ctx := context.Background()
	dummyErr := errors.New("cursor failure")
	cur := &mockCursor{next: false, err: dummyErr}
	_, err := decodeCursor[string](ctx, cur)
	if !errors.Is(err, dummyErr) {
		t.Fatalf("expected dummyErr, got: %v", err)
	}
}

func TestDecodeCursor_DecodeErr(t *testing.T) {
	ctx := context.Background()
	dummyErr := errors.New("decode failure")
	cur := &mockCursor{next: true, decodeErr: dummyErr}
	_, err := decodeCursor[string](ctx, cur)
	if !errors.Is(err, dummyErr) {
		t.Fatalf("expected decodeErr, got: %v", err)
	}
}
