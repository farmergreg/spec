package award

import (
	"github.com/hamradiolog-net/adif-spec/v6/spectype"
)

// SpecMapContainer contains all Award specifications as defined by the ADIF Workgroup specification exports.
type SpecMapContainer struct {
	// Header  []string         `json:"Header"`
	Records map[Award]Spec `json:"Records"`
}

// Spec represents the specification for a single Award as defined by the ADIF Workgroup specification exports.
type Spec struct {
	// EnumerationName string           `json:"Enumeration Name"`
	IsImportOnly spectype.Boolean `json:"Import-only,omitempty"`
	// Comments     string           `json:"Comments,omitempty"`
	Key Award `json:"Award"` // Award
}

// Depreciated: CodeGeneratorMetadata is not part of the stable API and may change without warning in the future even for minor version numbers.
func (s Spec) CodeGeneratorMetadata() string {
	return string(s.Key)
}
