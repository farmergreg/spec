package aditype

import (
	"strings"

	"github.com/hamradiolog-net/spec/v6/internal/codegen"
)

// Type represents the ADIF data type of a data field
type Type string

var _ codegen.CodeGenKey = Type("")

// New creates a new Type from the provided string.
func New(value string) Type {
	return Type(strings.ToUpper(value))
}

// String returns the string representation of the Type.
func (t Type) String() string {
	return string(t)
}

// Compare returns an integer comparing two Type values lexicographically.
// ADIF enums are case-insensitive.
func (t Type) Compare(other Type) int {
	return strings.Compare(strings.ToUpper(string(t)), strings.ToUpper(string(other)))
}

// Equals returns true if this Type equals the other Type.
// ADIF enums are case-insensitive.
func (t Type) Equals(other Type) bool {
	return strings.EqualFold(string(t), string(other))
}
