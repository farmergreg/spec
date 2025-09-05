package secondaryadministrativesubdivision

import "github.com/hamradiolog-net/adif-spec/v6/internal/codegen"

type SecondaryAdministrativeSubdivision string

var _ codegen.CodeGeneratorEnumValue = SecondaryAdministrativeSubdivision("")

func (s SecondaryAdministrativeSubdivision) String() string {
	return string(s)
}
