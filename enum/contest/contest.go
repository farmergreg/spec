package contest

import (
	"strings"

	"github.com/hamradiolog-net/spec/v6/internal/codegen"
)

// Contest represents the contest identifier
type Contest string

var _ codegen.CodeGenKey = Contest("")

// String returns the string representation of the Contest.
func (c Contest) String() string {
	return string(c)
}

// Compare implements the Comparable interface.
// ADIF enums are case-insensitive.
func (t Contest) Compare(other Contest) int {
	return strings.Compare(strings.ToUpper(string(t)), strings.ToUpper(other.String()))
}
