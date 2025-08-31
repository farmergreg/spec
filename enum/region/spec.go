package region

import (
	"fmt"

	"github.com/hamradiolog-net/adif-spec/v6/enum/dxccentitycode"
	"github.com/hamradiolog-net/adif-spec/v6/spectype"
)

// SpecMapContainer contains all Region specifications as defined by the ADIF Workgroup specification exports.
type SpecMapContainer struct {
	// Header  []string         `json:"Header"`
	Records map[RegionCompositeKey]Spec `json:"Records"`
}

// Spec represents the specification for a single Region as defined by the ADIF Workgroup specification exports.
type Spec struct {
	// EnumerationName string           `json:"Enumeration Name"`
	IsImportOnly spectype.Boolean `json:"Import-only,omitempty"`
	// Comments       string            `json:"Comments,omitempty"`
	Code           RegionCode                    `json:"Region Entity Code"` // Region Entity Code
	DXCCEntityCode dxccentitycode.DXCCEntityCode `json:"DXCC Entity Code"`
	Region         string                        `json:"Region"`
	Prefix         string                        `json:"Prefix,omitempty"`
	Applicability  spectype.StringSlice          `json:"Applicability,omitempty"` // TODO custom type
	StartDate      spectype.DateTime             `json:"Start Date,omitempty"`
	EndDate        spectype.DateTime             `json:"End Date,omitempty"`
}

// RegionCode represents a region entity code.
type RegionCode string

func (s Spec) String() string {
	return fmt.Sprintf("%4s.%-3s = %-5s %-15s; IMPORTANT: This is NOT the Region Code. It is a lookup key for use with RegionCompositeKeyMap", s.Code, s.DXCCEntityCode, s.Code, s.Region)
}
