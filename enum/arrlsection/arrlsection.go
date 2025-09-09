package arrlsection

import "github.com/hamradiolog-net/spec/v6/internal/codegen"

// ARRLSection represents an ARRL section identifier
type ARRLSection string

var _ codegen.CodeGenKey = ARRLSection("")

// String returns the string representation of the ARRLSection.
func (a ARRLSection) String() string {
	return string(a)
}
