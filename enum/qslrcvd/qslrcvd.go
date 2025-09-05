package qslrcvd

import "github.com/hamradiolog-net/adif-spec/v6/internal/codegen"

// QSLRcvd represents the QSL received status
type QSLRcvd string

var _ codegen.CodeGeneratorEnumValue = QSLRcvd("")

func (q QSLRcvd) String() string {
	return string(q)
}
