package mode

import (
	"github.com/hamradiolog-net/adif-spec/src/pkg/enum/submode"
	"github.com/hamradiolog-net/adif-spec/src/pkg/spectype"
)

// Spec represents the specification for a single Mode as defined by the ADIF Workgroup specification exports.
type Spec struct {
	EnumerationName string              `json:"Enumeration Name"`
	IsImportOnly    spectype.Boolean    `json:"Import-only,omitempty"`
	Comments        string              `json:"Comments,omitempty"`
	Id              string              `json:"Mode"` // Mode
	Submodes        submode.SubModeList `json:"Submodes,omitempty"`
	Description     string              `json:"Description,omitempty"`
}

// SpecMap contains all Mode specifications as defined by the ADIF Workgroup specification exports.
type SpecMap struct {
	Header  []string      `json:"Header"`
	Records map[Mode]Spec `json:"Records"`
}
