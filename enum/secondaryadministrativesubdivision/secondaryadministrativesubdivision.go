package secondaryadministrativesubdivision

import (
	"strings"

	"github.com/farmergreg/spec/v6/internal/codegen"
)

type SecondaryAdministrativeSubdivision string

var _ codegen.CodeGenKey = SecondaryAdministrativeSubdivision("")

// New creates a new SecondaryAdministrativeSubdivision from the provided string.
func New(value string) SecondaryAdministrativeSubdivision {
	return SecondaryAdministrativeSubdivision(strings.ToLower(value))
}

// String returns the string representation of the SecondaryAdministrativeSubdivision.
func (s SecondaryAdministrativeSubdivision) String() string {
	return string(s)
}

// Compare returns an integer comparing two SecondaryAdministrativeSubdivision values lexicographically.
// ADIF enums are case-insensitive.
func (s SecondaryAdministrativeSubdivision) Compare(other SecondaryAdministrativeSubdivision) int {
	return strings.Compare(strings.ToLower(string(s)), strings.ToLower(string(other)))
}

// Equals returns true if this SecondaryAdministrativeSubdivision equals the other SecondaryAdministrativeSubdivision.
// ADIF enums are case-insensitive.
func (s SecondaryAdministrativeSubdivision) Equals(other SecondaryAdministrativeSubdivision) bool {
	return strings.EqualFold(string(s), string(other))
}
