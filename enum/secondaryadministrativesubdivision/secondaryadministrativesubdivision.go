package secondaryadministrativesubdivision

import "github.com/hamradiolog-net/spec/v6/internal/codegen"

type SecondaryAdministrativeSubdivision string

var _ codegen.CodeGenKey = SecondaryAdministrativeSubdivision("")

// String returns the string representation of the SecondaryAdministrativeSubdivision.
func (s SecondaryAdministrativeSubdivision) String() string {
	return string(s)
}
