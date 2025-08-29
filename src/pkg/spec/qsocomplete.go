package spec

import "github.com/hamradiolog-net/adif-spec/src/pkg/enum/qsocomplete"

// QSOCompleteSpec represents the specification for a single QSOComplete
type QSOCompleteSpec struct {
	EnumerationName string          `json:"Enumeration Name"`
	IsImportOnly    AdifSpecBoolean `json:"Import-only,omitempty"`
	Comments        string          `json:"Comments,omitempty"`
	Id              string          `json:"Abbreviation"` // Abbreviation
	Description     string          `json:"Meaning"`
}

// QSOCompleteSpecMap contains all QSOCompleteSpec specifications.
type QSOCompleteSpecMap struct {
	Header  []string                                    `json:"Header"`
	Records map[qsocomplete.QSOComplete]QSOCompleteSpec `json:"Records"`
}
