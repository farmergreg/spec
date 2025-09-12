package qslsent

import (
	"strings"

	"github.com/hamradiolog-net/spec/v6/internal/codegen"
)

// QSLSent represents the QSL sent status
type QSLSent string

var _ codegen.CodeGenKey = QSLSent("")

// New creates a new QSLSent from the provided string.
func New(value string) QSLSent {
	return QSLSent(strings.ToUpper(value))
}

// String returns the string representation of the QSLSent.
func (q QSLSent) String() string {
	return string(q)
}

// ADIF enums are case-insensitive.
func (q QSLSent) Compare(other QSLSent) int {
	return strings.Compare(string(q), string(other))
}
