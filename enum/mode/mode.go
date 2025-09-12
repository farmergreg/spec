package mode

import (
	"strings"

	"github.com/hamradiolog-net/spec/v6/internal/codegen"
)

// Mode is the ADIF mode of a radio communication.
type Mode string

var _ codegen.CodeGenKey = Mode("")

// String returns the string representation of the Mode.
func (m Mode) String() string {
	return string(m)
}

// Compare implements the Comparable interface.
// ADIF enums are case-insensitive.
func (t Mode) Compare(other Mode) int {
	return strings.Compare(strings.ToUpper(string(t)), strings.ToUpper(other.String()))
}
