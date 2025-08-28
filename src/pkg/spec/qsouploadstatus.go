package spec

// QSOUploadStatusSpec represents the specification for a single QSOUploadStatus
type QSOUploadStatusSpec struct {
	BaseEnumerationSpec
	Id          string `json:"Status"` // Status
	Description string `json:"Description"`
}

// QSOUploadStatusSpecMap contains all QSOUploadStatusSpec specifications.
type QSOUploadStatusSpecMap struct {
	Header  []string                       `json:"Header"`
	Records map[string]QSOUploadStatusSpec `json:"Records"`
}
