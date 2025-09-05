package qsocomplete

import "github.com/hamradiolog-net/adif-spec/v6/codegen"

// QSOComplete represents the completion status of a QSO
type QSOComplete string

var _ codegen.CodeGeneratorEnumValue = QSOComplete("")

func (q QSOComplete) String() string {
	return string(q)
}
