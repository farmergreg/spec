package award

import (
	"strings"

	"github.com/farmergreg/spec/v6/internal/codegen"
)

// Deprecated: Award represents an ADIF award type
type Award string

var _ codegen.CodeGenKey = Award("")

// Deprecated: New creates a new Award from the provided string.
func New(value string) Award {
	return Award(strings.ToLower(value))
}

// Deprecated: String returns the string representation of the Award.
func (a Award) String() string {
	return string(a)
}

// Deprecated: Compare returns an integer comparing two Award values lexicographically.
// ADIF enums are case-insensitive.
func (t Award) Compare(other Award) int {
	return strings.Compare(strings.ToLower(string(t)), strings.ToLower(string(other)))
}

// Deprecated: Equals returns true if this Award equals the other Award.
// ADIF enums are case-insensitive.
func (t Award) Equals(other Award) bool {
	return strings.EqualFold(string(t), string(other))
}
