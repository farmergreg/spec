package gendef

import (
	"fmt"
	"strconv"

	"github.com/hamradiolog-net/adif-spec/v6/enum/continent"
)

var _ GenDef = ContinentSpec{}

type ContinentSpec struct {
	Spec continent.Spec
}

func (c ContinentSpec) ConstName() string {
	return string(c.Spec.Key)
}

func (c ContinentSpec) ConstValue() string {
	return strconv.QuoteToASCII(string(c.Spec.Key))
}

func (c ContinentSpec) Comments() string {
	return fmt.Sprintf("%s = %s", c.Spec.Key, c.Spec.Continent)
}
