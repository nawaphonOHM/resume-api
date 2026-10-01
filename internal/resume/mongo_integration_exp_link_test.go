// Package resume defines domain data models, repository interfaces, and services for resume data.
package resume

import (
	"context"
	"testing"
)

func verifyExpCompanyRole(t *testing.T, item *ExperienceItem) {
	t.Helper()
	if item.Company != "Future Inc." {
		t.Fatalf("unexpected company: %s", item.Company)
	}
	if item.Role != "Tech Lead" {
		t.Fatalf("unexpected role: %s", item.Role)
	}
}

func verifyExpClient(t *testing.T, client *ClientItem) {
	t.Helper()
	if client == nil {
		t.Fatal("expected non-nil client")
	}
	if client.Name != "Big Client" {
		t.Fatalf("unexpected client name: %s", client.Name)
	}
}

func seedExperience(h *testMongoHelper) {
	exp1 := []*ExperienceItem{{Company: "Old Corp", Role: "Dev"}}
	exp2 := []*ExperienceItem{{
		Company:      "Future Inc.",
		Role:         "Tech Lead",
		Client:       &ClientItem{Name: "Big Client", Logo: &Logo{Src: "/c.png"}},
		Technologies: []string{"Go", "MongoDB"},
	}}
	h.insert("experience", 1, exp1)
	h.insert("experience", 2, exp2)
}

func verifyExpResult(t *testing.T, res []*ExperienceItem) {
	t.Helper()
	if len(res) == 0 {
		t.Fatal("expected non-empty experience slice")
	}
	verifyExpCompanyRole(t, res[0])
	verifyExpClient(t, res[0].Client)
}

func TestMongoRepository_Integration_Experience(t *testing.T) {
	h := newTestMongoHelper(t)
	ctx := context.Background()
	_, err := h.repo.GetLatestExperiences(ctx)
	assertNotFound(t, err)

	seedExperience(h)
	res, err := h.repo.GetLatestExperiences(ctx)
	assertNoError(t, err)
	verifyExpResult(t, res)
}

func verifyLinkItem(t *testing.T, item *LinkItem) {
	t.Helper()
	if item.Label != "GitHub" {
		t.Fatalf("unexpected label: %s", item.Label)
	}
	if item.URL != "https://github.com/test" {
		t.Fatalf("unexpected URL: %s", item.URL)
	}
}

func verifyLinkLogo(t *testing.T, logo *Logo) {
	t.Helper()
	if logo == nil {
		t.Fatal("expected non-nil logo")
	}
	if logo.Src != "/gh.png" {
		t.Fatalf("unexpected logo src: %s", logo.Src)
	}
}

func verifyLinksResult(t *testing.T, res []*LinkItem) {
	t.Helper()
	if len(res) == 0 {
		t.Fatal("expected non-empty links slice")
	}
	verifyLinkItem(t, res[0])
	verifyLinkLogo(t, res[0].Logo)
}

func seedLinks(h *testMongoHelper) {
	l1 := []*LinkItem{{Label: "OldLink", URL: "https://old.com"}}
	l2 := []*LinkItem{
		{Label: "GitHub", URL: "https://github.com/test", Logo: &Logo{Src: "/gh.png"}},
	}
	h.insert("links", 1, l1)
	h.insert("links", 2, l2)
}

func TestMongoRepository_Integration_Links(t *testing.T) {
	h := newTestMongoHelper(t)
	ctx := context.Background()
	_, err := h.repo.GetLatestLinks(ctx)
	assertNotFound(t, err)

	seedLinks(h)
	res, err := h.repo.GetLatestLinks(ctx)
	assertNoError(t, err)
	verifyLinksResult(t, res)
}
