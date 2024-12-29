package contest

import (
	"testing"
)

func TestInit(t *testing.T) {
	// Test that maps and lists are not nil
	if EnumContestMap == nil {
		t.Error("EnumContestMap is nil")
	}
	if EnumContestListAll == nil {
		t.Error("EnumContestListAll is nil")
	}
	if EnumContestList == nil {
		t.Error("EnumContestList is nil")
	}

	// Test that filtered list is smaller or equal to full list
	if len(EnumContestList) > len(EnumContestListAll) {
		t.Errorf("Filtered list length (%d) is larger than full list length (%d)",
			len(EnumContestList), len(EnumContestListAll))
	}

	// Test that map contains all items from filtered list
	for _, item := range EnumContestList {
		if mapped, exists := EnumContestMap[item.ID]; !exists {
			t.Errorf("Item %s not found in map", item.ID)
		} else if mapped != item {
			t.Errorf("Map entry for %s points to different item", item.ID)
		}
	}

	// Test that map size matches filtered list size
	if len(EnumContestMap) != len(EnumContestList) {
		t.Errorf("Map size (%d) does not match filtered list size (%d)",
			len(EnumContestMap), len(EnumContestList))
	}

	// Test filtering logic
	for _, item := range EnumContestListAll {
		if bool(item.IsReleased) && !bool(item.IsImportOnly) {
			// Should be in filtered list
			found := false
			for _, filteredItem := range EnumContestList {
				if filteredItem == item {
					found = true
					break
				}
			}
			if !found {
				t.Errorf("Released non-import-only item %s not found in filtered list", item.ID)
			}
		}
	}
}
