package gendef

import (
	"fmt"
	"strconv"

	"github.com/hamradiolog-net/adif-spec/v6/adifield"
)

var _ GenDef = ADIFieldSpec{}

type ADIFieldSpec struct {
	Spec adifield.Spec
}

func (a ADIFieldSpec) ConstName() string {
	return string(a.Spec.Key)
}

func (a ADIFieldSpec) ConstValue() string {
	return strconv.QuoteToASCII(string(a.Spec.Key))
}

func (a ADIFieldSpec) Comments() string {
	if a.Spec.IsHeaderField {
		return fmt.Sprintf("Header: %s", a.Spec.Description)
	} else {
		return fmt.Sprintf("Record: %s", a.Spec.Description)
	}
}
