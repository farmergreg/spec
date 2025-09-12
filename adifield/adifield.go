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

// String returns the string representation of the ADIField.
func (f ADIField) String() string {
	return string(f)
}

// Compare implements the Comparable interface.
// ADIF enums are case-insensitive.
func (t ADIField) Compare(other ADIField) int {
	return strings.Compare(strings.ToUpper(string(t)), strings.ToUpper(other.String()))
}
