package awardsponsor

import (
	"fmt"

	"github.com/hamradiolog-net/adif-spec/v2/src/pkg/spectype"
)

// Spec represents the specification for a single AwardSponsor as defined by the ADIF Workgroup specification exports.
type Spec struct {
	// EnumerationName string           `json:"Enumeration Name"`
	IsImportOnly spectype.Boolean `json:"Import-only,omitempty"`
	// Comments     string             `json:"Comments,omitempty"`
	Key         AwardSponsorPrefix `json:"Sponsor"`                 // Sponsor Prefix
	Description string             `json:"Sponsoring Organization"` // Sponsoring Organization
}

func (s Spec) String() string {
	return fmt.Sprintf("%-6s = %s", s.Key, s.Description)
}

// SpecMap contains all AwardSponsor specifications as defined by the ADIF Workgroup specification exports.
type SpecMapContainer struct {
	// Header  []string         `json:"Header"`
	Records map[AwardSponsorPrefix]Spec `json:"Records"`
}
