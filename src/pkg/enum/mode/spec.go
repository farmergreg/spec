package mode

import (
	"github.com/hamradiolog-net/adif-spec/src/pkg/spectype"
)

// ModeSpec represents the specification for a single Mode
type ModeSpec struct {
	EnumerationName string                   `json:"Enumeration Name"`
	IsImportOnly    spectype.AdifSpecBoolean `json:"Import-only,omitempty"`
	Comments        string                   `json:"Comments,omitempty"`
	Id              string                   `json:"Mode"` // Mode
	Submodes        EnumModeSubModeList      `json:"Submodes,omitempty"`
	Description     string                   `json:"Description,omitempty"`
}

// SpecMap contains all ModeSpec specifications.
type SpecMap struct {
	Header  []string          `json:"Header"`
	Records map[Mode]ModeSpec `json:"Records"`
}
