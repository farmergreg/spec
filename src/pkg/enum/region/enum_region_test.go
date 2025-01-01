package region

import (
	"testing"
)

func TestInit(t *testing.T) {
	// Test that maps and slices are initialized
	if EnumRegionMap == nil {
		t.Error("EnumRegionMap was not initialized")
	}
	if EnumRegionListAll == nil {
		t.Error("EnumRegionListAll was not initialized")
	}
	if EnumRegionList == nil {
		t.Error("EnumRegionList was not initialized")
	}

	// Test that EnumRegionList only contains released and non-import-only items
	for _, item := range EnumRegionList {
		if !bool(item.IsReleased) {
			t.Errorf("EnumRegionList contains unreleased item: %s", item.ID)
		}
		if bool(item.IsImportOnly) {
			t.Errorf("EnumRegionList contains import-only item: %s", item.ID)
		}
	}

	// Test that EnumRegionMap contains all items from EnumRegionList
	if len(EnumRegionMap) != len(EnumRegionList) {
		t.Errorf("EnumRegionMap size (%d) does not match EnumRegionList size (%d)",
			len(EnumRegionMap), len(EnumRegionList))
	}

	// Verify each item in EnumRegionList exists in EnumRegionMap
	for _, item := range EnumRegionList {
		if mapItem, exists := EnumRegionMap[item.ID]; !exists {
			t.Errorf("Item %s exists in EnumRegionList but not in EnumRegionMap", item.ID)
		} else if mapItem != item {
			t.Errorf("Item %s in map does not match item in list", item.ID)
		}
	}
}
