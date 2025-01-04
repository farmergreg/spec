package qsouploadstatus

import (
	"testing"
)

func TestInit(t *testing.T) {
	// Test that the maps and lists are not nil
	if EnumQSOUploadStatusMap == nil {
		t.Error("EnumQSOUploadStatusMap is nil")
	}
	if EnumQSOUploadStatusListAll == nil {
		t.Error("EnumQSOUploadStatusListAll is nil")
	}
	if EnumQSOUploadStatusList == nil {
		t.Error("EnumQSOUploadStatusList is nil")
	}

	// Test that filtered list only contains released and non-import-only items
	for _, item := range EnumQSOUploadStatusList {
		if !bool(item.IsReleased) {
			t.Errorf("Found unreleased item in filtered list: %s", item.ID)
		}
		if bool(item.IsImportOnly) {
			t.Errorf("Found import-only item in filtered list: %s", item.ID)
		}
	}

	// Test that map contains all items from the filtered list
	if len(EnumQSOUploadStatusMap) != len(EnumQSOUploadStatusList) {
		t.Errorf("Map size (%d) does not match filtered list size (%d)",
			len(EnumQSOUploadStatusMap), len(EnumQSOUploadStatusList))
	}

	// Verify each item in the list exists in the map
	for _, item := range EnumQSOUploadStatusList {
		if mapItem, exists := EnumQSOUploadStatusMap[item.ID]; !exists {
			t.Errorf("Item %s exists in list but not in map", item.ID)
		} else if mapItem != item {
			t.Errorf("Item %s in map does not match item in list", item.ID)
		}
	}

	// Verify the filtered list is smaller than or equal to the full list
	if len(EnumQSOUploadStatusList) > len(EnumQSOUploadStatusListAll) {
		t.Errorf("Filtered list (%d) is larger than full list (%d)",
			len(EnumQSOUploadStatusList), len(EnumQSOUploadStatusListAll))
	}
}
