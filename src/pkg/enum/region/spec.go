package region

import (
	"github.com/hamradiolog-net/adif-spec/src/pkg/spectype"
)

// Spec represents the specification for a single Region as defined by the ADIF Workgroup specification exports.
type Spec struct {
	EnumerationName string                    `json:"Enumeration Name"`
	IsImportOnly    spectype.AdifSpecBoolean  `json:"Import-only,omitempty"`
	Comments        string                    `json:"Comments,omitempty"`
	Id              string                    `json:"Region Entity Code"` // Region Entity Code
	DXCCEntityCode  string                    `json:"DXCC Entity Code"`   // TODO: get this to deserialize into dxccentitycode.DXCCEntityCode...
	Region          string                    `json:"Region"`
	Prefix          string                    `json:"Prefix,omitempty"`
	Applicability   string                    `json:"Applicability,omitempty"`
	StartDate       spectype.AdifSpecDateOnly `json:"Start Date,omitempty"`
	EndDate         spectype.AdifSpecDateOnly `json:"End Date,omitempty"`
}

func (r Spec) Description() string {
	return r.Region
}

// SpecMap contains all Region specifications as defined by the ADIF Workgroup specification exports.
type SpecMap struct {
	Header  []string        `json:"Header"`
	Records map[Region]Spec `json:"Records"`
}
