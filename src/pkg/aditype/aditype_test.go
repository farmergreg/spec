package aditype

import (
	"testing"
)

func TestInit(t *testing.T) {
	// Test DataTypeListAll is not empty
	if len(DataTypeListAll) == 0 {
		t.Error("DataTypeListAll should not be empty after init")
	}

	// Test DataTypeList is not empty and smaller than DataTypeListAll
	if len(DataTypeList) == 0 {
		t.Error("DataTypeList should not be empty after init")
	}
	if len(DataTypeList) >= len(DataTypeListAll) {
		t.Error("DataTypeList should be smaller than DataTypeListAll due to filtering")
	}

	// Test DataTypeMap has same length as DataTypeList
	if len(DataTypeMap) != len(DataTypeList) {
		t.Errorf("DataTypeMap length (%d) should match DataTypeList length (%d)",
			len(DataTypeMap), len(DataTypeList))
	}

	// Test each item in DataTypeList meets the filtering criteria
	for _, item := range DataTypeList {
		if !bool(item.IsReleased) {
			t.Error("DataTypeList contains unreleased item:", item.ID)
		}
		if bool(item.IsImportOnly) {
			t.Error("DataTypeList contains import-only item:", item.ID)
		}
	}

	// Test each item in DataTypeList is properly mapped in DataTypeMap
	for _, item := range DataTypeList {
		mapped, exists := DataTypeMap[item.ID]
		if !exists {
			t.Error("Item from DataTypeList not found in DataTypeMap:", item.ID)
		}
		if mapped != item {
			t.Error("Mapped item does not match original item for ID:", item.ID)
		}
	}
}
