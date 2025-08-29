package spec

// QSLViaSpec represents the specification for a single QSLVia
type QSLViaSpec struct {
	EnumerationName string          `json:"Enumeration Name"`
	IsImportOnly    AdifSpecBoolean `json:"Import-only,omitempty"`
	Comments        string          `json:"Comments,omitempty"`
	Id              string          `json:"Via"` // Via
	Description     string          `json:"Description"`
}

// QSLViaSpecMap contains all QSLViaSpec specifications.
type QSLViaSpecMap struct {
	Header  []string              `json:"Header"`
	Records map[string]QSLViaSpec `json:"Records"`
}
