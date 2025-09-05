package eqslag

import "github.com/hamradiolog-net/adif-spec/v6/internal/codegen"

// EQSLAG - "Authenticity Guaranteed" by eQSL
type EQSLAG string

var _ codegen.CodeGenKey = EQSLAG("")

func (e EQSLAG) String() string {
	return string(e)
}
