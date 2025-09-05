package antpath

import "github.com/hamradiolog-net/adif-spec/v6/internal/codegen"

// AntPath represents the antenna path used for a QSO
type AntPath string

var _ codegen.CodeGenKey = AntPath("")

func (a AntPath) String() string {
	return string(a)
}
