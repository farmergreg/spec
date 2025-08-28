package spec

import (
	"github.com/hamradiolog-net/adif-spec/src/pkg/enum/mode"
)

// ModeRecord represents a single mode record
type ModeRecord struct {
	BaseEnumerationRecord
	Id          string                   `json:"Mode"` // Mode
	Submodes    mode.EnumModeSubModeList `json:"Submodes,omitempty"`
	Description string                   `json:"Description,omitempty"`
}

// ModeEnumeration represents the complete mode enumeration
type ModeEnumeration struct {
	Header  []string              `json:"Header"`
	Records map[string]ModeRecord `json:"Records"`
}
