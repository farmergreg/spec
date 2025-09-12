package qsocomplete

import (
	"strings"

	"github.com/hamradiolog-net/spec/v6/internal/codegen"
)

// QSOComplete represents the completion status of a QSO
type QSOComplete string

var _ codegen.CodeGenKey = QSOComplete("")

// New creates a new QSOComplete from the provided string.
func New(value string) QSOComplete {
	return QSOComplete(strings.ToUpper(value))
}

// String returns the string representation of the QSOComplete.
func (q QSOComplete) String() string {
	return string(q)
}

// ADIF enums are case-insensitive.
func (q QSOComplete) Compare(other QSOComplete) int {
	return strings.Compare(string(q), string(other))
}
