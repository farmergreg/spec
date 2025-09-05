package region

import (
	"github.com/hamradiolog-net/adif-spec/v6/enum/dxccentitycode"
	"github.com/hamradiolog-net/adif-spec/v6/internal/codegen"
)

// RegionCompositeKey works differently that almost all of the other enumerations!
// It is a COMPOSITE KEY, made up of the RegionCompositeKey code and the DXCC Entity Code.
// It is NOT the value that is stored in the ADIF field.
type RegionCompositeKey string

var _ codegen.CodeGeneratorEnumValue = RegionCompositeKey("")

func (r RegionCompositeKey) String() string {
	return string(r)
}

// LookupRegion looks up a Region specification by its composite key (RegionCompositeKey + DXCCEntityCode).
func LookupRegion(code RegionCode, dxccEntityCode dxccentitycode.DXCCEntityCode) (Spec, bool) {
	spec, ok := RegionCompositeKeyMap[RegionCompositeKey(string(code)+"."+dxccEntityCode.String())]
	return spec, ok
}
