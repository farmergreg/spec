package awardsponsor

import (
	"github.com/hamradiolog-net/adif-spec/src/pkg/spectype"
)

// Spec represents the specification for a single AwardSponsor as defined by the ADIF Workgroup specification exports.
type Spec struct {
	EnumerationName string           `json:"Enumeration Name"`
	IsImportOnly    spectype.Boolean `json:"Import-only,omitempty"`
	Comments        string           `json:"Comments,omitempty"`
	Id              string           `json:"Sponsor"`                 // Sponsor Prefix
	Description     string           `json:"Sponsoring Organization"` // Sponsoring Organization
}

// SpecMap contains all AwardSponsor specifications as defined by the ADIF Workgroup specification exports.
type SpecMap struct {
	Header  []string                    `json:"Header"`
	Records map[AwardSponsorPrefix]Spec `json:"Records"`
}
