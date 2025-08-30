package region

import (
	"fmt"

	"github.com/hamradiolog-net/adif-spec/v2/src/pkg/spectype"
)

// SpecMap contains all Region specifications as defined by the ADIF Workgroup specification exports.
type SpecMapContainer struct {
	// Header  []string         `json:"Header"`
	Records map[Region]Spec `json:"Records"`
}

// Spec represents the specification for a single Region as defined by the ADIF Workgroup specification exports.
type Spec struct {
	// EnumerationName string           `json:"Enumeration Name"`
	IsImportOnly   spectype.Boolean  `json:"Import-only,omitempty"`
	Comments       string            `json:"Comments,omitempty"`
	Key            Region            `json:"Region Entity Code"` // Region Entity Code
	DXCCEntityCode string            `json:"DXCC Entity Code"`   // TODO: get this to deserialize into dxccentitycode.DXCCEntityCode...
	Region         string            `json:"Region"`
	Prefix         string            `json:"Prefix,omitempty"`
	Applicability  string            `json:"Applicability,omitempty"`
	StartDate      spectype.DateTime `json:"Start Date,omitempty"`
	EndDate        spectype.DateTime `json:"End Date,omitempty"`
}

func (s Spec) String() string {
	return fmt.Sprintf("%4s.%-3s = %s", s.Key, s.DXCCEntityCode, s.Region)
}
