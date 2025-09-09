package propagationmode

import "github.com/hamradiolog-net/spec/v6/internal/codegen"

// PropagationMode represents the propagation mode
type PropagationMode string

var _ codegen.CodeGenKey = PropagationMode("")

// String returns the string representation of the PropagationMode.
func (p PropagationMode) String() string {
	return string(p)
}
