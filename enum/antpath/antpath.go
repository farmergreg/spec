package antpath

import "github.com/hamradiolog-net/adif-spec/v6/codegen"

// AntPath represents the antenna path used for a QSO
type AntPath string

var _ codegen.CodeGeneratorEnumValue = AntPath("")

func (a AntPath) String() string {
	return string(a)
}
