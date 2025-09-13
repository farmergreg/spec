package qslmedium

import (
	"strings"

	"github.com/hamradiolog-net/spec/v6/internal/codegen"
)

// QSLMedium represents the medium used for QSL exchange
type QSLMedium string

var _ codegen.CodeGenKey = QSLMedium("")

// New creates a new QSLMedium from the provided string.
func New(value string) QSLMedium {
	return QSLMedium(strings.ToUpper(value))
}

// String returns the string representation of the QSLMedium.
func (q QSLMedium) String() string {
	return string(q)
}

// Compare returns an integer comparing two QSLMedium values lexicographically.
// ADIF enums are case-insensitive.
func (q QSLMedium) Compare(other QSLMedium) int {
	return strings.Compare(strings.ToUpper(string(q)), strings.ToUpper(string(other)))
}

// Equals returns true if this QSLMedium equals the other QSLMedium.
// ADIF enums are case-insensitive.
func (q QSLMedium) Equals(other QSLMedium) bool {
	return strings.EqualFold(string(q), string(other))
}
