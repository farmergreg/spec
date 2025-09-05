package qslmedium

import "github.com/hamradiolog-net/adif-spec/v6/internal/codegen"

// QSLMedium represents the medium used for QSL exchange
type QSLMedium string

var _ codegen.CodeGeneratorEnumValue = QSLMedium("")

func (q QSLMedium) String() string {
	return string(q)
}
