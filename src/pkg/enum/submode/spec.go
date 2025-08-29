package submode

import (
	"github.com/hamradiolog-net/adif-spec/src/pkg/spectype"
)

// SubmodeSpec represents the specification for a single Submode
type SubmodeSpec struct {
	EnumerationName string                   `json:"Enumeration Name"`
	IsImportOnly    spectype.AdifSpecBoolean `json:"Import-only,omitempty"`
	Comments        string                   `json:"Comments,omitempty"`
	Id              string                   `json:"Submode"` // Submode
	Mode            string                   `json:"Mode"`
	Description     string                   `json:"Description,omitempty"`
}

// SpecMap contains all SubmodeSpec specifications.
type SpecMap struct {
	Header  []string                `json:"Header"`
	Records map[SubMode]SubmodeSpec `json:"Records"`
}
