package qsodownloadstatus

import (
	"github.com/hamradiolog-net/adif-spec/src/pkg/spectype"
)

// Spec represents the specification for a single QSODownloadStatus as defined by the ADIF Workgroup specification exports.
type Spec struct {
	EnumerationName string                   `json:"Enumeration Name"`
	IsImportOnly    spectype.AdifSpecBoolean `json:"Import-only,omitempty"`
	Comments        string                   `json:"Comments,omitempty"`
	Id              string                   `json:"Status"` // Status
	Description     string                   `json:"Description"`
}

// SpecMap contains all QSODownloadStatus specifications as defined by the ADIF Workgroup specification exports.
type SpecMap struct {
	Header  []string                   `json:"Header"`
	Records map[QSODownloadStatus]Spec `json:"Records"`
}
