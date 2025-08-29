package spec

// QSLMediumSpec represents the specification for a single QSLMedium
type QSLMediumSpec struct {
	EnumerationName string          `json:"Enumeration Name"`
	IsImportOnly    AdifSpecBoolean `json:"Import-only,omitempty"`
	Comments        string          `json:"Comments,omitempty"`
	Id              string          `json:"Medium"` // Medium
	Description     string          `json:"Description"`
}

// QSLMediumSpecMap contains all QSLMediumSpec specifications.
type QSLMediumSpecMap struct {
	Header  []string                 `json:"Header"`
	Records map[string]QSLMediumSpec `json:"Records"`
}
