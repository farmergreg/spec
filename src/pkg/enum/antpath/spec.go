package antpath

import (
	"github.com/hamradiolog-net/adif-spec/src/pkg/spectype"
)

// AntPathSpec represents the specification for a single AntPath
type AntPathSpec struct {
	EnumerationName string                   `json:"Enumeration Name"`
	IsImportOnly    spectype.AdifSpecBoolean `json:"Import-only,omitempty"`
	Comments        string                   `json:"Comments,omitempty"`
	Id              string                   `json:"Abbreviation"` // Abbreviation
	Description     string                   `json:"Meaning"`
}

// AntPathEnumerationSpecMap contains all AntPathSpec specifications
type AntPathEnumerationSpecMap struct {
	Header  []string        `json:"Header"`
	Records map[AntPath]AntPathSpec `json:"Records"`
}
