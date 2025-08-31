package region

import (
	"testing"

	"github.com/hamradiolog-net/adif-spec/v6/enum/dxccentitycode"
)

func TestLookupRegion(t *testing.T) {
	tests := []struct {
		region RegionCode
		dxcc   dxccentitycode.DXCCEntityCode
		wantOk bool
	}{
		{"AI", 248, true},
		{"BI", 259, true},
		{"XX", 9999, false},
	}
	for _, tt := range tests {
		_, ok := LookupRegion(tt.region, tt.dxcc)
		if ok != tt.wantOk {
			t.Errorf("LookupRegion(%v, %v) = %v, want %v", tt.region, tt.dxcc, ok, tt.wantOk)
		}
	}
}
