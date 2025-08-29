package qsocomplete

import (
	"github.com/hamradiolog-net/adif-spec/src/pkg/spectype"
)

// QSOCompleteSpec represents the specification for a single QSOComplete
type QSOCompleteSpec struct {
	EnumerationName string                   `json:"Enumeration Name"`
	IsImportOnly    spectype.AdifSpecBoolean `json:"Import-only,omitempty"`
	Comments        string                   `json:"Comments,omitempty"`
	Id              string                   `json:"Abbreviation"` // Abbreviation
	Description     string                   `json:"Meaning"`
}

// SpecMap contains all QSOCompleteSpec specifications.
type SpecMap struct {
	Header  []string                        `json:"Header"`
	Records map[QSOComplete]QSOCompleteSpec `json:"Records"`
}
