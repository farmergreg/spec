package qslvia

import (
	"github.com/hamradiolog-net/adif-spec/src/pkg/spectype"
)

// Spec represents the specification for a single QSLVia as defined by the ADIF Workgroup specification exports.
type Spec struct {
	EnumerationName string           `json:"Enumeration Name"`
	IsImportOnly    spectype.Boolean `json:"Import-only,omitempty"`
	Comments        string           `json:"Comments,omitempty"`
	Id              string           `json:"Via"` // Via
	Description     string           `json:"Description"`
}

// SpecMap contains all QSLVia specifications as defined by the ADIF Workgroup specification exports.
type SpecMap struct {
	Header  []string        `json:"Header"`
	Records map[QSLVia]Spec `json:"Records"`
}
