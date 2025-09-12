package qslmedium

import (
	"strings"

	"github.com/hamradiolog-net/spec/v6/internal/codegen"
)

// QSLMedium represents the medium used for QSL exchange
type QSLMedium string

var _ codegen.CodeGenKey = QSLMedium("")

// String returns the string representation of the QSLMedium.
func (q QSLMedium) String() string {
	return string(q)
}

// Compare implements the Comparable interface.
// ADIF enums are case-insensitive.
func (t QSLMedium) Compare(other QSLMedium) int {
	return strings.Compare(strings.ToUpper(string(t)), strings.ToUpper(other.String()))
}
