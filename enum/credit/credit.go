package credit

import (
	"strings"

	"github.com/hamradiolog-net/spec/v6/internal/codegen"
)

// Credit represents an award credit identifier
type Credit string

var _ codegen.CodeGenKey = Credit("")

// New creates a new Credit from the provided string.
func New(value string) Credit {
	return Credit(strings.ToLower(value))
}

// String returns the string representation of the Credit.
func (c Credit) String() string {
	return string(c)
}

// Compare returns an integer comparing two Credit values lexicographically.
// ADIF enums are case-insensitive.
func (c Credit) Compare(other Credit) int {
	return strings.Compare(strings.ToLower(string(c)), strings.ToLower(string(other)))
}

// Equals returns true if this Credit equals the other Credit.
// ADIF enums are case-insensitive.
func (c Credit) Equals(other Credit) bool {
	return strings.EqualFold(string(c), string(other))
}
