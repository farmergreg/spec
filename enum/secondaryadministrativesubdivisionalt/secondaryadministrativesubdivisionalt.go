package secondaryadministrativesubdivisionalt

import (
	"strings"

	"github.com/hamradiolog-net/spec/v6/internal/codegen"
)

type SecondaryAdministrativeSubdivisionAlt string

var _ codegen.CodeGenKey = SecondaryAdministrativeSubdivisionAlt("")

// New creates a new SecondaryAdministrativeSubdivisionAlt from the provided string.
func New(value string) SecondaryAdministrativeSubdivisionAlt {
	return SecondaryAdministrativeSubdivisionAlt(strings.ToLower(value))
}

// String returns the string representation of the SecondaryAdministrativeSubdivisionAlt.
func (s SecondaryAdministrativeSubdivisionAlt) String() string {
	return string(s)
}

// Compare returns an integer comparing two SecondaryAdministrativeSubdivisionAlt values lexicographically.
// ADIF enums are case-insensitive.
func (t SecondaryAdministrativeSubdivisionAlt) Compare(other SecondaryAdministrativeSubdivisionAlt) int {
	return strings.Compare(strings.ToLower(string(t)), strings.ToLower(string(other)))
}

// Equals returns true if this SecondaryAdministrativeSubdivisionAlt equals the other SecondaryAdministrativeSubdivisionAlt.
// ADIF enums are case-insensitive.
func (t SecondaryAdministrativeSubdivisionAlt) Equals(other SecondaryAdministrativeSubdivisionAlt) bool {
	return strings.EqualFold(string(t), string(other))
}
