package contest

import (
	"strings"

	"github.com/farmergreg/spec/v6/internal/codegen"
)

// Contest represents the contest identifier
type Contest string

var _ codegen.CodeGenKey = Contest("")

// New creates a new Contest from the provided string.
func New(value string) Contest {
	return Contest(strings.ToLower(value))
}

// String returns the string representation of the Contest.
func (c Contest) String() string {
	return string(c)
}

// Compare returns an integer comparing two Contest values lexicographically.
// ADIF enums are case-insensitive.
func (c Contest) Compare(other Contest) int {
	return strings.Compare(strings.ToLower(string(c)), strings.ToLower(string(other)))
}

// Equals returns true if this Contest equals the other Contest.
// ADIF enums are case-insensitive.
func (c Contest) Equals(other Contest) bool {
	return strings.EqualFold(string(c), string(other))
}
