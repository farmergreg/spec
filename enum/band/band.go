package band

import "github.com/hamradiolog-net/spec/v6/internal/codegen"

// Band represents a range of radio frequencies
type Band string

var _ codegen.CodeGenKey = Band("")

// String returns the string representation of the Band.
func (b Band) String() string {
	return string(b)
}
