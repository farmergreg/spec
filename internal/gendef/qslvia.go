package gendef

import (
	"fmt"
	"strconv"

	"github.com/hamradiolog-net/adif-spec/v6/enum/qslvia"
)

var _ GenDef = QSLViaSpec{}

type QSLViaSpec struct {
	Spec qslvia.Spec
}

func (q QSLViaSpec) ConstName() string {
	return string(q.Spec.Key)
}

func (q QSLViaSpec) ConstValue() string {
	return strconv.QuoteToASCII(string(q.Spec.Key))
}

func (q QSLViaSpec) Comments() string {
	return fmt.Sprintf("%s = %s", q.Spec.Key, q.Spec.Description)
}
