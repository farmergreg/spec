package qsocomplete

import (
	"testing"
)

func TestInit(t *testing.T) {
	// Test that the maps and slices are properly initialized
	if len(EnumQSOCompleteListAll) == 0 {
		t.Error("EnumQSOCompleteListAll should not be empty")
	}

	if len(EnumQSOCompleteList) == 0 {
		t.Error("EnumQSOCompleteList should not be empty")
	}

	if len(EnumQSOCompleteMap) == 0 {
		t.Error("EnumQSOCompleteMap should not be empty")
	}

	// Test that EnumQSOCompleteList only contains released and non-import-only items
	for _, item := range EnumQSOCompleteList {
		if !bool(item.IsReleased) {
			t.Errorf("EnumQSOCompleteList contains unreleased item: %s", item.ID)
		}
		if bool(item.IsImportOnly) {
			t.Errorf("EnumQSOCompleteList contains import-only item: %s", item.ID)
		}
	}

	// Test that map entries correspond to list entries
	for _, item := range EnumQSOCompleteList {
		mapItem, exists := EnumQSOCompleteMap[item.ID]
		if !exists {
			t.Errorf("Item %s exists in list but not in map", item.ID)
		}
		if mapItem != item {
			t.Errorf("Map entry for %s doesn't match list entry", item.ID)
		}
	}

	// Test that map doesn't contain extra entries
	if len(EnumQSOCompleteMap) != len(EnumQSOCompleteList) {
		t.Errorf("Map size (%d) doesn't match list size (%d)",
			len(EnumQSOCompleteMap), len(EnumQSOCompleteList))
	}
}
