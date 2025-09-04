package gendef

import (
	"fmt"
	"strconv"

	"github.com/hamradiolog-net/adif-spec/v6/enum/qslrcvd"
)

var _ GenDef = QSLRcvdSpec{}

type QSLRcvdSpec struct {
	Spec qslrcvd.Spec
}

func (q QSLRcvdSpec) ConstName() string {
	return string(q.Spec.Key)
}

func (q QSLRcvdSpec) ConstValue() string {
	return strconv.QuoteToASCII(string(q.Spec.Key))
}

func (q QSLRcvdSpec) Comments() string {
	return fmt.Sprintf("%s = %s", q.Spec.Key, q.Spec.Description)
}
