package spec

import (
	"testing"
)

func TestInit(t *testing.T) {
	// Test that both lists were initialized
	if EnumPrimaryAdministrativeSubdivisionListAll == nil {
		t.Error("EnumPrimaryAdministrativeSubdivisionListAll was not initialized")
	}
	if EnumPrimaryAdministrativeSubdivisionList == nil {
		t.Error("EnumPrimaryAdministrativeSubdivisionList was not initialized")
	}

	// Test that filtered list is smaller than or equal to the full list
	if len(EnumPrimaryAdministrativeSubdivisionList) > len(EnumPrimaryAdministrativeSubdivisionListAll) {
		t.Errorf("Filtered list (%d) should not be larger than full list (%d)",
			len(EnumPrimaryAdministrativeSubdivisionList),
			len(EnumPrimaryAdministrativeSubdivisionListAll))
	}

	// Test that filtered list only contains valid items
	for _, item := range EnumPrimaryAdministrativeSubdivisionList {
		if !bool(item.IsReleased) {
			t.Errorf("Found unreleased item in filtered list: %v", item.ID)
		}
		if bool(item.IsImportOnly) {
			t.Errorf("Found import-only item in filtered list: %v", item.ID)
		}
		if bool(item.IsDeleted) {
			t.Errorf("Found deleted item in filtered list: %v", item.ID)
		}
	}

	// Test that the filtered list's capacity matches its length (was clipped)
	if cap(EnumPrimaryAdministrativeSubdivisionList) != len(EnumPrimaryAdministrativeSubdivisionList) {
		t.Errorf("Filtered list was not clipped: length=%d, capacity=%d",
			len(EnumPrimaryAdministrativeSubdivisionList),
			cap(EnumPrimaryAdministrativeSubdivisionList))
	}
}
