package spec

// ContestSpec represents the specification for a single Contest
type ContestSpec struct {
	EnumerationName string          `json:"Enumeration Name"`
	IsImportOnly    AdifSpecBoolean `json:"Import-only,omitempty"`
	Comments        string          `json:"Comments,omitempty"`
	Id              string          `json:"Contest-ID"` // Contest ID
	Description     string          `json:"Description"`
}

// ContestSpecMap contains all ContestSpec specifications.
type ContestSpecMap struct {
	Header  []string               `json:"Header"`
	Records map[string]ContestSpec `json:"Records"`
}
