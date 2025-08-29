package eqslag

import (
	"github.com/hamradiolog-net/adif-spec/src/pkg/spectype"
)

// EQSLAGSpec represents the specification for a single EQSLAG
type EQSLAGSpec struct {
	EnumerationName string                   `json:"Enumeration Name"`
	IsImportOnly    spectype.AdifSpecBoolean `json:"Import-only,omitempty"`
	Comments        string                   `json:"Comments,omitempty"`
	Id              string                   `json:"Status"` // Status
	Description     string                   `json:"Description"`
}

// EQSLAGSpecMap contains all EQSLAGSpec specifications.
type EQSLAGSpecMap struct {
	Header  []string     `json:"Header"`
	Records map[EQSLAG]EQSLAGSpec `json:"Records"`
}
