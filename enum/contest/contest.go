package contest

import (
	"strings"

	"github.com/hamradiolog-net/spec/v6/internal/codegen"
)

// Contest represents the contest identifier
type Contest string

var _ codegen.CodeGenKey = Contest("")

// New creates a new Contest from the provided string.
func New(value string) Contest {
	return Contest(strings.ToUpper(value))
}

// String returns the string representation of the Contest.
func (c Contest) String() string {
	return string(c)
}

// ADIF enums are case-insensitive.
func (c Contest) Compare(other Contest) int {
	return strings.Compare(string(c), string(other))
}
