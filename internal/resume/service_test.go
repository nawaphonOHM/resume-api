// Package resume defines domain data models, repository interfaces, and services for resume data.
package resume

import (
	"context"
	"testing"
)

func assertEducationProject(t *testing.T, res *Education) {
	t.Helper()
	if res.SeniorProject == nil {
		t.Fatal("expected non-nil project")
	}
	if res.InstitutionLogo == nil {
		t.Fatal("expected non-nil logo")
	}
}

func assertEducation(t *testing.T, res *Education) {
	t.Helper()
	if res.Degree != "B.Sc." {
		t.Fatalf("unexpected degree: %s", res.Degree)
	}
	assertEducationProject(t, res)
}

func assertExperienceItem(t *testing.T, item *ExperienceItem) {
	t.Helper()
	if item.CompanyLogo == nil {
		t.Fatal("expected non-nil company logo")
	}
	if item.Client == nil {
		t.Fatal("expected non-nil client")
	}
}

func assertExperience(t *testing.T, res []*ExperienceItem) {
	t.Helper()
	if len(res) == 0 {
		t.Fatal("expected experience items")
	}
	assertExperienceItem(t, res[0])
}

func TestService_GetDetails(t *testing.T) {
	repo := &MemoryRepository{Details: &Details{Email: "test@example.com"}}
	svc := NewService(repo)
	res, err := svc.GetDetails(context.Background())
	assertNoError(t, err)
	if res.Email != "test@example.com" {
		t.Fatalf("unexpected email: %s", res.Email)
	}
	repo.Details = nil
	_, err = svc.GetDetails(context.Background())
	assertNotFound(t, err)
}

func TestService_GetEducations(t *testing.T) {
	repo := &MemoryRepository{Education: &Education{
		SeniorProject:   &SeniorProject{Name: "Proj"},
		InstitutionLogo: &Logo{Src: "logo.png"},
		Degree:          "B.Sc.",
	}}
	svc := NewService(repo)
	res, err := svc.GetEducations(context.Background())
	assertNoError(t, err)
	assertEducation(t, res)
	repo.Education = nil
	_, err = svc.GetEducations(context.Background())
	assertNotFound(t, err)
}

func TestService_GetExperiences(t *testing.T) {
	repo := &MemoryRepository{Experience: []*ExperienceItem{{
		CompanyLogo: &Logo{Src: "corp.png"},
		Client:      &ClientItem{Logo: &Logo{Src: "client.png"}, Name: "Client"},
		Company:     "Corp",
	}}}
	svc := NewService(repo)
	res, err := svc.GetExperiences(context.Background())
	assertNoError(t, err)
	assertExperience(t, res)
	repo.Experience = nil
	_, err = svc.GetExperiences(context.Background())
	assertNotFound(t, err)
}
