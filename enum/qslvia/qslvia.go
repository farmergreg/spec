package qslvia

import (
	"strings"

	"github.com/hamradiolog-net/spec/v6/internal/codegen"
)

// QSLVia represents the method used to exchange QSL cards
type QSLVia string

var _ codegen.CodeGenKey = QSLVia("")

// New creates a new QSLVia from the provided string.
func New(value string) QSLVia {
	return QSLVia(strings.ToUpper(value))
}

// String returns the string representation of the QSLVia.
func (q QSLVia) String() string {
	return string(q)
}

// Compare returns an integer comparing two QSLVia values lexicographically.
// ADIF enums are case-insensitive.
func (q QSLVia) Compare(other QSLVia) int {
	return strings.Compare(strings.ToUpper(string(q)), strings.ToUpper(string(other)))
}

// Equals returns true if this QSLVia equals the other QSLVia.
// ADIF enums are case-insensitive.
func (q QSLVia) Equals(other QSLVia) bool {
	return strings.EqualFold(string(q), string(other))
}
