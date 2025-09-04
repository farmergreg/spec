package gendef

import (
	"fmt"
	"strconv"

	"github.com/hamradiolog-net/adif-spec/v6/enum/mode"
)

var _ GenDef = ModeSpec{}

type ModeSpec struct {
	Spec mode.Spec
}

func (m ModeSpec) ConstName() string {
	return string(m.Spec.Key)
}

func (m ModeSpec) ConstValue() string {
	return strconv.QuoteToASCII(string(m.Spec.Key))
}

func (m ModeSpec) Comments() string {
	if m.Spec.IsImportOnly {
		return fmt.Sprintf("%-10s = %s", m.Spec.Key, m.Spec.Submodes)
	} else {
		return fmt.Sprintf("%-22s = %s", m.Spec.Key, m.Spec.Submodes)
	}
}
