package awardsponsorprefix

import (
	"testing"
)

func TestInit(t *testing.T) {
	// Test that the maps and slices are initialized
	if EnumAwardSponsorMap == nil {
		t.Error("EnumAwardSponsorMap was not initialized")
	}

	if EnumAwardSponsorListAll == nil {
		t.Error("EnumAwardSponsorListAll was not initialized")
	}

	if EnumAwardSponsorList == nil {
		t.Error("EnumAwardSponsorList was not initialized")
	}

	// Test that EnumAwardSponsorList only contains released and non-import-only items
	for _, item := range EnumAwardSponsorList {
		if !bool(item.IsReleased) {
			t.Errorf("Found unreleased item in EnumAwardSponsorList: %s", item.IDPrefix)
		}
		if bool(item.IsImportOnly) {
			t.Errorf("Found import-only item in EnumAwardSponsorList: %s", item.IDPrefix)
		}
	}

	// Test that map entries correspond to list entries
	for _, item := range EnumAwardSponsorList {
		mapItem, exists := EnumAwardSponsorMap[item.IDPrefix]
		if !exists {
			t.Errorf("Item %s exists in list but not in map", item.IDPrefix)
		}
		if mapItem != item {
			t.Errorf("Map item pointer %p does not match list item pointer %p for %s", mapItem, item, item.IDPrefix)
		}
	}

	// Test that map doesn't contain extra entries
	if len(EnumAwardSponsorMap) != len(EnumAwardSponsorList) {
		t.Errorf("Map size (%d) does not match list size (%d)", len(EnumAwardSponsorMap), len(EnumAwardSponsorList))
	}
}
