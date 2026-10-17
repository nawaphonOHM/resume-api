//go:build realmongo

// Package resume defines domain data models, repository interfaces, and services for resume data.
package resume

import (
	"net/http"
	"testing"

	"github.com/nawaphonOHM/whatever/v3/pkg/rest"
)

func verifyRouteAPI(t *testing.T, api *rest.ExportableAPI, path string) {
	t.Helper()
	if string(api.Path) != path {
		t.Fatalf("unexpected path: %s", api.Path)
	}
	if api.Method != rest.GET {
		t.Fatalf("unexpected method: %s", api.Method)
	}
}

func verifyRoutesHeader(t *testing.T, reg *rest.RRestAPIRegistration) {
	t.Helper()
	if reg.Prefix != "" {
		t.Fatalf("unexpected prefix: %s", reg.Prefix)
	}
	if reg.Version != 1 {
		t.Fatalf("unexpected version: %d", reg.Version)
	}
}

func verifyRoutesAPIs(t *testing.T, apis []*rest.ExportableAPI) {
	t.Helper()
	paths := []string{"/details", "/educations", "/experiences", "/links", "/names", "/skills", "/summaries", "/titles"}
	if len(apis) != len(paths) {
		t.Fatalf("expected %d apis, got %d", len(paths), len(apis))
	}
	for i, path := range paths {
		verifyRouteAPI(t, apis[i], path)
	}
}

func isNewEndpoint(path rest.Pathz) bool {
	return path == "/skills" || path == "/summaries"
}

// verifyNewEndpointStatus ensures new endpoints return 200 (if collection is seeded)
// or 404 (if unseeded in the live DB), but never accept 500 internal errors.
func verifyNewEndpointStatus(t *testing.T, path rest.Pathz, status int) {
	t.Helper()
	if status != http.StatusOK && status != http.StatusNotFound {
		t.Fatalf("new endpoint %s returned unexpected status: %d (expected 200 or 404)", path, status)
	}
}

func verifyEndpointStatus(t *testing.T, api *rest.ExportableAPI) {
	t.Helper()
	status := api.Handler(nil).StatusCode()
	if isNewEndpoint(api.Path) {
		verifyNewEndpointStatus(t, api.Path, status)
		return
	}
	if status != http.StatusOK {
		t.Fatalf("endpoint %s returned unexpected status: %d", api.Path, status)
	}
}

func verifyRegisteredEndpoints(t *testing.T, apis []*rest.ExportableAPI) {
	t.Helper()
	for _, api := range apis {
		verifyEndpointStatus(t, api)
	}
}

func TestRoutes_Integration_Mongo(t *testing.T) {
	repo := connectURIRepository(t)
	svc := NewService(repo)
	reg := Routes(NewHandler(svc))

	verifyRoutesHeader(t, reg)
	verifyRoutesAPIs(t, reg.Apis)
	verifyRegisteredEndpoints(t, reg.Apis)
}
