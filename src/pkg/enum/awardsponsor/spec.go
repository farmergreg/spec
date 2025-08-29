package awardsponsor

import (
	"github.com/hamradiolog-net/adif-spec/src/pkg/spectype"
)

// AwardSponsorSpec represents the specification for a single AwardSponsor
type AwardSponsorSpec struct {
	EnumerationName string                   `json:"Enumeration Name"`
	IsImportOnly    spectype.AdifSpecBoolean `json:"Import-only,omitempty"`
	Comments        string                   `json:"Comments,omitempty"`
	Id              string                   `json:"Sponsor"`                 // Sponsor Prefix
	Description     string                   `json:"Sponsoring Organization"` // Sponsoring Organization
}

// SpecMap contains all AwardSponsorSpec specifications.
type SpecMap struct {
	Header  []string                                `json:"Header"`
	Records map[AwardSponsorPrefix]AwardSponsorSpec `json:"Records"`
}
