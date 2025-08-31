package region

import (
	"fmt"

	"github.com/hamradiolog-net/adif-spec/v8/enum/dxccentitycode"
	"github.com/hamradiolog-net/adif-spec/v8/spectype"
)

// SpecMap contains all Region specifications as defined by the ADIF Workgroup specification exports.
type SpecMapContainer struct {
	// Header  []string         `json:"Header"`
	Records map[RegionCompositeKey]Spec `json:"Records"`
}

// Spec represents the specification for a single Region as defined by the ADIF Workgroup specification exports.
type Spec struct {
	// EnumerationName string           `json:"Enumeration Name"`
	IsImportOnly spectype.Boolean `json:"Import-only,omitempty"`
	// Comments       string            `json:"Comments,omitempty"`
	Key            RegionCompositeKey            `json:"Region Entity Code"` // Region Entity Code
	DXCCEntityCode dxccentitycode.DXCCEntityCode `json:"DXCC Entity Code"`
	Region         Region                        `json:"Region"`
	Prefix         string                        `json:"Prefix,omitempty"`
	Applicability  spectype.StringSlice          `json:"Applicability,omitempty"` // TODO custom type
	StartDate      spectype.DateTime             `json:"Start Date,omitempty"`
	EndDate        spectype.DateTime             `json:"End Date,omitempty"`
}

// Region represents a region entity code.
type Region string

func (s Spec) String() string {
	return fmt.Sprintf("%4s.%-3s = %-5s %-15s; IMPORTANT: This is NOT the Region Code. It is a lookup key for use with RegionCompositeKeyMap", s.Key, s.DXCCEntityCode, s.Key, s.Region)
}
