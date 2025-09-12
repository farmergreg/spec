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

// ADIF enums are case-insensitive.
func (t ADIType) Compare(other ADIType) int {
	return strings.Compare(string(t), string(other))
}
