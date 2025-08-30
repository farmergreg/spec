package mode

import (
	"fmt"

	"github.com/hamradiolog-net/adif-spec/v2/src/pkg/enum/submode"
	"github.com/hamradiolog-net/adif-spec/v2/src/pkg/spectype"
)

// SpecMap contains all Mode specifications as defined by the ADIF Workgroup specification exports.
type SpecMapContainer struct {
	// Header  []string         `json:"Header"`
	Records map[Mode]Spec `json:"Records"`
}

// Spec represents the specification for a single Mode as defined by the ADIF Workgroup specification exports.
type Spec struct {
	// EnumerationName string           `json:"Enumeration Name"`
	IsImportOnly spectype.Boolean    `json:"Import-only,omitempty"`
	Comments     string              `json:"Comments,omitempty"`
	Key          Mode                `json:"Mode"` // Mode
	Submodes     submode.SubModeList `json:"Submodes,omitempty"`
	Description  string              `json:"Description,omitempty"` // TODO may want to improve this. e.g. SSB = [ LSB, SSB ]
}

func (s Spec) String() string {
	if s.IsImportOnly {
		return fmt.Sprintf("%-10s = %s", s.Key, s.Submodes)
	} else {
		return fmt.Sprintf("%-22s = %s", s.Key, s.Submodes)
	}
}
