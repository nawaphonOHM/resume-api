//go:build testcontainers

// Package resume defines domain data models, repository interfaces, and services for resume data.
package resume

import (
	"context"
	"testing"
)

func verifyDetailsBasic(t *testing.T, res *Details) {
	t.Helper()
	if res.BirthDate != "2000-01-01" {
		t.Fatalf("unexpected birth date: %s", res.BirthDate)
	}
	if res.Email != "v2@test.com" {
		t.Fatalf("unexpected email: %s", res.Email)
	}
}

func verifyDetailsLocation(t *testing.T, res *Details) {
	t.Helper()
	if res.Location != "Bangkok" {
		t.Fatalf("unexpected location: %s", res.Location)
	}
	if res.PhoneLabel != "+66-2-000-0000" {
		t.Fatalf("unexpected phone: %s", res.PhoneLabel)
	}
}

func seedDetails(h *testMongoHelper) {
	d1 := Details{BirthDate: "1990-01-01", Email: "v1@test.com"}
	d2 := Details{
		BirthDate:   "2000-01-01",
		Email:       "v2@test.com",
		Location:    "Bangkok",
		Nationality: "Thai",
		PhoneLabel:  "+66-2-000-0000",
	}
	h.insert("details", 1, d1)
	h.insert("details", 2, d2)
}

func TestMongoRepository_Integration_Details(t *testing.T) {
	h := newTestMongoHelper(t)
	ctx := context.Background()
	_, err := h.repo.GetLatestDetails(ctx)
	assertNotFound(t, err)

	seedDetails(h)
	res, err := h.repo.GetLatestDetails(ctx)
	assertNoError(t, err)
	verifyDetailsBasic(t, res)
	verifyDetailsLocation(t, res)
}

func verifyEduBasic(t *testing.T, res *Education) {
	t.Helper()
	if res.Institution != "Tech University" {
		t.Fatalf("unexpected institution: %s", res.Institution)
	}
	if res.Degree != "M.Sc." {
		t.Fatalf("unexpected degree: %s", res.Degree)
	}
}

func verifyEduSeniorProject(t *testing.T, proj *SeniorProject) {
	t.Helper()
	if proj == nil {
		t.Fatal("expected non-nil senior project")
	}
	if proj.Name != "AI Project" {
		t.Fatalf("unexpected project name: %s", proj.Name)
	}
}

func verifyEduLogo(t *testing.T, logo *Logo) {
	t.Helper()
	if logo == nil {
		t.Fatal("expected non-nil logo")
	}
	if logo.Src != "/logo2.png" {
		t.Fatalf("unexpected logo src: %s", logo.Src)
	}
}

func seedEducation(h *testMongoHelper) {
	e1 := Education{Institution: "Old College", Degree: "B.Sc."}
	e2 := Education{
		Institution:     "Tech University",
		Degree:          "M.Sc.",
		SeniorProject:   &SeniorProject{Name: "AI Project", URL: "https://example.com/ai"},
		InstitutionLogo: &Logo{Src: "/logo2.png"},
	}
	h.insert("education", 1, e1)
	h.insert("education", 2, e2)
}

func TestMongoRepository_Integration_Education(t *testing.T) {
	h := newTestMongoHelper(t)
	ctx := context.Background()
	_, err := h.repo.GetLatestEducations(ctx)
	assertNotFound(t, err)

	seedEducation(h)
	res, err := h.repo.GetLatestEducations(ctx)
	assertNoError(t, err)
	verifyEduBasic(t, res)
	verifyEduSeniorProject(t, res.SeniorProject)
	verifyEduLogo(t, res.InstitutionLogo)
}
