package mode

import (
	"strings"

	"github.com/hamradiolog-net/spec/v6/internal/codegen"
)

// Mode is the ADIF mode of a radio communication.
type Mode string

var _ codegen.CodeGenKey = Mode("")

// New creates a new Mode from the provided string.
func New(value string) Mode {
	return Mode(strings.ToUpper(value))
}

// String returns the string representation of the Mode.
func (m Mode) String() string {
	return string(m)
}

// ADIF enums are case-insensitive.
func (m Mode) Compare(other Mode) int {
	return strings.Compare(string(m), string(other))
}
