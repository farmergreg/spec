package qslvia

import (
	"testing"
)

func TestInit(t *testing.T) {
	// Test that the maps and lists are not nil
	if EnumQSLViaMap == nil {
		t.Error("EnumQSLViaMap is nil")
	}
	if EnumQSLViaListAll == nil {
		t.Error("EnumQSLViaListAll is nil")
	}
	if EnumQSLViaList == nil {
		t.Error("EnumQSLViaList is nil")
	}

	// Test that EnumQSLViaList only contains released and non-import-only items
	for _, item := range EnumQSLViaList {
		if !bool(item.IsReleased) {
			t.Errorf("Found unreleased item in EnumQSLViaList: %s", item.ID)
		}
		if bool(item.IsImportOnly) {
			t.Errorf("Found import-only item in EnumQSLViaList: %s", item.ID)
		}
	}

	// Test that all items in EnumQSLViaMap are present in EnumQSLViaList
	for id, mapItem := range EnumQSLViaMap {
		found := false
		for _, listItem := range EnumQSLViaList {
			if listItem.ID == id {
				found = true
				// Check that the items are the same instance
				if mapItem != listItem {
					t.Errorf("Map and list items don't match for ID %s", id)
				}
				break
			}
		}
		if !found {
			t.Errorf("Item %s found in map but not in list", id)
		}
	}

	// Test that all items in EnumQSLViaList are present in EnumQSLViaMap
	for _, listItem := range EnumQSLViaList {
		if mapItem, exists := EnumQSLViaMap[listItem.ID]; !exists {
			t.Errorf("Item %s found in list but not in map", listItem.ID)
		} else if mapItem != listItem {
			t.Errorf("List and map items don't match for ID %s", listItem.ID)
		}
	}
}
