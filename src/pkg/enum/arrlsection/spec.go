package arrlsection

import (
	"fmt"

	"github.com/hamradiolog-net/adif-spec/v2/src/pkg/enum/dxccentitycode"
	"github.com/hamradiolog-net/adif-spec/v2/src/pkg/spectype"
)

// Spec represents the specification for a single ARRLSection as defined by the ADIF Workgroup specification exports.
type Spec struct {
	// EnumerationName string           `json:"Enumeration Name"`
	IsImportOnly   spectype.Boolean                  `json:"Import-only,omitempty"`
	Comments       string                            `json:"Comments,omitempty"`
	Key            ARRLSection                       `json:"Abbreviation"` // Abbreviation
	Description    string                            `json:"Section Name"` // Section Name
	DXCCEntityCode dxccentitycode.DXCCEntityCodeList `json:"DXCC Entity Code"`
	FromDate       spectype.DateTime                 `json:"From Date,omitempty"`
	DeletedDate    spectype.DateTime                 `json:"Deleted Date,omitempty"`
}

func (s Spec) String() string {
	return fmt.Sprintf("%-4s = %s", s.Key, s.Description)
}

// SpecMap contains all ARRLSection specifications as defined by the ADIF Workgroup specification exports.
type SpecMapContainer struct {
	// Header  []string         `json:"Header"`
	Records map[ARRLSection]Spec `json:"Records"`
}
