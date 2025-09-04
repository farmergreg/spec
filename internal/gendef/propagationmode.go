package gendef

import (
	"fmt"
	"strconv"

	"github.com/hamradiolog-net/adif-spec/v6/enum/propagationmode"
)

var _ GenDef = PropagationModeSpec{}

type PropagationModeSpec struct {
	Spec propagationmode.Spec
}

func (p PropagationModeSpec) ConstName() string {
	return string(p.Spec.Key)
}

func (p PropagationModeSpec) ConstValue() string {
	return strconv.QuoteToASCII(string(p.Spec.Key))
}

func (p PropagationModeSpec) Comments() string {
	return fmt.Sprintf("%-10s = %s", p.Spec.Key, p.Spec.Description)
}
