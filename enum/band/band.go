package band

import (
	"strings"

	"github.com/hamradiolog-net/spec/v6/internal/codegen"
)

// Band represents a range of radio frequencies
type Band string

var _ codegen.CodeGenKey = Band("")

// String returns the string representation of the Band.
func (b Band) String() string {
	return string(b)
}

// Compare implements the Comparable interface.
// ADIF enums are case-insensitive.
func (t Band) Compare(other Band) int {
	return strings.Compare(strings.ToUpper(string(t)), strings.ToUpper(other.String()))
}
