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

// Compare returns an integer comparing two AntPath values lexicographically.
// ADIF enums are case-insensitive.
func (a AntPath) Compare(other AntPath) int {
	return strings.Compare(strings.ToUpper(string(a)), strings.ToUpper(string(other)))
}

// Equals returns true if this AntPath equals the other AntPath.
// ADIF enums are case-insensitive.
func (a AntPath) Equals(other AntPath) bool {
	return strings.EqualFold(string(a), string(other))
}
