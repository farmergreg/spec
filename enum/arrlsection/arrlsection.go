package arrlsection

import (
	"strings"

	"github.com/hamradiolog-net/spec/v6/internal/codegen"
)

// ARRLSection represents an ARRL section identifier
type ARRLSection string

var _ codegen.CodeGenKey = ARRLSection("")

// New creates a new ARRLSection from the provided string.
func New(value string) ARRLSection {
	return ARRLSection(strings.ToUpper(value))
}

// String returns the string representation of the ARRLSection.
func (a ARRLSection) String() string {
	return string(a)
}

// ADIF enums are case-insensitive.
func (t ARRLSection) Compare(other ARRLSection) int {
	return strings.Compare(string(t), string(other))
}
