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
	return PropagationMode(strings.ToUpper(value))
}

// String returns the string representation of the PropagationMode.
func (p PropagationMode) String() string {
	return string(p)
}

// ADIF enums are case-insensitive.
func (p PropagationMode) Compare(other PropagationMode) int {
	return strings.Compare(string(p), string(other))
}
