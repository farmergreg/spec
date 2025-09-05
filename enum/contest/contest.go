package contest

import "github.com/hamradiolog-net/adif-spec/v6/codegen"

// Contest represents the contest identifier
type Contest string

var _ codegen.CodeGeneratorEnumValue = Contest("")

func (c Contest) String() string {
	return string(c)
}
