package adifield

import (
	"strings"

	"github.com/hamradiolog-net/spec/v6/internal/codegen"
)

// ADIField is the ADIF field name in and ADI file.
// By convention in the adifield package, field name constants are always UPPERCASE.
// This is a departure from the ADIF specification, which allows field names to include lowercase letters.
type ADIField string

var _ codegen.CodeGenKey = ADIField("")

// New creates a new ADIField from the provided string.
func New(value string) ADIField {
	return ADIField(strings.ToUpper(value))
}

// String returns the string representation of the ADIField.
func (f ADIField) String() string {
	return string(f)
}

// Compare returns an integer comparing two ADIField values lexicographically.
// ADIF enums are case-insensitive.
func (f ADIField) Compare(other ADIField) int {
	return strings.Compare(strings.ToUpper(string(f)), strings.ToUpper(string(other)))
}

// Equals returns true if this ADIField equals the other ADIField.
// ADIF enums are case-insensitive.
func (f ADIField) Equals(other ADIField) bool {
	return strings.EqualFold(string(f), string(other))
}
