package morsekey

import (
	"testing"
)

func TestMorseKeyInit(t *testing.T) {
	// Test that the lists are not empty
	if len(EnumMorseKeyListAll) == 0 {
		t.Error("EnumMorseKeyListAll should not be empty")
	}

	if len(EnumMorseKeyList) == 0 {
		t.Error("EnumMorseKeyList should not be empty")
	}

	if len(EnumMorseKeyMap) == 0 {
		t.Error("EnumMorseKeyMap should not be empty")
	}

	// Test that filtered list is subset of all list
	if len(EnumMorseKeyList) > len(EnumMorseKeyListAll) {
		t.Error("Filtered list cannot be larger than complete list")
	}

	// Test that map entries correspond to filtered list
	if len(EnumMorseKeyMap) != len(EnumMorseKeyList) {
		t.Errorf("Map size (%d) should equal filtered list size (%d)",
			len(EnumMorseKeyMap), len(EnumMorseKeyList))
	}

	// Verify each item in filtered list exists in map
	for _, item := range EnumMorseKeyList {
		if mapped, exists := EnumMorseKeyMap[item.ID]; !exists {
			t.Errorf("Item %s from filtered list not found in map", item.ID)
		} else if mapped != item {
			t.Errorf("Map entry for %s does not match list item", item.ID)
		}
	}

	// Verify filtering logic
	for _, item := range EnumMorseKeyListAll {
		if bool(item.IsReleased) && !bool(item.IsImportOnly) {
			// Should be in filtered list
			found := false
			for _, filteredItem := range EnumMorseKeyList {
				if filteredItem.ID == item.ID {
					found = true
					break
				}
			}
			if !found {
				t.Errorf("Released non-import item %s not found in filtered list", item.ID)
			}
		}
	}
}
