package qslvia

import (
	"strings"

	"github.com/hamradiolog-net/spec/v6/internal/codegen"
)

// QSLVia represents the method used to exchange QSL cards
type QSLVia string

var _ codegen.CodeGenKey = QSLVia("")

// String returns the string representation of the QSLVia.
func (q QSLVia) String() string {
	return string(q)
}

// Compare implements the Comparable interface.
// ADIF enums are case-insensitive.
func (t QSLVia) Compare(other QSLVia) int {
	return strings.Compare(strings.ToUpper(string(t)), strings.ToUpper(other.String()))
}
