package secondaryadministrativesubdivisionalt

import (
	"strings"

	"github.com/hamradiolog-net/spec/v6/internal/codegen"
)

type SecondaryAdministrativeSubdivisionAlt string

var _ codegen.CodeGenKey = SecondaryAdministrativeSubdivisionAlt("")

// New creates a new SecondaryAdministrativeSubdivisionAlt from the provided string.
func New(value string) SecondaryAdministrativeSubdivisionAlt {
	return SecondaryAdministrativeSubdivisionAlt(strings.ToUpper(value))
}

// String returns the string representation of the SecondaryAdministrativeSubdivisionAlt.
func (s SecondaryAdministrativeSubdivisionAlt) String() string {
	return string(s)
}

// ADIF enums are case-insensitive.
func (t SecondaryAdministrativeSubdivisionAlt) Compare(other SecondaryAdministrativeSubdivisionAlt) int {
	return strings.Compare(string(t), string(other))
}
