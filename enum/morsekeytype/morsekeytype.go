package morsekeytype

import "github.com/hamradiolog-net/adif-spec/v6/internal/codegen"

type MorseKeyType string

var _ codegen.CodeGenKey = MorseKeyType("")

func (m MorseKeyType) String() string {
	return string(m)
}
