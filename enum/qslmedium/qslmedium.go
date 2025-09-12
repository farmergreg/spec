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

// ADIF enums are case-insensitive.
func (q QSLMedium) Compare(other QSLMedium) int {
	return strings.Compare(string(q), string(other))
}
