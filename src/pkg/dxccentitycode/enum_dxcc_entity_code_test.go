package dxccentitycode

import (
	"testing"
)

func TestInit(t *testing.T) {
	// Test that the maps and lists are not nil
	if EnumDXCCEntityCodeMap == nil {
		t.Error("EnumDXCCEntityCodeMap is nil")
	}
	if EnumDXCCEntityCodeListAll == nil {
		t.Error("EnumDXCCEntityCodeListAll is nil")
	}
	if EnumDXCCEntityCodeList == nil {
		t.Error("EnumDXCCEntityCodeList is nil")
	}

	// Test that filtered list is smaller than or equal to the full list
	if len(EnumDXCCEntityCodeList) > len(EnumDXCCEntityCodeListAll) {
		t.Errorf("EnumDXCCEntityCodeList (%d) is larger than EnumDXCCEntityCodeListAll (%d)",
			len(EnumDXCCEntityCodeList), len(EnumDXCCEntityCodeListAll))
	}

	// Test that map size matches filtered list size
	if len(EnumDXCCEntityCodeMap) != len(EnumDXCCEntityCodeList) {
		t.Errorf("EnumDXCCEntityCodeMap size (%d) does not match EnumDXCCEntityCodeList size (%d)",
			len(EnumDXCCEntityCodeMap), len(EnumDXCCEntityCodeList))
	}

	// Test that all items in the filtered list meet the filtering criteria
	for _, item := range EnumDXCCEntityCodeList {
		if !bool(item.IsReleased) || bool(item.IsImportOnly) || bool(item.IsDeleted) {
			t.Errorf("Found invalid item in filtered list: ID=%d, Released=%v, ImportOnly=%v, Deleted=%v",
				item.ID, item.IsReleased, item.IsImportOnly, item.IsDeleted)
		}
	}

	// Test that all map entries correspond to items in the filtered list
	for code, item := range EnumDXCCEntityCodeMap {
		if item.ID != code {
			t.Errorf("Map key (%d) does not match item ID (%d)", code, item.ID)
		}
	}
}
