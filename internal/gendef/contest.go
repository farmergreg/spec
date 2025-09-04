package gendef

import (
	"fmt"
	"strconv"

	"github.com/hamradiolog-net/adif-spec/v6/enum/contest"
)

var _ GenDef = ContestSpec{}

type ContestSpec struct {
	Spec contest.Spec
}

func (c ContestSpec) ConstName() string {
	return string(c.Spec.Key)
}

func (c ContestSpec) ConstValue() string {
	return strconv.QuoteToASCII(string(c.Spec.Key))
}

func (c ContestSpec) Comments() string {
	return fmt.Sprintf("%-20s = %s", c.Spec.Key, c.Spec.Description)
}
