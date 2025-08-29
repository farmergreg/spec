package antpath

import (
	"github.com/hamradiolog-net/adif-spec/src/pkg/spectype"
)

// Spec represents the specification for a single AntPath as defined by the ADIF Workgroup specification exports.
type Spec struct {
	EnumerationName string           `json:"Enumeration Name"`
	IsImportOnly    spectype.Boolean `json:"Import-only,omitempty"`
	Comments        string           `json:"Comments,omitempty"`
	Id              string           `json:"Abbreviation"` // Abbreviation
	Description     string           `json:"Meaning"`
}

// SpecMap contains all AntPath specifications as defined by the ADIF Workgroup specification exports.
type SpecMap struct {
	Header  []string         `json:"Header"`
	Records map[AntPath]Spec `json:"Records"`
}
