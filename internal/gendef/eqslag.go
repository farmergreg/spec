package gendef

import (
	"fmt"
	"strconv"

	"github.com/hamradiolog-net/adif-spec/v6/enum/eqslag"
)

var _ GenDef = EQSLAGSpec{}

type EQSLAGSpec struct {
	Spec eqslag.Spec
}

func (e EQSLAGSpec) ConstName() string {
	return string(e.Spec.Key)
}

func (e EQSLAGSpec) ConstValue() string {
	return strconv.QuoteToASCII(string(e.Spec.Key))
}

func (e EQSLAGSpec) Comments() string {
	return fmt.Sprintf("%s = %s", e.Spec.Key, e.Spec.Description)
}
