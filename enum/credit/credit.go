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
	return Credit(strings.ToUpper(value))
}

// String returns the string representation of the Credit.
func (c Credit) String() string {
	return string(c)
}

// ADIF enums are case-insensitive.
func (c Credit) Compare(other Credit) int {
	return strings.Compare(string(c), string(other))
}
