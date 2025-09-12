package antpath

import (
	"strings"

	"github.com/hamradiolog-net/spec/v6/internal/codegen"
)

// AntPath represents the antenna path used for a QSO
type AntPath string

var _ codegen.CodeGenKey = AntPath("")

// New creates a new AntPath from the provided string.
func New(value string) AntPath {
	return AntPath(strings.ToUpper(value))
}

func (a AntPath) String() string {
	return string(a)
}

// ADIF enums are case-insensitive.
func (a AntPath) Compare(other AntPath) int {
	return strings.Compare(string(a), string(other))
}
