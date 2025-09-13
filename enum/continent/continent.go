package continent

import (
	"strings"

	"github.com/hamradiolog-net/spec/v6/internal/codegen"
)

// Continent represents a continent such as EU for Europe
type Continent string

var _ codegen.CodeGenKey = Continent("")

// New creates a new Continent from the provided string.
func New(value string) Continent {
	return Continent(strings.ToUpper(value))
}

// String returns the string representation of the Continent.
func (c Continent) String() string {
	return string(c)
}

// Compare returns an integer comparing two Continent values lexicographically.
// ADIF enums are case-insensitive.
func (c Continent) Compare(other Continent) int {
	return strings.Compare(strings.ToUpper(string(c)), strings.ToUpper(string(other)))
}

// Equals returns true if this Continent equals the other Continent.
// ADIF enums are case-insensitive.
func (c Continent) Equals(other Continent) bool {
	return strings.EqualFold(string(c), string(other))
}
