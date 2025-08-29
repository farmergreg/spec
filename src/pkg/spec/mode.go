package spec

import (
	"github.com/hamradiolog-net/adif-spec/src/pkg/enum/mode"
	"github.com/hamradiolog-net/adif-spec/src/pkg/spectype"
)

// ModeSpec represents the specification for a single Mode
type ModeSpec struct {
	EnumerationName string                   `json:"Enumeration Name"`
	IsImportOnly    spectype.AdifSpecBoolean `json:"Import-only,omitempty"`
	Comments        string                   `json:"Comments,omitempty"`
	Id              string                   `json:"Mode"` // Mode
	Submodes        mode.EnumModeSubModeList `json:"Submodes,omitempty"`
	Description     string                   `json:"Description,omitempty"`
}

// ModeSpecMap contains all ModeSpec specifications.
type ModeSpecMap struct {
	Header  []string               `json:"Header"`
	Records map[mode.Mode]ModeSpec `json:"Records"`
}
