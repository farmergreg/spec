package credit

import "github.com/hamradiolog-net/adif-spec/v6/codegen"

// Credit represents an award credit identifier
type Credit string

var _ codegen.CodeGeneratorEnumValue = Credit("")

func (c Credit) String() string {
	return string(c)
}
