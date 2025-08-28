package spec

// AwardSponsorSpec represents the specification for a single AwardSponsor
type AwardSponsorSpec struct {
	BaseEnumerationSpec
	Id          string `json:"Sponsor"`                 // Sponsor Prefix
	Description string `json:"Sponsoring Organization"` // Sponsoring Organization
}

// AwardSponsorSpecMap contains all AwardSponsorSpec specifications.
type AwardSponsorSpecMap struct {
	Header  []string                    `json:"Header"`
	Records map[string]AwardSponsorSpec `json:"Records"`
}
