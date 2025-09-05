package qsocomplete

import "github.com/hamradiolog-net/adif-spec/v6/internal/codegen"

// QSOComplete represents the completion status of a QSO
type QSOComplete string

var _ codegen.CodeGenKey = QSOComplete("")

func (q QSOComplete) String() string {
	return string(q)
}
