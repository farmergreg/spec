package band

import (
	"strings"

	"github.com/hamradiolog-net/spec/v6/internal/codegen"
)

// Band represents a range of radio frequencies
type Band string

var _ codegen.CodeGenKey = Band("")

// New creates a new Band from the provided string.
func New(value string) Band {
	return Band(strings.ToUpper(value))
}

// String returns the string representation of the Band.
func (b Band) String() string {
	return string(b)
}

// ADIF enums are case-insensitive.
func (b Band) Compare(other Band) int {
	return strings.Compare(string(b), string(other))
}
