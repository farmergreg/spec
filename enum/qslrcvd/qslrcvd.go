package qslrcvd

import (
	"strings"

	"github.com/hamradiolog-net/spec/v6/internal/codegen"
)

// QSLRcvd represents the QSL received status
type QSLRcvd string

var _ codegen.CodeGenKey = QSLRcvd("")

// New creates a new QSLRcvd from the provided string.
func New(value string) QSLRcvd {
	return QSLRcvd(strings.ToUpper(value))
}

// String returns the string representation of the QSLRcvd.
func (q QSLRcvd) String() string {
	return string(q)
}

// ADIF enums are case-insensitive.
func (q QSLRcvd) Compare(other QSLRcvd) int {
	return strings.Compare(string(q), string(other))
}
