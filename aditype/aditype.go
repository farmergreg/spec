package aditype

import (
	"strings"

	"github.com/hamradiolog-net/spec/v6/internal/codegen"
)

// ADIType represents the ADIF data type of a data field
type ADIType string

var _ codegen.CodeGenKey = ADIType("")

// New creates a new ADIType from the provided string.
func New(value string) ADIType {
	return ADIType(strings.ToUpper(value))
}

// String returns the string representation of the ADIType.
func (t ADIType) String() string {
	return string(t)
}

// Compare returns an integer comparing two ADIType values lexicographically.
// ADIF enums are case-insensitive.
func (t ADIType) Compare(other ADIType) int {
	return strings.Compare(strings.ToUpper(string(t)), strings.ToUpper(string(other)))
}

// Equals returns true if this ADIType equals the other ADIType.
// ADIF enums are case-insensitive.
func (t ADIType) Equals(other ADIType) bool {
	return strings.EqualFold(string(t), string(other))
}
