package mode

import (
	"testing"
)

func TestInit(t *testing.T) {
	// Test that EnumModeListAll is populated
	if len(EnumModeListAll) == 0 {
		t.Error("EnumModeListAll should not be empty")
	}

	// Test that EnumModeList is populated and smaller than or equal to EnumModeListAll
	if len(EnumModeList) == 0 {
		t.Error("EnumModeList should not be empty")
	}
	if len(EnumModeList) > len(EnumModeListAll) {
		t.Errorf("EnumModeList length (%d) should not be greater than EnumModeListAll length (%d)",
			len(EnumModeList), len(EnumModeListAll))
	}

	// Test that EnumModeMap is populated and matches EnumModeList length
	if len(EnumModeMap) != len(EnumModeList) {
		t.Errorf("EnumModeMap length (%d) should match EnumModeList length (%d)",
			len(EnumModeMap), len(EnumModeList))
	}

	// Test that all items in EnumModeList are properly filtered
	for _, item := range EnumModeList {
		if !bool(item.IsReleased) {
			t.Errorf("Item %s in EnumModeList should be released", item.ID)
		}
		if bool(item.IsImportOnly) {
			t.Errorf("Item %s in EnumModeList should not be import-only", item.ID)
		}
	}

	// Test that all items in EnumModeMap are present in EnumModeList
	for id, item := range EnumModeMap {
		found := false
		for _, listItem := range EnumModeList {
			if listItem.ID == id {
				found = true
				if listItem != item {
					t.Errorf("Item mismatch in map and list for ID %s", id)
				}
				break
			}
		}
		if !found {
			t.Errorf("Item %s found in map but not in list", id)
		}
	}
}
