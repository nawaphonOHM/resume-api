//go:build realmongo

// Package resume defines domain data models, repository interfaces, and services for resume data.
package resume

import (
	"testing"

	"github.com/nawaphonOHM/whatever/v3/pkg/rest"
)

func verifyDetailsPayload(t *testing.T, apis []*rest.ExportableAPI) {
	t.Helper()
	d := decodePayload[*Details](t, apis, "/details")
	if d == nil {
		t.Fatal("details payload is nil")
	}
	checkNonEmptyString(t, "birthDate", d.BirthDate)
	checkNonEmptyString(t, "email", d.Email)
	checkNonEmptyString(t, "location", d.Location)
	checkNonEmptyString(t, "nationality", d.Nationality)
	checkNonEmptyString(t, "phoneLabel", d.PhoneLabel)
}

func verifyEducationLogo(t *testing.T, logo *Logo) {
	t.Helper()
	if logo != nil {
		checkNonEmptyString(t, "institutionLogo.src", logo.Src)
	}
}

func verifySeniorProject(t *testing.T, sp *SeniorProject) {
	t.Helper()
	if sp != nil {
		checkNonEmptyString(t, "seniorProject.name", sp.Name)
		checkNonEmptyString(t, "seniorProject.url", sp.URL)
	}
}

func verifyEducationPayload(t *testing.T, apis []*rest.ExportableAPI) {
	t.Helper()
	e := decodePayload[*Education](t, apis, "/educations")
	if e == nil {
		t.Fatal("education payload is nil")
	}
	checkNonEmptyString(t, "institution", e.Institution)
	checkNonEmptyString(t, "degree", e.Degree)
	checkNonEmptyString(t, "period", e.Period)
	checkNonEmptyString(t, "gpax", e.GPAX)
	verifyEducationLogo(t, e.InstitutionLogo)
	verifySeniorProject(t, e.SeniorProject)
}

func verifyExperienceItem(t *testing.T, item *ExperienceItem) {
	t.Helper()
	if item == nil {
		t.Fatal("experience item is nil")
	}
	checkNonEmptyString(t, "company", item.Company)
	checkNonEmptyString(t, "role", item.Role)
	checkNonEmptyString(t, "period", item.Period)
	checkNonEmptyString(t, "location", item.Location)
}

func verifyExperiencesPayload(t *testing.T, apis []*rest.ExportableAPI) {
	t.Helper()
	items := decodePayload[[]*ExperienceItem](t, apis, "/experiences")
	if len(items) == 0 {
		t.Fatal("experiences payload is empty")
	}
	for _, item := range items {
		verifyExperienceItem(t, item)
	}
}

func verifyLinkItem(t *testing.T, item *LinkItem) {
	t.Helper()
	if item == nil {
		t.Fatal("link item is nil")
	}
	checkNonEmptyString(t, "label", item.Label)
	checkNonEmptyString(t, "url", item.URL)
	if item.Logo != nil {
		checkNonEmptyString(t, "logo.src", item.Logo.Src)
	}
}

func verifyLinksPayload(t *testing.T, apis []*rest.ExportableAPI) {
	t.Helper()
	items := decodePayload[[]*LinkItem](t, apis, "/links")
	if len(items) == 0 {
		t.Fatal("links payload is empty")
	}
	for _, item := range items {
		verifyLinkItem(t, item)
	}
}
