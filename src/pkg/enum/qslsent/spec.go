package qslsent

import (
	"github.com/hamradiolog-net/adif-spec/src/pkg/spectype"
)

// QSLSentSpec represents the specification for a single QSLSent
type QSLSentSpec struct {
	EnumerationName string                   `json:"Enumeration Name"`
	IsImportOnly    spectype.AdifSpecBoolean `json:"Import-only,omitempty"`
	Comments        string                   `json:"Comments,omitempty"`
	Id              string                   `json:"Status"` // Status
	Meaning         string                   `json:"Meaning"`
	Description     string                   `json:"Description"`
}

// SpecMap contains all QSLSentSpec specifications.
type SpecMap struct {
	Header  []string                `json:"Header"`
	Records map[QSLSent]QSLSentSpec `json:"Records"`
}
