package gendef

import (
	"strconv"

	"github.com/hamradiolog-net/adif-spec/v6/enum/award"
)

var _ GenDef = AwardSpec{}

type AwardSpec struct {
	Spec award.Spec
}

func (a AwardSpec) ConstName() string {
	return string(a.Spec.Key)
}

func (a AwardSpec) ConstValue() string {
	return strconv.QuoteToASCII(string(a.Spec.Key))
}

func (a AwardSpec) Comments() string {
	return string(a.Spec.Key)
}
