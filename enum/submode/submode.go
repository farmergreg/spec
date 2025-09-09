package submode

import "github.com/hamradiolog-net/spec/v6/internal/codegen"

// SubMode represents the submode of an ADIF record
type SubMode string

var _ codegen.CodeGenKey = SubMode("")

// String returns the string representation of the SubMode.
func (s SubMode) String() string {
	return string(s)
}
