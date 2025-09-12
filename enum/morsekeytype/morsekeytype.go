package morsekeytype

import (
	"strings"

	"github.com/hamradiolog-net/spec/v6/internal/codegen"
)

type MorseKeyType string

var _ codegen.CodeGenKey = MorseKeyType("")

// String returns the string representation of the MorseKeyType.
func (m MorseKeyType) String() string {
	return string(m)
}

// Compare implements the Comparable interface.
// ADIF enums are case-insensitive.
func (t MorseKeyType) Compare(other MorseKeyType) int {
	return strings.Compare(strings.ToUpper(string(t)), strings.ToUpper(other.String()))
}
