package gendef

import (
	"fmt"
	"strconv"

	"github.com/hamradiolog-net/adif-spec/v6/enum/qsouploadstatus"
)

var _ GenDef = QSOUploadStatusSpec{}

type QSOUploadStatusSpec struct {
	Spec qsouploadstatus.Spec
}

func (q QSOUploadStatusSpec) ConstName() string {
	return string(q.Spec.Key)
}

func (q QSOUploadStatusSpec) ConstValue() string {
	return strconv.QuoteToASCII(string(q.Spec.Key))
}

func (q QSOUploadStatusSpec) Comments() string {
	return fmt.Sprintf("%s = %s", q.Spec.Key, q.Spec.Description)
}
