package credit

import (
	"strings"

	"github.com/hamradiolog-net/spec/v6/internal/codegen"
)

// Credit represents an award credit identifier
type Credit string

var _ codegen.CodeGenKey = Credit("")

// String returns the string representation of the Credit.
func (c Credit) String() string {
	return string(c)
}

// Compare implements the Comparable interface.
// ADIF enums are case-insensitive.
func (t Credit) Compare(other Credit) int {
	return strings.Compare(strings.ToUpper(string(t)), strings.ToUpper(other.String()))
}
