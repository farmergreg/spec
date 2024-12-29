package continent

import (
	"testing"
)

func TestInit(t *testing.T) {
	// Test that maps and slices are initialized
	if EnumContinentMap == nil {
		t.Error("EnumContinentMap was not initialized")
	}
	if EnumContinentListAll == nil {
		t.Error("EnumContinentListAll was not initialized")
	}
	if EnumContinentList == nil {
		t.Error("EnumContinentList was not initialized")
	}

	// Test that filtered list is smaller or equal to full list
	if len(EnumContinentList) > len(EnumContinentListAll) {
		t.Errorf("Filtered list (%d) should not be larger than full list (%d)",
			len(EnumContinentList), len(EnumContinentListAll))
	}

	// Test that map contains all items from filtered list
	for _, item := range EnumContinentList {
		if mapped, exists := EnumContinentMap[item.ID]; !exists {
			t.Errorf("Item %s missing from map", item.ID)
		} else if mapped != item {
			t.Errorf("Map entry for %s points to different item", item.ID)
		}
	}

	// Test that map size matches filtered list size
	if len(EnumContinentMap) != len(EnumContinentList) {
		t.Errorf("Map size (%d) does not match filtered list size (%d)",
			len(EnumContinentMap), len(EnumContinentList))
	}

	// Test that filtered list only contains released and non-import-only items
	for _, item := range EnumContinentList {
		if !bool(item.IsReleased) {
			t.Errorf("Filtered list contains unreleased item: %s", item.ID)
		}
		if bool(item.IsImportOnly) {
			t.Errorf("Filtered list contains import-only item: %s", item.ID)
		}
	}
}
