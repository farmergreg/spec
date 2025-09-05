package qslvia

import "github.com/hamradiolog-net/adif-spec/v6/codegen"

// QSLVia represents the method used to exchange QSL cards
type QSLVia string

var _ codegen.CodeGeneratorEnumValue = QSLVia("")

func (q QSLVia) String() string {
	return string(q)
}
