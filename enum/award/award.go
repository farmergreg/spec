package award

import (
	"strings"

	"github.com/hamradiolog-net/spec/v6/internal/codegen"
)

// Deprecated: Award represents an ADIF award type
type Award string

var _ codegen.CodeGenKey = Award("")

// Deprecated: String returns the string representation of the Award.
func (a Award) String() string {
	return string(a)
}

// Deprecated: Compare implements the Comparable interface.
// ADIF enums are case-insensitive.
func (t Award) Compare(other Award) int {
	return strings.Compare(strings.ToUpper(string(t)), strings.ToUpper(other.String()))
}
