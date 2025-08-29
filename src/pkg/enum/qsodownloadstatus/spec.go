package qsodownloadstatus

import (
	"github.com/hamradiolog-net/adif-spec/src/pkg/spectype"
)

// QSODownloadStatusSpec represents the specification for a single QSODownloadStatus
type QSODownloadStatusSpec struct {
	EnumerationName string                   `json:"Enumeration Name"`
	IsImportOnly    spectype.AdifSpecBoolean `json:"Import-only,omitempty"`
	Comments        string                   `json:"Comments,omitempty"`
	Id              string                   `json:"Status"` // Status
	Description     string                   `json:"Description"`
}

// SpecMap contains all QSODownloadStatusSpec specifications.
type SpecMap struct {
	Header  []string                                    `json:"Header"`
	Records map[QSODownloadStatus]QSODownloadStatusSpec `json:"Records"`
}
