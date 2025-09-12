package qslsent

import (
	"strings"

	"github.com/hamradiolog-net/spec/v6/internal/codegen"
)

// QSLSent represents the QSL sent status
type QSLSent string

var _ codegen.CodeGenKey = QSLSent("")

// String returns the string representation of the QSLSent.
func (q QSLSent) String() string {
	return string(q)
}

// Compare implements the Comparable interface.
// ADIF enums are case-insensitive.
func (t QSLSent) Compare(other QSLSent) int {
	return strings.Compare(strings.ToUpper(string(t)), strings.ToUpper(other.String()))
}
