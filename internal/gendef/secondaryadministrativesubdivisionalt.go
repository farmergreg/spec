package gendef

import (
	"fmt"
	"strconv"

	"github.com/hamradiolog-net/adif-spec/v6/enum/secondaryadministrativesubdivisionalt"
)

var _ GenDef = SecondaryAdministrativeSubdivisionAltSpec{}

type SecondaryAdministrativeSubdivisionAltSpec struct {
	Spec secondaryadministrativesubdivisionalt.Spec
}

func (s SecondaryAdministrativeSubdivisionAltSpec) ConstName() string {
	return string(s.Spec.Key)
}

func (s SecondaryAdministrativeSubdivisionAltSpec) ConstValue() string {
	return strconv.QuoteToASCII(string(s.Spec.Key))
}

func (s SecondaryAdministrativeSubdivisionAltSpec) Comments() string {
	return fmt.Sprintf("%-50s = DXCC: %s %s", s.Spec.Key, s.Spec.DXCCEntityCode, s.Spec.Region)
}
