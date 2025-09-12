package secondaryadministrativesubdivision

import (
	"strings"

	"github.com/hamradiolog-net/spec/v6/internal/codegen"
)

type SecondaryAdministrativeSubdivision string

var _ codegen.CodeGenKey = SecondaryAdministrativeSubdivision("")

// New creates a new SecondaryAdministrativeSubdivision from the provided string.
func New(value string) SecondaryAdministrativeSubdivision {
	return SecondaryAdministrativeSubdivision(strings.ToUpper(value))
}

// String returns the string representation of the SecondaryAdministrativeSubdivision.
func (s SecondaryAdministrativeSubdivision) String() string {
	return string(s)
}

// ADIF enums are case-insensitive.
func (s SecondaryAdministrativeSubdivision) Compare(other SecondaryAdministrativeSubdivision) int {
	return strings.Compare(string(s), string(other))
}
