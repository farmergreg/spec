package gendef

import (
	"strconv"

	"github.com/hamradiolog-net/adif-spec/v6/aditype"
)

var _ GenDef = ADITypeSpec{}

type ADITypeSpec struct {
	Spec aditype.Spec
}

func (a ADITypeSpec) ConstName() string {
	return string(a.Spec.Key)
}

func (a ADITypeSpec) ConstValue() string {
	return strconv.QuoteToASCII(string(a.Spec.Key))
}

func (a ADITypeSpec) Comments() string {
	return a.Spec.Description
}
