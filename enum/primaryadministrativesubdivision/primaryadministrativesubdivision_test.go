package primaryadministrativesubdivision

import (
	"testing"

	"github.com/hamradiolog-net/spec/v6/enum/dxccentitycode"
)

func TestLookupLookupByCodeAndDXCC(t *testing.T) {
	tests := []struct {
		code   PrimaryAdministrativeSubdivisionCode
		dxcc   dxccentitycode.DXCCEntityCode
		wantOk bool
	}{
		{PrimaryAdministrativeSubdivision_001, 5, true},
		{PrimaryAdministrativeSubdivision_ZVO, 504, true},
		{"XX", 9999, false},
	}
	for _, tt := range tests {
		_, ok := LookupByCodeAndDXCC(tt.code, tt.dxcc)
		if ok != tt.wantOk {
			t.Errorf("LookupPrimaryAdministrativeSubdivision(%v, %v) = %v, want %v", tt.code, tt.dxcc, ok, tt.wantOk)
		}
	}
}
