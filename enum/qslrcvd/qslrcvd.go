package qslrcvd

import (
	"strings"

	"github.com/hamradiolog-net/spec/v6/internal/codegen"
)

// QSLRcvd represents the QSL received status
type QSLRcvd string

var _ codegen.CodeGenKey = QSLRcvd("")

// String returns the string representation of the QSLRcvd.
func (q QSLRcvd) String() string {
	return string(q)
}

// Compare implements the Comparable interface.
// ADIF enums are case-insensitive.
func (t QSLRcvd) Compare(other QSLRcvd) int {
	return strings.Compare(strings.ToUpper(string(t)), strings.ToUpper(other.String()))
}
