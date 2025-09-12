package secondaryadministrativesubdivisionalt

import (
	"strings"

	"github.com/hamradiolog-net/spec/v6/internal/codegen"
)

type SecondaryAdministrativeSubdivisionAlt string

var _ codegen.CodeGenKey = SecondaryAdministrativeSubdivisionAlt("")

// String returns the string representation of the SecondaryAdministrativeSubdivisionAlt.
func (s SecondaryAdministrativeSubdivisionAlt) String() string {
	return string(s)
}

// Compare implements the Comparable interface.
// ADIF enums are case-insensitive.
func (t SecondaryAdministrativeSubdivisionAlt) Compare(other SecondaryAdministrativeSubdivisionAlt) int {
	return strings.Compare(strings.ToUpper(string(t)), strings.ToUpper(other.String()))
}
