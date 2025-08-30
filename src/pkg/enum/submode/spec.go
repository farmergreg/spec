package submode

import (
	"fmt"

	"github.com/hamradiolog-net/adif-spec/v2/src/pkg/spectype"
)

// SpecMap contains all Submode specifications as defined by the ADIF Workgroup specification exports.
type SpecMapContainer struct {
	// Header  []string         `json:"Header"`
	Records map[SubMode]Spec `json:"Records"`
}

// Spec represents the specification for a single Submode as defined by the ADIF Workgroup specification exports.
type Spec struct {
	// EnumerationName string           `json:"Enumeration Name"`
	IsImportOnly spectype.Boolean `json:"Import-only,omitempty"`
	Comments     string           `json:"Comments,omitempty"`
	Key          SubMode          `json:"Submode"` // Submode
	Mode         string           `json:"Mode"`
	Description  string           `json:"Description,omitempty"` // TODO would be nice to clean this up Submode Mode Description
}

func (s Spec) String() string {
	return fmt.Sprintf("%-15s = %-15s %s", s.Key, s.Mode, s.Description)
}
