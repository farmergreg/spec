package mode

import (
	"testing"
)

func TestEnumSubModeInit(t *testing.T) {
	// Test that maps and slices are initialized
	if EnumSubModeMap == nil {
		t.Error("EnumSubModeMap was not initialized")
	}
	if EnumSubModeListAll == nil {
		t.Error("EnumSubModeListAll was not initialized")
	}
	if EnumSubModeList == nil {
		t.Error("EnumSubModeList was not initialized")
	}

	// Test that filtered list is smaller than full list
	if len(EnumSubModeList) > len(EnumSubModeListAll) {
		t.Error("Filtered list should be smaller than or equal to full list")
	}

	// Test that map contains all items from filtered list
	if len(EnumSubModeMap) != len(EnumSubModeList) {
		t.Errorf("Map size (%d) does not match filtered list size (%d)",
			len(EnumSubModeMap), len(EnumSubModeList))
	}

	// Test that each item in the filtered list is in the map
	for _, item := range EnumSubModeList {
		if mapped, exists := EnumSubModeMap[item.ID]; !exists {
			t.Errorf("Item %s from filtered list not found in map", item.ID)
		} else if mapped != item {
			t.Errorf("Item %s in map does not match item in filtered list", item.ID)
		}
	}

	// Test that filtered list only contains released and non-import-only items
	for _, item := range EnumSubModeList {
		if !bool(item.IsReleased) {
			t.Errorf("Filtered list contains unreleased item: %s", item.ID)
		}
		if bool(item.IsImportOnly) {
			t.Errorf("Filtered list contains import-only item: %s", item.ID)
		}
	}
}
