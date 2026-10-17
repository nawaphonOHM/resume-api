package resume

import (
	"testing"
)

func TestCollectionConstants(t *testing.T) {
	if collectionSummary != "summary" {
		t.Errorf("expected collectionSummary to be 'summary', got '%s'", collectionSummary)
	}
	if collectionSkill != "skill" {
		t.Errorf("expected collectionSkill to be 'skill', got '%s'", collectionSkill)
	}
}
