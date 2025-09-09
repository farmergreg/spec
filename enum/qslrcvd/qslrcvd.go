package qslrcvd

import "github.com/hamradiolog-net/spec/v6/internal/codegen"

// QSLRcvd represents the QSL received status
type QSLRcvd string

var _ codegen.CodeGenKey = QSLRcvd("")

// String returns the string representation of the QSLRcvd.
func (q QSLRcvd) String() string {
	return string(q)
}
