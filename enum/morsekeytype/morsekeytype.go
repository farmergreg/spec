package morsekeytype

import (
	"strings"

	"github.com/hamradiolog-net/spec/v6/internal/codegen"
)

type MorseKeyType string

var _ codegen.CodeGenKey = MorseKeyType("")

// New creates a new MorseKeyType from the provided string.
func New(value string) MorseKeyType {
	return MorseKeyType(strings.ToUpper(value))
}

// String returns the string representation of the MorseKeyType.
func (m MorseKeyType) String() string {
	return string(m)
}

// ADIF enums are case-insensitive.
func (m MorseKeyType) Compare(other MorseKeyType) int {
	return strings.Compare(string(m), string(other))
}
