package primaryadministrativesubdivision

import (
	"strings"

	"github.com/hamradiolog-net/spec/v6/enum/dxccentitycode"
	"github.com/hamradiolog-net/spec/v6/internal/codegen"
)

type PrimaryAdministrativeSubdivisionCompositeKey string

var _ codegen.CodeGenKey = PrimaryAdministrativeSubdivisionCompositeKey("")

// String returns the string representation of the PrimaryAdministrativeSubdivisionCompositeKey.
func (p PrimaryAdministrativeSubdivisionCompositeKey) String() string {
	return string(p)
}

// ADIF enums are case-insensitive.
func (p PrimaryAdministrativeSubdivisionCode) Compare(other PrimaryAdministrativeSubdivisionCode) int {
	return strings.Compare(string(p), string(other))
}

// LookupByCodeAndDXCC looks up a Primary Administrative Subdivision specification by its composite key (Code + DXCCEntityCode).
func LookupByCodeAndDXCC(code PrimaryAdministrativeSubdivisionCode, dxccEntityCode dxccentitycode.DXCCEntityCode) (Spec, bool) {
	spec, ok := Lookup(PrimaryAdministrativeSubdivisionCompositeKey(string(code) + "." + dxccEntityCode.String()))
	return spec, ok
}
