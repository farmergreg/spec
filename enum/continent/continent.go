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

// ADIF enums are case-insensitive.
func (c Continent) Compare(other Continent) int {
	return strings.Compare(string(c), string(other))
}
