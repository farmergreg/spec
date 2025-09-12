package eqslag

import (
	"strings"

	"github.com/hamradiolog-net/spec/v6/internal/codegen"
)

// EQSLAG - "Authenticity Guaranteed" by eQSL
type EQSLAG string

var _ codegen.CodeGenKey = EQSLAG("")

// String returns the string representation of the EQSLAG.
func (e EQSLAG) String() string {
	return string(e)
}

// Compare implements the Comparable interface.
// ADIF enums are case-insensitive.
func (t EQSLAG) Compare(other EQSLAG) int {
	return strings.Compare(strings.ToUpper(string(t)), strings.ToUpper(other.String()))
}
