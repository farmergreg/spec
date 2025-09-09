package eqslag

import "github.com/hamradiolog-net/spec/v6/internal/codegen"

// EQSLAG - "Authenticity Guaranteed" by eQSL
type EQSLAG string

var _ codegen.CodeGenKey = EQSLAG("")

// String returns the string representation of the EQSLAG.
func (e EQSLAG) String() string {
	return string(e)
}
