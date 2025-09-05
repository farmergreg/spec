package qslsent

import "github.com/hamradiolog-net/adif-spec/v6/internal/codegen"

// QSLSent represents the QSL sent status
type QSLSent string

var _ codegen.CodeGeneratorEnumValue = QSLSent("")

func (q QSLSent) String() string {
	return string(q)
}
