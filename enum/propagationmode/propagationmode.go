package propagationmode

import (
	"strings"

	"github.com/hamradiolog-net/spec/v6/internal/codegen"
)

// PropagationMode represents the propagation mode
type PropagationMode string

var _ codegen.CodeGenKey = PropagationMode("")

// String returns the string representation of the PropagationMode.
func (p PropagationMode) String() string {
	return string(p)
}

// Compare implements the Comparable interface.
// ADIF enums are case-insensitive.
func (t PropagationMode) Compare(other PropagationMode) int {
	return strings.Compare(strings.ToUpper(string(t)), strings.ToUpper(other.String()))
}
