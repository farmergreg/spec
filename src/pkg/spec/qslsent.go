package spec

// QSLSentSpec represents the specification for a single QSLSent
type QSLSentSpec struct {
	EnumerationName string          `json:"Enumeration Name"`
	IsImportOnly    AdifSpecBoolean `json:"Import-only,omitempty"`
	Comments        string          `json:"Comments,omitempty"`
	Id              string          `json:"Status"` // Status
	Meaning         string          `json:"Meaning"`
	Description     string          `json:"Description"`
}

// QSLSentSpecMap contains all QSLSentSpec specifications.
type QSLSentSpecMap struct {
	Header  []string               `json:"Header"`
	Records map[string]QSLSentSpec `json:"Records"`
}
