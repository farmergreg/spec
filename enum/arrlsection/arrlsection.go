package arrlsection

import (
	"strings"

	"github.com/hamradiolog-net/spec/v6/internal/codegen"
)

// ARRLSection represents an ARRL section identifier
type ARRLSection string

var _ codegen.CodeGenKey = ARRLSection("")

// String returns the string representation of the ARRLSection.
func (a ARRLSection) String() string {
	return string(a)
}

// Compare implements the Comparable interface.
// ADIF enums are case-insensitive.
func (t ARRLSection) Compare(other ARRLSection) int {
	return strings.Compare(strings.ToUpper(string(t)), strings.ToUpper(other.String()))
}
