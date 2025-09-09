package contest

import "github.com/hamradiolog-net/spec/v6/internal/codegen"

// Contest represents the contest identifier
type Contest string

var _ codegen.CodeGenKey = Contest("")

// String returns the string representation of the Contest.
func (c Contest) String() string {
	return string(c)
}
