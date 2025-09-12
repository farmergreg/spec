package antpath

import (
	"strings"

	"github.com/hamradiolog-net/spec/v6/internal/codegen"
)

// AntPath represents the antenna path used for a QSO
type AntPath string

var _ codegen.CodeGenKey = AntPath("")

func (a AntPath) String() string {
	return string(a)
}

// Compare implements the Comparable interface.
// ADIF enums are case-insensitive.
func (t AntPath) Compare(other AntPath) int {
	return strings.Compare(strings.ToUpper(string(t)), strings.ToUpper(other.String()))
}
