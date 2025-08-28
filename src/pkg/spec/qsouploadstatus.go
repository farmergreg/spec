package spec

// QSOUploadStatusRecord represents a single QSO upload status record
type QSOUploadStatusRecord struct {
	BaseEnumerationRecord
	Id          string `json:"Status"` // Status
	Description string `json:"Description"`
}

// QSOUploadStatusEnumeration represents the complete QSO upload status enumeration
type QSOUploadStatusEnumeration struct {
	Header  []string                         `json:"Header"`
	Records map[string]QSOUploadStatusRecord `json:"Records"`
}
