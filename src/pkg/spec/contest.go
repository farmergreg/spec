package spec

// ContestIDRecord represents a single contest ID record
type ContestIDRecord struct {
	BaseEnumerationRecord
	Id          string `json:"Contest-ID"` // Contest ID
	Description string `json:"Description"`
}

// ContestIDEnumeration represents the complete contest ID enumeration
type ContestIDEnumeration struct {
	Header  []string                   `json:"Header"`
	Records map[string]ContestIDRecord `json:"Records"`
}
