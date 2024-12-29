package dxccentitycode

import (
	"testing"
)

func TestInit(t *testing.T) {
	// Test that the maps and lists are not nil
	if EnumDXCCMap == nil {
		t.Error("EnumDXCCMap is nil")
	}
	if EnumDXCCListAll == nil {
		t.Error("EnumDXCCListAll is nil")
	}
	if EnumDXCCList == nil {
		t.Error("EnumDXCCList is nil")
	}

	// Test that filtered list is smaller than or equal to the full list
	if len(EnumDXCCList) > len(EnumDXCCListAll) {
		t.Errorf("Filtered list (%d) is larger than full list (%d)",
			len(EnumDXCCList), len(EnumDXCCListAll))
	}

	// Test that map size matches filtered list size
	if len(EnumDXCCMap) != len(EnumDXCCList) {
		t.Errorf("Map size (%d) does not match filtered list size (%d)",
			len(EnumDXCCMap), len(EnumDXCCList))
	}

	// Test that all items in the filtered list meet the filtering criteria
	for _, item := range EnumDXCCList {
		if !bool(item.IsReleased) || bool(item.IsImportOnly) || bool(item.IsDeleted) {
			t.Errorf("Found invalid item in filtered list: ID=%d, Released=%v, ImportOnly=%v, Deleted=%v",
				item.ID, item.IsReleased, item.IsImportOnly, item.IsDeleted)
		}
	}

	// Test that all map entries correspond to items in the filtered list
	for code, item := range EnumDXCCMap {
		if item.ID != code {
			t.Errorf("Map key (%d) does not match item ID (%d)", code, item.ID)
		}
	}
}
