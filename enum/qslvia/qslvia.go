package qslvia

import "github.com/hamradiolog-net/spec/v6/internal/codegen"

// QSLVia represents the method used to exchange QSL cards
type QSLVia string

var _ codegen.CodeGenKey = QSLVia("")

// String returns the string representation of the QSLVia.
func (q QSLVia) String() string {
	return string(q)
}
