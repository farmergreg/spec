package gendef

import (
	"fmt"
	"strconv"

	"github.com/hamradiolog-net/adif-spec/v6/enum/band"
)

var _ GenDef = BandSpec{}

type BandSpec struct {
	Spec band.Spec
}

func (b BandSpec) ConstName() string {
	return string(b.Spec.Key)
}

func (b BandSpec) ConstValue() string {
	return strconv.QuoteToASCII(string(b.Spec.Key))
}

func (b BandSpec) Comments() string {
	return fmt.Sprintf("%-6s = %12.4f MHz to %12.4f MHz", b.Spec.Key, b.Spec.LowerFreqMHz, b.Spec.UpperFreqMHz)
}
