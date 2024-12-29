package qslrcvd

import (
	"testing"
)

func TestInit(t *testing.T) {
	// Test that the maps and lists are not empty
	if len(EnumQSLRcvdListAll) == 0 {
		t.Error("EnumQSLRcvdListAll should not be empty")
	}

	if len(EnumQSLRcvdList) == 0 {
		t.Error("EnumQSLRcvdList should not be empty")
	}

	if len(EnumQSLRcvdMap) == 0 {
		t.Error("EnumQSLRcvdMap should not be empty")
	}

	// Test that filtered list only contains released and non-import-only items
	for _, item := range EnumQSLRcvdList {
		if !bool(item.IsReleased) {
			t.Errorf("EnumQSLRcvdList contains unreleased item: %s", item.ID)
		}
		if bool(item.IsImportOnly) {
			t.Errorf("EnumQSLRcvdList contains import-only item: %s", item.ID)
		}
	}

	// Test that map contains all items from the filtered list
	for _, item := range EnumQSLRcvdList {
		if mapped, exists := EnumQSLRcvdMap[item.ID]; !exists {
			t.Errorf("Item %s from list not found in map", item.ID)
		} else if mapped != item {
			t.Errorf("Map entry for %s does not match list item", item.ID)
		}
	}

	// Test that map only contains items from the filtered list
	for id := range EnumQSLRcvdMap {
		found := false
		for _, item := range EnumQSLRcvdList {
			if item.ID == id {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("Map contains item %s not present in filtered list", id)
		}
	}
}
