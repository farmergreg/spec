package arrlsection

import (
	"strings"

	"github.com/farmergreg/spec/v6/internal/codegen"
)

// ARRLSection represents an ARRL section identifier
type ARRLSection string

var _ codegen.CodeGenKey = ARRLSection("")

// New creates a new ARRLSection from the provided string.
func New(value string) ARRLSection {
	return ARRLSection(strings.ToLower(value))
}

// String returns the string representation of the ARRLSection.
func (a ARRLSection) String() string {
	return string(a)
}

// Compare returns an integer comparing two ARRLSection values lexicographically.
// ADIF enums are case-insensitive.
func (t ARRLSection) Compare(other ARRLSection) int {
	return strings.Compare(strings.ToLower(string(t)), strings.ToLower(string(other)))
}

// Equals returns true if this ARRLSection equals the other ARRLSection.
// ADIF enums are case-insensitive.
func (t ARRLSection) Equals(other ARRLSection) bool {
	return strings.EqualFold(string(t), string(other))
}
