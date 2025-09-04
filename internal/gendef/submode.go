package gendef

import (
	"fmt"
	"strconv"

	"github.com/hamradiolog-net/adif-spec/v6/enum/submode"
)

var _ GenDef = SubModeSpec{}

type SubModeSpec struct {
	Spec submode.Spec
}

func (s SubModeSpec) ConstName() string {
	return string(s.Spec.Key)
}

func (s SubModeSpec) ConstValue() string {
	return strconv.QuoteToASCII(string(s.Spec.Key))
}

func (s SubModeSpec) Comments() string {
	return fmt.Sprintf("%-15s = %-15s %s", s.Spec.Key, s.Spec.Mode, s.Spec.Description)
}
