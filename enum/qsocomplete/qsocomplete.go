package qsocomplete

import (
	"strings"

	"github.com/hamradiolog-net/spec/v6/internal/codegen"
)

// QSOComplete represents the completion status of a QSO
type QSOComplete string

var _ codegen.CodeGenKey = QSOComplete("")

// String returns the string representation of the QSOComplete.
func (q QSOComplete) String() string {
	return string(q)
}

// Compare implements the Comparable interface.
// ADIF enums are case-insensitive.
func (t QSOComplete) Compare(other QSOComplete) int {
	return strings.Compare(strings.ToUpper(string(t)), strings.ToUpper(other.String()))
}
