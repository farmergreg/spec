package aditype

import "github.com/hamradiolog-net/adif-spec/v6/internal/codegen"

// ADIType represents the ADIF data type of a data field
type ADIType string

var _ codegen.CodeGeneratorEnumValue = ADIType("")

func (t ADIType) String() string {
	return string(t)
}
