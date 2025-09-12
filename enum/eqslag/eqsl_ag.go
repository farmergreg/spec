package eqslag

import (
	"strings"

	"github.com/hamradiolog-net/spec/v6/internal/codegen"
)

// EQSLAG - "Authenticity Guaranteed" by eQSL
type EQSLAG string

var _ codegen.CodeGenKey = EQSLAG("")

// New creates a new EQSLAG from the provided string.
func New(value string) EQSLAG {
	return EQSLAG(strings.ToUpper(value))
}

// String returns the string representation of the EQSLAG.
func (e EQSLAG) String() string {
	return string(e)
}

// ADIF enums are case-insensitive.
func (e EQSLAG) Compare(other EQSLAG) int {
	return strings.Compare(string(e), string(other))
}
