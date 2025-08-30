package secondaryadministrativesubdivisionalt

import (
	"fmt"

	"github.com/hamradiolog-net/adif-spec/v2/src/pkg/spectype"
)

// SpecMap contains all SecondaryAdministrativeSubdivisionAlt specifications as defined by the ADIF Workgroup specification exports.
type SpecMapContainer struct {
	// Header  []string         `json:"Header"`
	Records map[SecondaryAdministrativeSubdivisionAlt]Spec `json:"Records"`
}

// Spec represents the specification for a single SecondaryAdministrativeSubdivisionAlt as defined by the ADIF Workgroup specification exports.
type Spec struct {
	// EnumerationName string           `json:"Enumeration Name"`
	IsImportOnly   spectype.Boolean                      `json:"Import-only,omitempty"`
	Comments       string                                `json:"Comments,omitempty"`
	Key            SecondaryAdministrativeSubdivisionAlt `json:"Code"`             // Code
	DXCCEntityCode string                                `json:"DXCC Entity Code"` // TODO: get this to deserialize into dxccentitycode.DXCCEntityCode...
	Region         string                                `json:"Region"`
	District       string                                `json:"District"`
	IsDeleted      spectype.Boolean                      `json:"Deleted,omitempty"`
}

func (s Spec) String() string {
	return fmt.Sprintf("%-50s = DXCC: %s %s", s.Key, s.DXCCEntityCode, s.Region)
}
