package spec

// AwardSponsorRecord represents a single award sponsor record
type AwardSponsorRecord struct {
	BaseEnumerationSpec
	Id          string `json:"Sponsor"`                 // Sponsor Prefix
	Description string `json:"Sponsoring Organization"` // Sponsoring Organization
}

// AwardSponsorEnumeration represents the complete award sponsor enumeration
type AwardSponsorEnumeration struct {
	Header  []string                      `json:"Header"`
	Records map[string]AwardSponsorRecord `json:"Records"`
}
