package morsekeytype

import "github.com/hamradiolog-net/spec/v6/internal/codegen"

type MorseKeyType string

var _ codegen.CodeGenKey = MorseKeyType("")

// String returns the string representation of the MorseKeyType.
func (m MorseKeyType) String() string {
	return string(m)
}
