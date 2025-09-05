package submode

import "github.com/hamradiolog-net/adif-spec/v6/internal/codegen"

// SubMode represents the submode of an ADIF record
type SubMode string

var _ codegen.CodeGeneratorEnumValue = SubMode("")

func (s SubMode) String() string {
	return string(s)
}
