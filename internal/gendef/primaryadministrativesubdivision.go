package gendef

import (
	"fmt"
	"strconv"

	"github.com/hamradiolog-net/adif-spec/v6/enum/primaryadministrativesubdivision"
)

var _ GenDef = PrimaryAdministrativeSubdivisionSpec{}

type PrimaryAdministrativeSubdivisionSpec struct {
	Spec primaryadministrativesubdivision.Spec
}

func (p PrimaryAdministrativeSubdivisionSpec) ConstName() string {
	return string(p.Spec.Code)
}

func (p PrimaryAdministrativeSubdivisionSpec) ConstValue() string {
	return strconv.QuoteToASCII(string(p.Spec.Code))
}

func (p PrimaryAdministrativeSubdivisionSpec) Comments() string {
	return fmt.Sprintf("%5s.%-5s = %-5s ( %-5s ); IMPORTANT: This is NOT the Primary Administrative Subdivision Code. It is a lookup key for use with PrimaryAdministrativeSubdivisionCompositeKeyMap", p.Spec.Code, p.Spec.DXCCEntityCode, p.Spec.Code, p.Spec.PrimaryAdminSub)
}
