package secondaryadministrativesubdivision

import (
	"strings"

	"github.com/hamradiolog-net/spec/v6/internal/codegen"
)

type SecondaryAdministrativeSubdivision string

var _ codegen.CodeGenKey = SecondaryAdministrativeSubdivision("")

// String returns the string representation of the SecondaryAdministrativeSubdivision.
func (s SecondaryAdministrativeSubdivision) String() string {
	return string(s)
}

// Compare implements the Comparable interface.
// ADIF enums are case-insensitive.
func (t SecondaryAdministrativeSubdivision) Compare(other SecondaryAdministrativeSubdivision) int {
	return strings.Compare(strings.ToUpper(string(t)), strings.ToUpper(other.String()))
}
