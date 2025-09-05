package mode

import "github.com/hamradiolog-net/adif-spec/v6/internal/codegen"

// Mode is the ADIF mode of a radio communication.
type Mode string

var _ codegen.CodeGeneratorEnumValue = Mode("")

func (m Mode) String() string {
	return string(m)
}
