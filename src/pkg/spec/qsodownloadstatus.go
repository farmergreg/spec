package spec

// QSODownloadStatusSpec represents the specification for a single QSODownloadStatus
type QSODownloadStatusSpec struct {
	EnumerationName string          `json:"Enumeration Name"`
	IsImportOnly    AdifSpecBoolean `json:"Import-only,omitempty"`
	Comments        string          `json:"Comments,omitempty"`
	Id              string          `json:"Status"` // Status
	Description     string          `json:"Description"`
}

// QSODownloadStatusSpecMap contains all QSODownloadStatusSpec specifications.
type QSODownloadStatusSpecMap struct {
	Header  []string                         `json:"Header"`
	Records map[string]QSODownloadStatusSpec `json:"Records"`
}
