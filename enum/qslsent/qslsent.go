package qslsent

import (
	"strings"

	"github.com/farmergreg/spec/v6/internal/codegen"
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

// Compare returns an integer comparing two QSLSent values lexicographically.
// ADIF enums are case-insensitive.
func (q QSLSent) Compare(other QSLSent) int {
	return strings.Compare(strings.ToUpper(string(q)), strings.ToUpper(string(other)))
}

// Equals returns true if this QSLSent equals the other QSLSent.
// ADIF enums are case-insensitive.
func (q QSLSent) Equals(other QSLSent) bool {
	return strings.EqualFold(string(q), string(other))
}
