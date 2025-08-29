package spec

import "github.com/hamradiolog-net/adif-spec/src/pkg/enum/submode"

// SubmodeSpec represents the specification for a single Submode
type SubmodeSpec struct {
	EnumerationName string          `json:"Enumeration Name"`
	IsImportOnly    AdifSpecBoolean `json:"Import-only,omitempty"`
	Comments        string          `json:"Comments,omitempty"`
	Id              string          `json:"Submode"` // Submode
	Mode            string          `json:"Mode"`
	Description     string          `json:"Description,omitempty"`
}

// SubmodeSpecMap contains all SubmodeSpec specifications.
type SubmodeSpecMap struct {
	Header  []string                        `json:"Header"`
	Records map[submode.SubMode]SubmodeSpec `json:"Records"`
}
