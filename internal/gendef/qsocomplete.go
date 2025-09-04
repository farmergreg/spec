package gendef

import (
	"fmt"
	"strconv"

	"github.com/hamradiolog-net/adif-spec/v6/enum/qsocomplete"
)

var _ GenDef = QSOCompleteSpec{}

type QSOCompleteSpec struct {
	Spec qsocomplete.Spec
}

func (q QSOCompleteSpec) ConstName() string {
	return string(q.Spec.Key)
}

func (q QSOCompleteSpec) ConstValue() string {
	return strconv.QuoteToASCII(string(q.Spec.Key))
}

func (q QSOCompleteSpec) Comments() string {
	return fmt.Sprintf("%-4s = %s", q.Spec.Key, q.Spec.Description)
}
