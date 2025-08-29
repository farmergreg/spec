package qslmedium

import (
	"github.com/hamradiolog-net/adif-spec/src/pkg/spectype"
)

// QSLMediumSpec represents the specification for a single QSLMedium
type QSLMediumSpec struct {
	EnumerationName string                   `json:"Enumeration Name"`
	IsImportOnly    spectype.AdifSpecBoolean `json:"Import-only,omitempty"`
	Comments        string                   `json:"Comments,omitempty"`
	Id              string                   `json:"Medium"` // Medium
	Description     string                   `json:"Description"`
}

// SpecMap contains all QSLMediumSpec specifications.
type SpecMap struct {
	Header  []string                    `json:"Header"`
	Records map[QSLMedium]QSLMediumSpec `json:"Records"`
}
