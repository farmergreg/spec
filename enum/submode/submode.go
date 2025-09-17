package submode

import (
	"strings"

	"github.com/hamradiolog-net/spec/v6/internal/codegen"
)

// SubMode represents the submode of an ADIF record
type SubMode string

var _ codegen.CodeGenKey = SubMode("")

// New creates a new SubMode from the provided string.
func New(value string) SubMode {
	return SubMode(strings.ToLower(value))
}

// String returns the string representation of the SubMode.
func (s SubMode) String() string {
	return string(s)
}

// Compare returns an integer comparing two SubMode values lexicographically.
// ADIF enums are case-insensitive.
func (s SubMode) Compare(other SubMode) int {
	return strings.Compare(strings.ToLower(string(s)), strings.ToLower(string(other)))
}

// Equals returns true if this SubMode equals the other SubMode.
// ADIF enums are case-insensitive.
func (s SubMode) Equals(other SubMode) bool {
	return strings.EqualFold(string(s), string(other))
}
