package gendef

import (
	"fmt"
	"strconv"

	"github.com/hamradiolog-net/adif-spec/v6/enum/credit"
)

var _ GenDef = CreditSpec{}

type CreditSpec struct {
	Spec credit.Spec
}

func (c CreditSpec) ConstName() string {
	return string(c.Spec.Key)
}

func (c CreditSpec) ConstValue() string {
	return strconv.QuoteToASCII(string(c.Spec.Key))
}

func (c CreditSpec) Comments() string {
	return fmt.Sprintf("%-20s = %-15s %-45s %-15s", c.Spec.Key, c.Spec.Sponsor, c.Spec.Award, c.Spec.Facet)
}
