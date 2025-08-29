package submode

import (
	"github.com/hamradiolog-net/adif-spec/src/pkg/spectype"
)

// Spec represents the specification for a single Submode as defined by the ADIF Workgroup specification exports.
type Spec struct {
	EnumerationName string           `json:"Enumeration Name"`
	IsImportOnly    spectype.Boolean `json:"Import-only,omitempty"`
	Comments        string           `json:"Comments,omitempty"`
	Id              string           `json:"Submode"` // Submode
	Mode            string           `json:"Mode"`
	Description     string           `json:"Description,omitempty"`
}

// SpecMap contains all Submode specifications as defined by the ADIF Workgroup specification exports.
type SpecMap struct {
	Header  []string         `json:"Header"`
	Records map[SubMode]Spec `json:"Records"`
}
