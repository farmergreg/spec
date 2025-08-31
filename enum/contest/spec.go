package contest

import (
	"fmt"

	"github.com/hamradiolog-net/adif-spec/v6/spectype"
)

// Spec represents the specification for a single Contest as defined by the ADIF Workgroup specification exports.
type Spec struct {
	// EnumerationName string           `json:"Enumeration Name"`
	IsImportOnly spectype.Boolean `json:"Import-only,omitempty"`
	// Comments     string           `json:"Comments,omitempty"`
	Key         Contest `json:"Contest-ID"` // Contest ID
	Description string  `json:"Description"`
}

func (s Spec) String() string {
	return fmt.Sprintf("%-20s = %s", s.Key, s.Description)
}

// SpecMap contains all Contest specifications as defined by the ADIF Workgroup specification exports.
type SpecMapContainer struct {
	// Header  []string         `json:"Header"`
	Records map[Contest]Spec `json:"Records"`
}
