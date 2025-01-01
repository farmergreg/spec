package qslmedium

import (
	"testing"
)

func TestInit(t *testing.T) {
	// Test that the maps and slices are initialized
	if EnumQSLMediumMap == nil {
		t.Error("EnumQSLMediumMap was not initialized")
	}
	if EnumQSLMediumList == nil {
		t.Error("EnumQSLMediumList was not initialized")
	}
	if EnumQSLMediumListAll == nil {
		t.Error("EnumQSLMediumListAll was not initialized")
	}

	// Test that EnumQSLMediumListAll contains items
	if len(EnumQSLMediumListAll) == 0 {
		t.Error("EnumQSLMediumListAll is empty")
	}

	// Test that filtered list (EnumQSLMediumList) only contains released and non-import-only items
	for _, item := range EnumQSLMediumList {
		if !bool(item.IsReleased) {
			t.Errorf("EnumQSLMediumList contains unreleased item: %s", item.ID)
		}
		if bool(item.IsImportOnly) {
			t.Errorf("EnumQSLMediumList contains import-only item: %s", item.ID)
		}
	}

	// Test that map contains all items from the filtered list
	if len(EnumQSLMediumMap) != len(EnumQSLMediumList) {
		t.Errorf("Map size (%d) does not match list size (%d)",
			len(EnumQSLMediumMap), len(EnumQSLMediumList))
	}

	// Verify each item in the list exists in the map
	for _, item := range EnumQSLMediumList {
		if mapItem, exists := EnumQSLMediumMap[item.ID]; !exists {
			t.Errorf("Item %s exists in list but not in map", item.ID)
		} else if mapItem != item {
			t.Errorf("Map item pointer does not match list item pointer for %s", item.ID)
		}
	}
}
