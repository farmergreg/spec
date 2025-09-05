package morsekeytype

import "github.com/hamradiolog-net/adif-spec/v6/codegen"

type MorseKeyType string

var _ codegen.CodeGeneratorEnumValue = MorseKeyType("")

func (m MorseKeyType) String() string {
	return string(m)
}
