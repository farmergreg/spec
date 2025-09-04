package gendef

import (
	"fmt"
	"strconv"

	"github.com/hamradiolog-net/adif-spec/v6/enum/arrlsection"
)

var _ GenDef = ARRLSectionSpec{}

type ARRLSectionSpec struct {
	Spec arrlsection.Spec
}

func (a ARRLSectionSpec) ConstName() string {
	return string(a.Spec.Key)
}

func (a ARRLSectionSpec) ConstValue() string {
	return strconv.QuoteToASCII(string(a.Spec.Key))
}

func (a ARRLSectionSpec) Comments() string {
	return fmt.Sprintf("%-4s = %s", a.Spec.Key, a.Spec.Description)
}
