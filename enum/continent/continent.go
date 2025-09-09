package continent

import "github.com/hamradiolog-net/spec/v6/internal/codegen"

// Continent represents a continent such as EU for Europe
type Continent string

var _ codegen.CodeGenKey = Continent("")

// String returns the string representation of the Continent.
func (c Continent) String() string {
	return string(c)
}
