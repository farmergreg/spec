package primaryadministrativesubdivision

import (
	"github.com/hamradiolog-net/adif-spec/v6/enum/dxccentitycode"
	"github.com/hamradiolog-net/adif-spec/v6/internal/codegen"
)

type PrimaryAdministrativeSubdivisionCompositeKey string

var _ codegen.CodeGenKey = PrimaryAdministrativeSubdivisionCompositeKey("")

func (p PrimaryAdministrativeSubdivisionCompositeKey) String() string {
	return string(p)
}

// LookupByCodeAndDXCC looks up a Primary Administrative Subdivision specification by its composite key (Code + DXCCEntityCode).
func LookupByCodeAndDXCC(code PrimaryAdministrativeSubdivisionCode, dxccEntityCode dxccentitycode.DXCCEntityCode) (Spec, bool) {
	spec, ok := Lookup(PrimaryAdministrativeSubdivisionCompositeKey(string(code) + "." + dxccEntityCode.String()))
	return spec, ok
}
