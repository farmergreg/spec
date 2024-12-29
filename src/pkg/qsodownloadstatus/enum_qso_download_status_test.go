package qsodownloadstatus

import (
	"testing"
)

func TestInit(t *testing.T) {
	// Test that maps and slices are initialized
	if EnumQSODownloadStatusMap == nil {
		t.Error("EnumQSODownloadStatusMap was not initialized")
	}
	if EnumQSODownloadStatusListAll == nil {
		t.Error("EnumQSODownloadStatusListAll was not initialized")
	}
	if EnumQSODownloadStatusList == nil {
		t.Error("EnumQSODownloadStatusList was not initialized")
	}

	// Test that filtered list only contains released and non-import-only items
	for _, item := range EnumQSODownloadStatusList {
		if !bool(item.IsReleased) {
			t.Errorf("Found unreleased item in filtered list: %s", item.ID)
		}
		if bool(item.IsImportOnly) {
			t.Errorf("Found import-only item in filtered list: %s", item.ID)
		}
	}

	// Test that map contains all items from the filtered list
	if len(EnumQSODownloadStatusMap) != len(EnumQSODownloadStatusList) {
		t.Errorf("Map size (%d) does not match filtered list size (%d)",
			len(EnumQSODownloadStatusMap), len(EnumQSODownloadStatusList))
	}

	// Verify each item in the list exists in the map
	for _, item := range EnumQSODownloadStatusList {
		if mapItem, exists := EnumQSODownloadStatusMap[item.ID]; !exists {
			t.Errorf("Item %s exists in list but not in map", item.ID)
		} else if mapItem != item {
			t.Errorf("Map item pointer does not match list item pointer for %s", item.ID)
		}
	}

	// Verify all items in EnumQSODownloadStatusListAll are valid
	for _, item := range EnumQSODownloadStatusListAll {
		if item == nil {
			t.Error("Found nil item in EnumQSODownloadStatusListAll")
		}
		if item.ID == "" {
			t.Error("Found item with empty ID in EnumQSODownloadStatusListAll")
		}
	}
}
