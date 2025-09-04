package gendef

import (
	"fmt"
	"strconv"

	"github.com/hamradiolog-net/adif-spec/v6/enum/qslsent"
)

var _ GenDef = QSLSentSpec{}

type QSLSentSpec struct {
	Spec qslsent.Spec
}

func (q QSLSentSpec) ConstName() string {
	return string(q.Spec.Key)
}

func (q QSLSentSpec) ConstValue() string {
	return strconv.QuoteToASCII(string(q.Spec.Key))
}

func (q QSLSentSpec) Comments() string {
	return fmt.Sprintf("%s = %s", q.Spec.Key, q.Spec.Description)
}
