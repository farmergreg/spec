package gendef

import (
	"fmt"
	"strconv"

	"github.com/hamradiolog-net/adif-spec/v6/enum/antpath"
)

var _ GenDef = AntPathSpec{}

type AntPathSpec struct {
	Spec antpath.Spec
}

func (a AntPathSpec) ConstName() string {
	return string(a.Spec.Key)
}

func (a AntPathSpec) ConstValue() string {
	return strconv.QuoteToASCII(string(a.Spec.Key))
}

func (a AntPathSpec) Comments() string {
	return fmt.Sprintf("%s = %s", a.Spec.Key, a.Spec.Description)
}
