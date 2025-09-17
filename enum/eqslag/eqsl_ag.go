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
	return EQSLAG(strings.ToLower(value))
}

// String returns the string representation of the EQSLAG.
func (e EQSLAG) String() string {
	return string(e)
}

// Compare returns an integer comparing two EQSLAG values lexicographically.
// ADIF enums are case-insensitive.
func (e EQSLAG) Compare(other EQSLAG) int {
	return strings.Compare(strings.ToLower(string(e)), strings.ToLower(string(other)))
}

// Equals returns true if this EQSLAG equals the other EQSLAG.
// ADIF enums are case-insensitive.
func (e EQSLAG) Equals(other EQSLAG) bool {
	return strings.EqualFold(string(e), string(other))
}
