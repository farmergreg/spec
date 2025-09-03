package arrlsection

import (
	"fmt"

	"github.com/hamradiolog-net/adif-spec/v6/enum/dxccentitycode"
	"github.com/hamradiolog-net/adif-spec/v6/spectype"
)

// SpecMapContainer contains all ARRLSection specifications as defined by the ADIF Workgroup specification exports.
type SpecMapContainer struct {
	// Header  []string         `json:"Header"`
	Records map[ARRLSection]Spec `json:"Records"`
}

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

// Depreciated: CodeGeneratorMetadata is not part of the stable API and may change without warning in the future even for minor version numbers.
func (s Spec) CodeGeneratorMetadata() string {
	return fmt.Sprintf("%-4s = %s", s.Key, s.Description)
}
