package adifield

import (
	"strings"
	"testing"
)

func TestInit(t *testing.T) {
	// Test that FieldListAll was populated
	if len(List) == 0 {
		t.Error("FieldListAll is empty")
	}

	// Test USERDEFn was renamed to USERDEF1
	found := false
	for _, item := range List {
		if item.ID == "USERDEFn" {
			t.Error("Found USERDEFn, should have been renamed to USERDEF1")
		}
		if item.ID == USERDEF1 {
			found = true
		}
	}
	if !found {
		t.Error("USERDEF1 not found in FieldListAll")
	}

	// Test FieldList only contains released and non-import-only fields
	for _, item := range ListCurrent {
		if !bool(item.IsReleased) {
			t.Errorf("Found unreleased field in FieldList: %s", item.ID)
		}
		if bool(item.IsImportOnly) {
			t.Errorf("Found import-only field in FieldList: %s", item.ID)
		}
	}

	// Test FieldMap contains all items from FieldList
	if len(Map) != len(ListCurrent) {
		t.Errorf("FieldMap length (%d) does not match FieldList length (%d)", len(Map), len(ListCurrent))
	}

	// Test all items in FieldList are in FieldMap
	for _, item := range ListCurrent {
		if _, exists := Map[item.ID]; !exists {
			t.Errorf("Field %s from FieldList not found in FieldMap", item.ID)
		}
	}

	// Test that APP_ fields are handled correctly
	for _, item := range ListCurrent {
		if strings.HasPrefix(string(item.ID), APP_) {
			// Verify APP_ fields follow the correct pattern
			parts := strings.Split(string(item.ID), "_")
			if len(parts) < 3 {
				t.Errorf("APP_ field %s does not follow pattern APP_PROGRAMID_FIELDNAME", item.ID)
			}
		}
	}
}
