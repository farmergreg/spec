package spec

import "github.com/hamradiolog-net/adif-spec/src/pkg/enum/propagationmode"

// PropagationModeSpec represents the specification for a single PropagationMode
type PropagationModeSpec struct {
	EnumerationName string          `json:"Enumeration Name"`
	IsImportOnly    AdifSpecBoolean `json:"Import-only,omitempty"`
	Comments        string          `json:"Comments,omitempty"`
	Id              string          `json:"Enumeration"` // Enumeration
	Description     string          `json:"Description"`
}

// PropagationModeSpecMap contains all PropagationModeSpec specifications.
type PropagationModeSpecMap struct {
	Header  []string                                                `json:"Header"`
	Records map[propagationmode.PropagationMode]PropagationModeSpec `json:"Records"`
}
