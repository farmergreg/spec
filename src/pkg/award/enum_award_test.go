package award

import (
	"testing"
)

func TestInit(t *testing.T) {
	t.Run("EnumAwardListAll initialization", func(t *testing.T) {
		if len(EnumAwardListAll) == 0 {
			t.Error("EnumAwardListAll should not be empty")
		}
	})

	t.Run("EnumAwardList initialization", func(t *testing.T) {
		if len(EnumAwardList) != 0 {
			t.Error("EnumAwardList should be empty")
		}
	})

	t.Run("EnumAwardMap initialization", func(t *testing.T) {
		if len(EnumAwardMap) != 0 {
			t.Error("EnumAwardMap should be empty")
		}
	})
}
