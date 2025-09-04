package gendef

import (
	"fmt"
	"strconv"

	"github.com/hamradiolog-net/adif-spec/v6/enum/qsodownloadstatus"
)

var _ GenDef = QSODownloadStatusSpec{}

type QSODownloadStatusSpec struct {
	Spec qsodownloadstatus.Spec
}

func (q QSODownloadStatusSpec) ConstName() string {
	return string(q.Spec.Key)
}

func (q QSODownloadStatusSpec) ConstValue() string {
	return strconv.QuoteToASCII(string(q.Spec.Key))
}

func (q QSODownloadStatusSpec) Comments() string {
	return fmt.Sprintf("%s = %s", q.Spec.Key, q.Spec.Description)
}
