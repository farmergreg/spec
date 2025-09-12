package aditype

import (
	"strings"

	"github.com/hamradiolog-net/spec/v6/internal/codegen"
)

// ADIType represents the ADIF data type of a data field
type ADIType string

var _ codegen.CodeGenKey = ADIType("")

// String returns the string representation of the ADIType.
func (t ADIType) String() string {
	return string(t)
}

// Compare implements the Comparable interface.
// ADIF enums are case-insensitive.
func (t ADIType) Compare(other ADIType) int {
	return strings.Compare(strings.ToUpper(string(t)), strings.ToUpper(other.String()))
}
