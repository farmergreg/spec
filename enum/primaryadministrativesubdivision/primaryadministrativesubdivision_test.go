package primaryadministrativesubdivision

import (
	"testing"

	"github.com/hamradiolog-net/adif-spec/v8/enum/dxccentitycode"
)

func TestLookupRegion(t *testing.T) {
	tests := []struct {
		code   PrimaryAdministrativeSubdivisionCode
		dxcc   dxccentitycode.DXCCEntityCode
		wantOk bool
	}{
		{"NS", 1, true},
		{"QC", 1, true},
		{"XX", 9999, false},
	}
	for _, tt := range tests {
		_, ok := LookupPrimaryAdministrativeSubdivision(tt.code, tt.dxcc)
		if ok != tt.wantOk {
			t.Errorf("LookupPrimaryAdministrativeSubdivision(%v, %v) = %v, want %v", tt.code, tt.dxcc, ok, tt.wantOk)
		}
	}
}
