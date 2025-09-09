package award

import "github.com/hamradiolog-net/spec/v6/internal/codegen"

// Deprecated: Award represents an ADIF award type
type Award string

var _ codegen.CodeGenKey = Award("")

// Deprecated: String returns the string representation of the Award.
func (a Award) String() string {
	return string(a)
}
