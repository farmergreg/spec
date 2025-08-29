package qsouploadstatus

import (
	"github.com/hamradiolog-net/adif-spec/src/pkg/spectype"
)

// Spec represents the specification for a single QSOUploadStatus as defined by the ADIF Workgroup specification exports.
type Spec struct {
	EnumerationName string           `json:"Enumeration Name"`
	IsImportOnly    spectype.Boolean `json:"Import-only,omitempty"`
	Comments        string           `json:"Comments,omitempty"`
	Id              string           `json:"Status"` // Status
	Description     string           `json:"Description"`
}

// SpecMap contains all QSOUploadStatus specifications as defined by the ADIF Workgroup specification exports.
type SpecMap struct {
	Header  []string                 `json:"Header"`
	Records map[QSOUploadStatus]Spec `json:"Records"`
}
