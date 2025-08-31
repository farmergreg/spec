package propagationmode

import (
	"fmt"

	"github.com/hamradiolog-net/adif-spec/v6/spectype"
)

// SpecMapContainer contains all PropagationMode specifications as defined by the ADIF Workgroup specification exports.
type SpecMapContainer struct {
	// Header  []string         `json:"Header"`
	Records map[PropagationMode]Spec `json:"Records"`
}

// Spec represents the specification for a single PropagationMode as defined by the ADIF Workgroup specification exports.
type Spec struct {
	// EnumerationName string           `json:"Enumeration Name"`
	IsImportOnly spectype.Boolean `json:"Import-only,omitempty"`
	// Comments     string           `json:"Comments,omitempty"`
	Key         PropagationMode `json:"Enumeration"` // Enumeration
	Description string          `json:"Description"`
}

func (s Spec) String() string {
	return fmt.Sprintf("%-10s = %s", s.Key, s.Description)
}
