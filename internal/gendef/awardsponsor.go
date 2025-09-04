package gendef

import (
	"fmt"
	"strconv"

	"github.com/hamradiolog-net/adif-spec/v6/enum/awardsponsor"
)

var _ GenDef = AwardSponsorSpec{}

type AwardSponsorSpec struct {
	Spec awardsponsor.Spec
}

func (a AwardSponsorSpec) ConstName() string {
	return string(a.Spec.Key)
}

func (a AwardSponsorSpec) ConstValue() string {
	return strconv.QuoteToASCII(string(a.Spec.Key))
}

func (a AwardSponsorSpec) Comments() string {
	return fmt.Sprintf("%-6s = %s", a.Spec.Key, a.Spec.Description)
}
