package arrlsection

import (
	"testing"
)

func TestInit(t *testing.T) {
	t.Run("EnumARRLSectionListAll initialization", func(t *testing.T) {
		if len(EnumARRLSectionListAll) == 0 {
			t.Error("EnumARRLSectionListAll should not be empty")
		}
	})

	t.Run("EnumARRLSectionList filtering", func(t *testing.T) {
		if len(EnumARRLSectionList) == 0 {
			t.Error("EnumARRLSectionList should not be empty")
		}
		if len(EnumARRLSectionList) >= len(EnumARRLSectionListAll) {
			t.Error("Filtered list should be smaller than or equal to the full list")
		}

		// Verify filtering logic
		for _, item := range EnumARRLSectionList {
			if !bool(item.IsReleased) {
				t.Error("All items should be released")
			}
			if bool(item.IsImportOnly) {
				t.Error("No items should be import-only")
			}
			if !item.DeletedDate.Time.IsZero() {
				t.Error("No items should have a deletion date")
			}
		}
	})

	t.Run("EnumARRLSectionMap initialization", func(t *testing.T) {
		if len(EnumARRLSectionList) != len(EnumARRLSectionMap) {
			t.Error("Map should have same number of entries as filtered list")
		}

		// Verify map contains all items from the filtered list
		for _, item := range EnumARRLSectionList {
			mapped, exists := EnumARRLSectionMap[item.ID]
			if !exists {
				t.Errorf("Map should contain item with ID %s", item.ID)
			}
			if item != mapped {
				t.Error("Mapped item should match list item")
			}
		}
	})
}
