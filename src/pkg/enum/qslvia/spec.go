package qslvia

import (
	"github.com/hamradiolog-net/adif-spec/src/pkg/spectype"
)

// QSLViaSpec represents the specification for a single QSLVia
type QSLViaSpec struct {
	EnumerationName string                   `json:"Enumeration Name"`
	IsImportOnly    spectype.AdifSpecBoolean `json:"Import-only,omitempty"`
	Comments        string                   `json:"Comments,omitempty"`
	Id              string                   `json:"Via"` // Via
	Description     string                   `json:"Description"`
}

// QSLViaSpecMap contains all QSLViaSpec specifications.
type QSLViaSpecMap struct {
	Header  []string     `json:"Header"`
	Records map[QSLVia]QSLViaSpec `json:"Records"`
}
