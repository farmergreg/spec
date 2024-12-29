package qslsent

import (
	"testing"
)

func TestInit(t *testing.T) {
	// Test that the maps and lists are properly initialized
	if len(EnumQSLSentListAll) == 0 {
		t.Error("EnumQSLSentListAll should not be empty")
	}

	if len(EnumQSLSentList) == 0 {
		t.Error("EnumQSLSentList should not be empty")
	}

	if len(EnumQSLSentMap) == 0 {
		t.Error("EnumQSLSentMap should not be empty")
	}

	// Test that EnumQSLSentList only contains released and non-import-only items
	for _, item := range EnumQSLSentList {
		if !bool(item.IsReleased) {
			t.Errorf("Found unreleased item in EnumQSLSentList: %s", item.ID)
		}
		if bool(item.IsImportOnly) {
			t.Errorf("Found import-only item in EnumQSLSentList: %s", item.ID)
		}
	}

	// Test that map entries correspond to list entries
	for _, item := range EnumQSLSentList {
		mapItem, exists := EnumQSLSentMap[item.ID]
		if !exists {
			t.Errorf("Item %s exists in list but not in map", item.ID)
		}
		if mapItem != item {
			t.Errorf("Map item pointer %p does not match list item pointer %p for ID %s", mapItem, item, item.ID)
		}
	}

	// Test that map doesn't contain extra entries
	if len(EnumQSLSentMap) != len(EnumQSLSentList) {
		t.Errorf("Map size (%d) does not match list size (%d)", len(EnumQSLSentMap), len(EnumQSLSentList))
	}
}
