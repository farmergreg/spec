package aditype

import (
	"testing"
)

func TestInit(t *testing.T) {
	// Test DataTypeListAll is not empty
	if len(List) == 0 {
		t.Error("DataTypeListAll should not be empty after init")
	}

	// Test DataTypeList is not empty and smaller than DataTypeListAll
	if len(ListCurrent) == 0 {
		t.Error("DataTypeList should not be empty after init")
	}
	if len(ListCurrent) >= len(List) {
		t.Error("DataTypeList should be smaller than DataTypeListAll due to filtering")
	}

	// Test DataTypeMap has same length as DataTypeList
	if len(Map) != len(ListCurrent) {
		t.Errorf("DataTypeMap length (%d) should match DataTypeList length (%d)",
			len(Map), len(ListCurrent))
	}

	// Test each item in DataTypeList meets the filtering criteria
	for _, item := range ListCurrent {
		if !bool(item.IsReleased) {
			t.Error("DataTypeList contains unreleased item:", item.ID)
		}
		if bool(item.IsImportOnly) {
			t.Error("DataTypeList contains import-only item:", item.ID)
		}
	}

	// Test each item in DataTypeList is properly mapped in DataTypeMap
	for _, item := range ListCurrent {
		mapped, exists := Map[item.ID]
		if !exists {
			t.Error("Item from DataTypeList not found in DataTypeMap:", item.ID)
		}
		if mapped != item {
			t.Error("Mapped item does not match original item for ID:", item.ID)
		}
	}
}
