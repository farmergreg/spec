package spec

import (
	"github.com/hamradiolog-net/adif-spec/src/pkg/enum/mode"
)

// ModeSpec represents the specification for a single Mode
type ModeSpec struct {
	BaseEnumerationSpec
	Id          string                   `json:"Mode"` // Mode
	Submodes    mode.EnumModeSubModeList `json:"Submodes,omitempty"`
	Description string                   `json:"Description,omitempty"`
}

// ModeSpecMap contains all ModeSpec specifications.
type ModeSpecMap struct {
	Header  []string            `json:"Header"`
	Records map[string]ModeSpec `json:"Records"`
}
