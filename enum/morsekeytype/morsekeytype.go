package morsekeytype

import (
	"strings"

	"github.com/hamradiolog-net/spec/v6/internal/codegen"
)

type MorseKeyType string

var _ codegen.CodeGenKey = MorseKeyType("")

// New creates a new MorseKeyType from the provided string.
func New(value string) MorseKeyType {
	return MorseKeyType(strings.ToLower(value))
}

// String returns the string representation of the MorseKeyType.
func (m MorseKeyType) String() string {
	return string(m)
}

// Compare returns an integer comparing two MorseKeyType values lexicographically.
// ADIF enums are case-insensitive.
func (m MorseKeyType) Compare(other MorseKeyType) int {
	return strings.Compare(strings.ToLower(string(m)), strings.ToLower(string(other)))
}

// Equals returns true if this MorseKeyType equals the other MorseKeyType.
// ADIF enums are case-insensitive.
func (m MorseKeyType) Equals(other MorseKeyType) bool {
	return strings.EqualFold(string(m), string(other))
}
