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
	return Band(strings.ToLower(value))
}

// String returns the string representation of the Band.
func (b Band) String() string {
	return string(b)
}

// Compare returns an integer comparing two Band values lexicographically.
// ADIF enums are case-insensitive.
func (b Band) Compare(other Band) int {
	return strings.Compare(strings.ToLower(string(b)), strings.ToLower(string(other)))
}

// Equals returns true if this Band equals the other Band.
// ADIF enums are case-insensitive.
func (b Band) Equals(other Band) bool {
	return strings.EqualFold(string(b), string(other))
}
