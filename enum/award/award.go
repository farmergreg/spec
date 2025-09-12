package award

import (
	"strings"

	"github.com/hamradiolog-net/spec/v6/internal/codegen"
)

// Deprecated: Award represents an ADIF award type
type Award string

var _ codegen.CodeGenKey = Award("")

// Deprecated: New creates a new Award from the provided string.
func New(value string) Award {
	return Award(strings.ToUpper(value))
}

// Deprecated: String returns the string representation of the Award.
func (a Award) String() string {
	return string(a)
}

// Deprecated: Compare implements the Comparable interface.
// ADIF enums are case-insensitive.
func (t Award) Compare(other Award) int {
	return strings.Compare(string(t), string(other))
}
