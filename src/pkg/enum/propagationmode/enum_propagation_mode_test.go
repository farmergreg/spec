package propagationmode

import (
	"testing"
)

func TestInit(t *testing.T) {
	// Test that the lists are not empty
	if len(EnumPropagationModeListAll) == 0 {
		t.Error("EnumPropagationModeListAll should not be empty")
	}

	if len(EnumPropagationModeList) == 0 {
		t.Error("EnumPropagationModeList should not be empty")
	}

	if len(EnumPropagationModeMap) == 0 {
		t.Error("EnumPropagationModeMap should not be empty")
	}

	// Test that filtered list is smaller or equal to the full list
	if len(EnumPropagationModeList) > len(EnumPropagationModeListAll) {
		t.Errorf("Filtered list (%d) should not be larger than full list (%d)",
			len(EnumPropagationModeList), len(EnumPropagationModeListAll))
	}

	// Test that map size matches filtered list size
	if len(EnumPropagationModeMap) != len(EnumPropagationModeList) {
		t.Errorf("Map size (%d) should match filtered list size (%d)",
			len(EnumPropagationModeMap), len(EnumPropagationModeList))
	}

	// Test that all items in the filtered list meet the criteria
	for _, item := range EnumPropagationModeList {
		if !bool(item.IsReleased) {
			t.Errorf("Item %s in filtered list should be released", item.ID)
		}
		if bool(item.IsImportOnly) {
			t.Errorf("Item %s in filtered list should not be import-only", item.ID)
		}
	}

	// Test that all map entries correspond to list entries
	for _, item := range EnumPropagationModeList {
		if mapItem, exists := EnumPropagationModeMap[item.ID]; !exists {
			t.Errorf("Item %s from list not found in map", item.ID)
		} else if mapItem != item {
			t.Errorf("Map entry for %s doesn't match list entry", item.ID)
		}
	}
}
