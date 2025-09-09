package qslmedium

import "github.com/hamradiolog-net/spec/v6/internal/codegen"

// QSLMedium represents the medium used for QSL exchange
type QSLMedium string

var _ codegen.CodeGenKey = QSLMedium("")

// String returns the string representation of the QSLMedium.
func (q QSLMedium) String() string {
	return string(q)
}
