package secondaryadministrativesubdivision

import (
	"testing"
)

func TestInit(t *testing.T) {
	// Test that the maps and lists are initialized
	if EnumSecondaryAdministrativeSubdivisionMap == nil {
		t.Error("EnumSecondaryAdministrativeSubdivisionMap was not initialized")
	}

	if EnumSecondaryAdministrativeSubdivisionListAll == nil {
		t.Error("EnumSecondaryAdministrativeSubdivisionListAll was not initialized")
	}

	if EnumSecondaryAdministrativeSubdivisionList == nil {
		t.Error("EnumSecondaryAdministrativeSubdivisionList was not initialized")
	}

	// Test that filtered list only contains released and non-import-only items
	for _, item := range EnumSecondaryAdministrativeSubdivisionList {
		if !bool(item.IsReleased) {
			t.Errorf("Found unreleased item in filtered list: %v", item.ID)
		}
		if bool(item.IsImportOnly) {
			t.Errorf("Found import-only item in filtered list: %v", item.ID)
		}
	}

	// Test that map contains all items from the filtered list
	if len(EnumSecondaryAdministrativeSubdivisionMap) != len(EnumSecondaryAdministrativeSubdivisionList) {
		t.Errorf("Map size (%d) does not match filtered list size (%d)",
			len(EnumSecondaryAdministrativeSubdivisionMap),
			len(EnumSecondaryAdministrativeSubdivisionList))
	}

	// Verify each item in the list exists in the map
	for _, item := range EnumSecondaryAdministrativeSubdivisionList {
		if _, exists := EnumSecondaryAdministrativeSubdivisionMap[item.ID]; !exists {
			t.Errorf("Item %v exists in list but not in map", item.ID)
		}
	}

	// Verify the filtered list is smaller than or equal to the full list
	if len(EnumSecondaryAdministrativeSubdivisionList) > len(EnumSecondaryAdministrativeSubdivisionListAll) {
		t.Error("Filtered list is larger than the full list")
	}
}
