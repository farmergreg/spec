package credit

import (
	"testing"
)

func TestInit(t *testing.T) {
	// Test that EnumCreditListAll is populated
	if len(EnumCreditListAll) == 0 {
		t.Error("EnumCreditListAll should not be empty")
	}

	// Test that EnumCreditList is populated and filtered correctly
	if len(EnumCreditList) == 0 {
		t.Error("EnumCreditList should not be empty")
	}
	if len(EnumCreditList) > len(EnumCreditListAll) {
		t.Error("EnumCreditList should not be larger than EnumCreditListAll")
	}

	// Test that EnumCreditMap is populated correctly
	if len(EnumCreditMap) != len(EnumCreditList) {
		t.Errorf("EnumCreditMap size (%d) should match EnumCreditList size (%d)",
			len(EnumCreditMap), len(EnumCreditList))
	}

	// Test that all items in EnumCreditList are properly filtered
	for _, item := range EnumCreditList {
		if !bool(item.IsReleased) {
			t.Error("EnumCreditList should only contain released items")
		}
		if bool(item.IsImportOnly) {
			t.Error("EnumCreditList should not contain import-only items")
		}
	}

	// Test that all items in EnumCreditMap are accessible by their ID
	for _, item := range EnumCreditList {
		if mapped, exists := EnumCreditMap[item.ID]; !exists {
			t.Errorf("Item with ID %s should exist in EnumCreditMap", item.ID)
		} else if mapped != item {
			t.Errorf("Mapped item for ID %s should be the same instance", item.ID)
		}
	}
}
