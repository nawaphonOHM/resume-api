// Package main provides unit tests for application initialization and blueprint configuration.
package main

import (
	"testing"

	"github.com/nawaphonOHM/whatever/v2/pkg/rest"
)

const expectedTotalRoutes = 6

func assertCorsOrigins(t *testing.T, origins []string) {
	t.Helper()
	if len(origins) != 1 {
		t.Fatalf("expected 1 origin, got %d", len(origins))
	}
	if origins[0] != allowedOrigin {
		t.Fatalf("unexpected origin: %s", origins[0])
	}
}

func assertCorsMethods(t *testing.T, methods []rest.HTTPMethod) {
	t.Helper()
	if len(methods) != 1 {
		t.Fatalf("expected 1 method, got %d", len(methods))
	}
	if methods[0] != rest.GET {
		t.Fatalf("unexpected method: %s", methods[0])
	}
}

func TestBuildCorsSetting(t *testing.T) {
	cors := buildCorsSetting()
	if cors == nil {
		t.Fatal("expected non-nil cors setting")
	}
	assertCorsOrigins(t, cors.AllowOrigin())
	assertCorsMethods(t, cors.AllowHTTPMethods())
}

func assertBlueprintApis(t *testing.T, bp *rest.BluePrint, expected *rest.RRestAPIRegistration) {
	t.Helper()
	apis := bp.Apis()
	if len(apis) != 1 {
		t.Fatalf("expected 1 api registration, got %d", len(apis))
	}
	if apis[0] != expected {
		t.Fatalf("unexpected api registration: %v", apis[0])
	}
}

func TestBuildBlueprint(t *testing.T) {
	routes := &rest.RRestAPIRegistration{}
	bp := buildBlueprint(routes)
	if bp == nil {
		t.Fatal("expected non-nil blueprint")
	}
	assertBlueprintApis(t, bp, routes)
}

func TestInitRoutes(t *testing.T) {
	routes := initRoutes(nil)
	if routes == nil {
		t.Fatal("expected non-nil routes")
	}
	if len(routes.Apis) != expectedTotalRoutes {
		t.Fatalf("expected %d routes, got %d", expectedTotalRoutes, len(routes.Apis))
	}
}
