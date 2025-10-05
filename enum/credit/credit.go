package credit

import (
	"strings"

	"github.com/farmergreg/spec/v6/internal/codegen"
)

// Credit represents an award credit identifier
type Credit string

var _ codegen.CodeGenKey = Credit("")

// New creates a new Credit from the provided string.
func New(value string) Credit {
	return Credit(strings.ToUpper(value))
}

// String returns the string representation of the Credit.
func (c Credit) String() string {
	return string(c)
}

// Compare returns an integer comparing two Credit values lexicographically.
// ADIF enums are case-insensitive.
func (c Credit) Compare(other Credit) int {
	return strings.Compare(strings.ToUpper(string(c)), strings.ToUpper(string(other)))
}

// Equals returns true if this Credit equals the other Credit.
// ADIF enums are case-insensitive.
func (c Credit) Equals(other Credit) bool {
	return strings.EqualFold(string(c), string(other))
}
