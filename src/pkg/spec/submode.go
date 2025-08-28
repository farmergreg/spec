package spec

// SubmodeRecord represents a single submode record
type SubmodeRecord struct {
	BaseEnumerationSpec
	Id          string `json:"Submode"` // Submode
	Mode        string `json:"Mode"`
	Description string `json:"Description,omitempty"`
}

// SubmodeEnumeration represents the complete submode enumeration
type SubmodeEnumeration struct {
	Header  []string                 `json:"Header"`
	Records map[string]SubmodeRecord `json:"Records"`
}
