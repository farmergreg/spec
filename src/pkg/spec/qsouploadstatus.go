package spec

import "github.com/hamradiolog-net/adif-spec/src/pkg/enum/qsouploadstatus"

// QSOUploadStatusSpec represents the specification for a single QSOUploadStatus
type QSOUploadStatusSpec struct {
	EnumerationName string          `json:"Enumeration Name"`
	IsImportOnly    AdifSpecBoolean `json:"Import-only,omitempty"`
	Comments        string          `json:"Comments,omitempty"`
	Id              string          `json:"Status"` // Status
	Description     string          `json:"Description"`
}

// QSOUploadStatusSpecMap contains all QSOUploadStatusSpec specifications.
type QSOUploadStatusSpecMap struct {
	Header  []string                                                `json:"Header"`
	Records map[qsouploadstatus.QSOUploadStatus]QSOUploadStatusSpec `json:"Records"`
}
