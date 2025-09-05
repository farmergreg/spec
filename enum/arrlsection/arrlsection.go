package arrlsection

import "github.com/hamradiolog-net/adif-spec/v6/internal/codegen"

// ARRLSection represents an ARRL section identifier
type ARRLSection string

var _ codegen.CodeGeneratorEnumValue = ARRLSection("")

func (a ARRLSection) String() string {
	return string(a)
}
