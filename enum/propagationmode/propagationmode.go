package propagationmode

import "github.com/hamradiolog-net/adif-spec/v6/internal/codegen"

// PropagationMode represents the propagation mode
type PropagationMode string

var _ codegen.CodeGenKey = PropagationMode("")

func (p PropagationMode) String() string {
	return string(p)
}
