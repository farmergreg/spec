package secondaryadministrativesubdivisionalt

import "github.com/hamradiolog-net/adif-spec/v6/codegen"

type SecondaryAdministrativeSubdivisionAlt string

var _ codegen.CodeGeneratorEnumValue = SecondaryAdministrativeSubdivisionAlt("")

func (s SecondaryAdministrativeSubdivisionAlt) String() string {
	return string(s)
}
