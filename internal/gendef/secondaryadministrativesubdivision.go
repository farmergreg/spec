package gendef

import (
	"fmt"
	"strconv"

	"github.com/hamradiolog-net/adif-spec/v6/enum/secondaryadministrativesubdivision"
)

var _ GenDef = SecondaryAdministrativeSubdivisionSpec{}

type SecondaryAdministrativeSubdivisionSpec struct {
	Spec secondaryadministrativesubdivision.Spec
}

func (s SecondaryAdministrativeSubdivisionSpec) ConstName() string {
	return string(s.Spec.Key)
}

func (s SecondaryAdministrativeSubdivisionSpec) ConstValue() string {
	return strconv.QuoteToASCII(string(s.Spec.Key))
}

func (s SecondaryAdministrativeSubdivisionSpec) Comments() string {
	return fmt.Sprintf("%-35s = DXCC %s: %s", s.Spec.Key, s.Spec.DXCCEntityCode, s.Spec.SecondaryAdminSub)
}
