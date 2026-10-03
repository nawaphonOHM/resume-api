//go:build realmongo

// Package resume defines domain data models, repository interfaces, and services for resume data.
package resume

import (
	"net/http"
	"testing"

	"github.com/nawaphonOHM/whatever/pkg/rest"
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
	paths := []string{"/details", "/educations", "/experiences", "/links", "/names", "/titles"}
	if len(apis) != len(paths) {
		t.Fatalf("expected %d apis, got %d", len(paths), len(apis))
	}
	for i, path := range paths {
		verifyRouteAPI(t, apis[i], path)
	}
}

func verifyEndpointOK(t *testing.T, api *rest.ExportableAPI) {
	t.Helper()
	status := api.Handler(nil).StatusCode()
	if status != http.StatusOK {
		t.Fatalf("endpoint %s returned unexpected status: %d", api.Path, status)
	}
}

func verifyRegisteredEndpoints(t *testing.T, apis []*rest.ExportableAPI) {
	t.Helper()
	for _, api := range apis {
		verifyEndpointOK(t, api)
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
