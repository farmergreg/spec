package gendef

import (
	"fmt"
	"strconv"

	"github.com/hamradiolog-net/adif-spec/v6/enum/dxccentitycode"
)

var _ GenDef = DXCCEntityCodeSpec{}

type DXCCEntityCodeSpec struct {
	Spec dxccentitycode.Spec
}

func (d DXCCEntityCodeSpec) ConstName() string {
	return string(d.Spec.Identifier())
}

func (d DXCCEntityCodeSpec) ConstValue() string {
	return strconv.QuoteToASCII(string(d.Spec.Key))
}

func (d DXCCEntityCodeSpec) Comments() string {
	deleted := ""
	if d.Spec.IsDeleted {
		deleted = " (DELETED) "
	}
	return fmt.Sprintf("%s = %s%s", d.Spec.Key, d.Spec.EntityName, deleted)
}
