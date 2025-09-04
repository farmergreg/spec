package gendef

import (
	"fmt"
	"strconv"

	"github.com/hamradiolog-net/adif-spec/v6/enum/qslmedium"
)

var _ GenDef = QSLMediumSpec{}

type QSLMediumSpec struct {
	Spec qslmedium.Spec
}

func (q QSLMediumSpec) ConstName() string {
	return string(q.Spec.Key)
}

func (q QSLMediumSpec) ConstValue() string {
	return strconv.QuoteToASCII(string(q.Spec.Key))
}

func (q QSLMediumSpec) Comments() string {
	return fmt.Sprintf("%s = %s", q.Spec.Key, q.Spec.Description)
}
