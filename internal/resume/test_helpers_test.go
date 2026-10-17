// Package resume defines domain data models, repository interfaces, and services for resume data.
package resume

import (
	"encoding/json"
	"errors"
	"net/http"
	"reflect"
	"testing"

	"github.com/nawaphonOHM/whatever/v3/pkg/rest"
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

func unmarshalJSONResponse(t *testing.T, resp rest.Response) *rest.SuccessResponse[[]string] {
	t.Helper()
	raw, err := json.Marshal(resp)
	if err != nil {
		t.Fatalf("failed to marshal response: %v", err)
	}
	var res rest.SuccessResponse[[]string]
	if err := json.Unmarshal(raw, &res); err != nil {
		t.Fatalf("failed to unmarshal response: %v", err)
	}
	return &res
}

func assertJSONData(t *testing.T, resp rest.Response, expected []string) {
	t.Helper()
	assertStatus(t, resp, http.StatusOK)
	res := unmarshalJSONResponse(t, resp)
	if !res.Success {
		t.Fatal("expected success to be true")
	}
	if !reflect.DeepEqual(res.Data, expected) {
		t.Fatalf("expected data %v, got %v", expected, res.Data)
	}
}
