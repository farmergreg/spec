package mode

import (
	"fmt"

	"github.com/hamradiolog-net/adif-spec/v3/enum/submode"
	"github.com/hamradiolog-net/adif-spec/v3/spectype"
)

// SpecMap contains all Mode specifications as defined by the ADIF Workgroup specification exports.
type SpecMapContainer struct {
	// Header  []string         `json:"Header"`
	Records map[Mode]Spec `json:"Records"`
}

// Spec represents the specification for a single Mode as defined by the ADIF Workgroup specification exports.
type Spec struct {
	// EnumerationName string           `json:"Enumeration Name"`
	IsImportOnly spectype.Boolean `json:"Import-only,omitempty"`
	// Comments     string              `json:"Comments,omitempty"`
	Key         Mode                `json:"Mode"` // Mode
	Submodes    submode.SubModeList `json:"Submodes,omitempty"`
	Description string              `json:"Description,omitempty"`
}

func (s Spec) String() string {
	if s.IsImportOnly {
		return fmt.Sprintf("%-10s = %s", s.Key, s.Submodes)
	} else {
		return fmt.Sprintf("%-22s = %s", s.Key, s.Submodes)
	}
}
