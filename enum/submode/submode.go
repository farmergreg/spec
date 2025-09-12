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
	return SubMode(strings.ToUpper(value))
}

// String returns the string representation of the SubMode.
func (s SubMode) String() string {
	return string(s)
}

// ADIF enums are case-insensitive.
func (s SubMode) Compare(other SubMode) int {
	return strings.Compare(string(s), string(other))
}
