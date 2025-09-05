package contest

import "github.com/hamradiolog-net/adif-spec/v6/internal/codegen"

// Contest represents the contest identifier
type Contest string

var _ codegen.CodeGenKey = Contest("")

func (c Contest) String() string {
	return string(c)
}
