//go:build realmongo

// Package resume defines domain data models, repository interfaces, and services for resume data.
package resume

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/nawaphonOHM/whatever/v3/pkg/rest"
)

func checkNonEmptyString(t *testing.T, field, value string) {
	t.Helper()
	if value == "" {
		t.Fatalf("expected non-empty %s", field)
	}
}

func getRoute(t *testing.T, apis []*rest.ExportableAPI, path string) *rest.ExportableAPI {
	t.Helper()
	api := findRouteAPI(apis, path)
	if api == nil {
		t.Fatalf("route %s not found", path)
	}
	return api
}

func unmarshalSuccess[T any](t *testing.T, resp rest.Response, path string) T {
	t.Helper()
	raw, err := json.Marshal(resp)
	assertNoError(t, err)
	var res rest.SuccessResponse[T]
	assertNoError(t, json.Unmarshal(raw, &res))
	if !res.Success {
		t.Fatalf("%s response success flag is false", path)
	}
	return res.Data
}

func decodePayload[T any](t *testing.T, apis []*rest.ExportableAPI, path string) T {
	t.Helper()
	api := getRoute(t, apis, path)
	resp := api.Handler(nil)
	assertStatus(t, resp, http.StatusOK)
	return unmarshalSuccess[T](t, resp, path)
}

func verifyStringPayload(t *testing.T, apis []*rest.ExportableAPI, path, expected string) {
	t.Helper()
	val := decodePayload[string](t, apis, path)
	if val != expected {
		t.Fatalf("%s returned %q, expected %q", path, val, expected)
	}
}

func verifySlicePayload(t *testing.T, apis []*rest.ExportableAPI, path string) {
	t.Helper()
	slice := decodePayload[[]string](t, apis, path)
	if len(slice) == 0 {
		t.Fatalf("%s returned empty slice", path)
	}
	for _, item := range slice {
		checkNonEmptyString(t, path+" item", item)
	}
}

func verifyPayloadIntegrity(t *testing.T, apis []*rest.ExportableAPI) {
	t.Helper()
	verifyDetailsPayload(t, apis)
	verifyEducationPayload(t, apis)
	verifyExperiencesPayload(t, apis)
	verifyLinksPayload(t, apis)
	verifyStringPayload(t, apis, "/names", "Nawaphon Isarathanachaikul")
	verifyStringPayload(t, apis, "/titles", "Backend Software Engineer")
	verifySlicePayload(t, apis, "/skills")
	verifySlicePayload(t, apis, "/summaries")
}
