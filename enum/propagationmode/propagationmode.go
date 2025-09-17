package propagationmode

import (
	"strings"

	"github.com/hamradiolog-net/spec/v6/internal/codegen"
)

// PropagationMode represents the propagation mode
type PropagationMode string

var _ codegen.CodeGenKey = PropagationMode("")

// New creates a new PropagationMode from the provided string.
func New(value string) PropagationMode {
	return PropagationMode(strings.ToLower(value))
}

// String returns the string representation of the PropagationMode.
func (p PropagationMode) String() string {
	return string(p)
}

// Compare returns an integer comparing two PropagationMode values lexicographically.
// ADIF enums are case-insensitive.
func (p PropagationMode) Compare(other PropagationMode) int {
	return strings.Compare(strings.ToLower(string(p)), strings.ToLower(string(other)))
}

// Equals returns true if this PropagationMode equals the other PropagationMode.
// ADIF enums are case-insensitive.
func (p PropagationMode) Equals(other PropagationMode) bool {
	return strings.EqualFold(string(p), string(other))
}
