package submode

import (
	"strings"

	"github.com/hamradiolog-net/spec/v6/internal/codegen"
)

// SubMode represents the submode of an ADIF record
type SubMode string

var _ codegen.CodeGenKey = SubMode("")

// String returns the string representation of the SubMode.
func (s SubMode) String() string {
	return string(s)
}

// Compare implements the Comparable interface.
// ADIF enums are case-insensitive.
func (t SubMode) Compare(other SubMode) int {
	return strings.Compare(strings.ToUpper(string(t)), strings.ToUpper(other.String()))
}
