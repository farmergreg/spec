package gendef

import (
	"fmt"
	"strconv"

	"github.com/hamradiolog-net/adif-spec/v6/enum/region"
)

var _ GenDef = RegionSpec{}

type RegionSpec struct {
	Spec region.Spec
}

func (r RegionSpec) ConstName() string {
	return string(r.Spec.Code) + "." + strconv.Itoa(int(r.Spec.DXCCEntityCode))
}

func (r RegionSpec) ConstValue() string {
	return strconv.QuoteToASCII(string(r.Spec.Code))
}

func (r RegionSpec) Comments() string {
	return fmt.Sprintf("%4s.%-3s = %-5s %-15s; IMPORTANT: This is NOT the Region Code. It is a lookup key for use with RegionCompositeKeyMap", r.Spec.Code, r.Spec.DXCCEntityCode, r.Spec.Code, r.Spec.Region)
}
