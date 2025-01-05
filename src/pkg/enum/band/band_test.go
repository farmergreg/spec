package band

import (
	"testing"
)

func TestInit(t *testing.T) {
	// Test that EnumBandListAll is populated
	if len(EnumBandListAll) == 0 {
		t.Error("EnumBandListAll should not be empty")
	}

	// Test that EnumBandList is populated and properly filtered
	if len(EnumBandList) == 0 {
		t.Error("EnumBandList should not be empty")
	}
	if len(EnumBandList) > len(EnumBandListAll) {
		t.Error("EnumBandList should be filtered and equal to or smaller than EnumBandListAll")
	}

	// Test that EnumBandMap is populated
	if len(EnumBandMap) != len(EnumBandList) {
		t.Errorf("EnumBandMap size (%d) should match EnumBandList size (%d)",
			len(EnumBandMap), len(EnumBandList))
	}

	// Verify that all items in EnumBandList meet the filtering criteria
	for _, item := range EnumBandList {
		if !bool(item.IsReleased) {
			t.Errorf("Band %s in EnumBandList is not released", item.ID)
		}
		if bool(item.IsImportOnly) {
			t.Errorf("Band %s in EnumBandList is import-only", item.ID)
		}
	}

	// Verify that all items in EnumBandMap are present in EnumBandList
	for bandID, mapItem := range EnumBandMap {
		found := false
		for _, listItem := range EnumBandList {
			if listItem.ID == bandID {
				found = true
				if listItem != mapItem {
					t.Errorf("Map item for band %s doesn't match list item", bandID)
				}
				break
			}
		}
		if !found {
			t.Errorf("Band %s found in map but not in list", bandID)
		}
	}
}

func TestBandwidth(t *testing.T) {
	for _, item := range EnumBandList {
		if item.Bandwidth() != item.UpperFreqMHz-item.LowerFreqMHz {
			t.Errorf("Band %s bandwidth is incorrect: %f", item.ID, item.Bandwidth())
		}
	}
}

func TestIsInBand(t *testing.T) {
	for _, item := range EnumBandList {
		if !item.IsInBand(item.LowerFreqMHz) {
			t.Errorf("Band %s should be in band", item.ID)
		}
		if !item.IsInBand(item.UpperFreqMHz) {
			t.Errorf("Band %s should be in band", item.ID)
		}
		if item.IsInBand(item.LowerFreqMHz - 1) {
			t.Errorf("Band %s should not be in band", item.ID)
		}
		if item.IsInBand(item.UpperFreqMHz + 1) {
			t.Errorf("Band %s should not be in band", item.ID)
		}
	}
}

func TestFindBandByMHz(t *testing.T) {
	for _, item := range EnumBandList {
		_, ok := FindBandByMHz(item.LowerFreqMHz)
		if !ok {
			t.Errorf("Band %s should be in band", item.ID)
		}
	}
}

func TestFindBandByMHz_NotInBand(t *testing.T) {
	for _, item := range EnumBandList {
		_, ok := FindBandByMHz(0.25)
		if ok {
			t.Errorf("Band %s should be in band", item.ID)
		}
	}
}
