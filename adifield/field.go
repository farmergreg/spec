package adifield

import (
	"strings"

	"github.com/farmergreg/spec/v6/internal/codegen"
)

// Field is the ADIF field name in and ADI file.
// By convention in the adifield package, field name constants are uppercase and their values are UPPERCASE.
// n.b. The ADIF specification defines field names as case-insensitive.
type Field string

var _ codegen.CodeGenKey = Field("")

// New creates a new Field from the provided string.
func New(value string) Field {
	return Field(strings.ToUpper(value))
}

// String returns the string representation of the Field.
func (f Field) String() string {
	return string(f)
}

// Compare returns an integer comparing two Field values lexicographically.
// ADIF enums are case-insensitive.
func (f Field) Compare(other Field) int {
	return strings.Compare(strings.ToUpper(string(f)), strings.ToUpper(string(other)))
}

// Equals returns true if this Field equals the other Field.
// ADIF enums are case-insensitive.
func (f Field) Equals(other Field) bool {
	return strings.EqualFold(string(f), string(other))
}
