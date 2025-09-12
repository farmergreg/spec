package region

import (
	"strings"

	"github.com/hamradiolog-net/spec/v6/enum/dxccentitycode"
	"github.com/hamradiolog-net/spec/v6/internal/codegen"
)

// RegionCompositeKey works differently that almost all of the other enumerations!
// It is a COMPOSITE KEY, made up of the RegionCompositeKey code and the DXCC Entity Code.
// It is NOT the value that is stored in the ADIF field.
type RegionCompositeKey string

var _ codegen.CodeGenKey = RegionCompositeKey("")

// String returns the string representation of the RegionCompositeKey.
func (r RegionCompositeKey) String() string {
	return string(r)
}

// ADIF enums are case-insensitive.
func (r RegionCompositeKey) Compare(other RegionCompositeKey) int {
	return strings.Compare(string(r), string(other))
}

// LookupByCodeAndDXCC looks up a Region specification by its composite key (Region Code + DXCCEntityCode).
func LookupByCodeAndDXCC(code RegionCode, dxccEntityCode dxccentitycode.DXCCEntityCode) (Spec, bool) {
	if code == NONE {
		spec, ok := Lookup(RegionCompositeKey("NONE"))
		return spec, ok
	}
	spec, ok := Lookup(RegionCompositeKey(string(code) + "." + dxccEntityCode.String()))
	return spec, ok
}
