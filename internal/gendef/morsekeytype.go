package gendef

import (
	"fmt"
	"strconv"

	"github.com/hamradiolog-net/adif-spec/v6/enum/morsekeytype"
)

var _ GenDef = MorseKeyTypeSpec{}

type MorseKeyTypeSpec struct {
	Spec morsekeytype.Spec
}

func (m MorseKeyTypeSpec) ConstName() string {
	return string(m.Spec.Key)
}

func (m MorseKeyTypeSpec) ConstValue() string {
	return strconv.QuoteToASCII(string(m.Spec.Key))
}

func (m MorseKeyTypeSpec) Comments() string {
	return fmt.Sprintf("%-4s = %s", m.Spec.Key, m.Spec.Description)
}
