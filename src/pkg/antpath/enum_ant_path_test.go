package antpath

import (
	"testing"
)

func TestEnumAntPathInitialization(t *testing.T) {
	// Test that the global variables are properly initialized
	t.Run("EnumAntPathListAll initialization", func(t *testing.T) {
		if len(EnumAntPathListAll) == 0 {
			t.Error("EnumAntPathListAll should not be empty after initialization")
		}
	})

	t.Run("EnumAntPathList filtering", func(t *testing.T) {
		if len(EnumAntPathList) == 0 {
			t.Error("EnumAntPathList should not be empty after initialization")
		}

		// Verify that filtered list only contains released and non-import-only items
		for _, item := range EnumAntPathList {
			if !bool(item.IsReleased) {
				t.Errorf("EnumAntPathList contains unreleased item: %s", item.ID)
			}
			if bool(item.IsImportOnly) {
				t.Errorf("EnumAntPathList contains import-only item: %s", item.ID)
			}
		}
	})

	t.Run("EnumAntPathMap initialization", func(t *testing.T) {
		if len(EnumAntPathMap) == 0 {
			t.Error("EnumAntPathMap should not be empty after initialization")
		}

		// Verify map contains all items from the filtered list
		if len(EnumAntPathMap) != len(EnumAntPathList) {
			t.Errorf("EnumAntPathMap size (%d) does not match EnumAntPathList size (%d)",
				len(EnumAntPathMap), len(EnumAntPathList))
		}

		// Verify each item in the list exists in the map
		for _, item := range EnumAntPathList {
			if mapItem, exists := EnumAntPathMap[item.ID]; !exists {
				t.Errorf("Item %s exists in list but not in map", item.ID)
			} else if mapItem != item {
				t.Errorf("Map item pointer mismatch for %s", item.ID)
			}
		}
	})
}
