package spec

// QSODownloadStatusRecord represents a single QSO download status record
type QSODownloadStatusRecord struct {
	BaseEnumerationSpec
	Id          string `json:"Status"` // Status
	Description string `json:"Description"`
}

// QSODownloadStatusEnumeration represents the complete QSO download status enumeration
type QSODownloadStatusEnumeration struct {
	Header  []string                           `json:"Header"`
	Records map[string]QSODownloadStatusRecord `json:"Records"`
}
