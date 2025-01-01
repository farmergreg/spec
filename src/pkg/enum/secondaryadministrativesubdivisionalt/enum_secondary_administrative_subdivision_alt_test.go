package secondaryadministrativesubdivisionalt

import (
	"testing"
)

func TestInit(t *testing.T) {
	// Test that the maps and lists are initialized
	if EnumSecondaryAdministrativeSubdivisionAltMap == nil {
		t.Error("EnumSecondaryAdministrativeSubdivisionAltMap was not initialized")
	}

	if EnumSecondaryAdministrativeSubdivisionAltListAll == nil {
		t.Error("EnumSecondaryAdministrativeSubdivisionAltListAll was not initialized")
	}

	if EnumSecondaryAdministrativeSubdivisionAltList == nil {
		t.Error("EnumSecondaryAdministrativeSubdivisionAltList was not initialized")
	}

	// Test that filtered list only contains released and non-import-only items
	for _, item := range EnumSecondaryAdministrativeSubdivisionAltList {
		if !bool(item.IsReleased) {
			t.Errorf("Found unreleased item in filtered list: %v", item.ID)
		}
		if bool(item.IsImportOnly) {
			t.Errorf("Found import-only item in filtered list: %v", item.ID)
		}
	}

	// Test that map contains all items from the filtered list
	if len(EnumSecondaryAdministrativeSubdivisionAltMap) != len(EnumSecondaryAdministrativeSubdivisionAltList) {
		t.Errorf("Map size (%d) does not match filtered list size (%d)",
			len(EnumSecondaryAdministrativeSubdivisionAltMap),
			len(EnumSecondaryAdministrativeSubdivisionAltList))
	}

	// Verify each item in the list exists in the map
	for _, item := range EnumSecondaryAdministrativeSubdivisionAltList {
		if mapItem, exists := EnumSecondaryAdministrativeSubdivisionAltMap[item.ID]; !exists {
			t.Errorf("Item %v exists in list but not in map", item.ID)
		} else if mapItem != item {
			t.Errorf("Map item pointer does not match list item pointer for ID %v", item.ID)
		}
	}

	// Verify the filtered list is smaller than or equal to the full list
	if len(EnumSecondaryAdministrativeSubdivisionAltList) > len(EnumSecondaryAdministrativeSubdivisionAltListAll) {
		t.Error("Filtered list is larger than the full list")
	}
}
