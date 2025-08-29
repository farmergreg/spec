package spec

import "github.com/hamradiolog-net/adif-spec/src/pkg/enum/region"

// RegionSpec represents the specification for a single Region
type RegionSpec struct {
	EnumerationName string           `json:"Enumeration Name"`
	IsImportOnly    AdifSpecBoolean  `json:"Import-only,omitempty"`
	Comments        string           `json:"Comments,omitempty"`
	Id              string           `json:"Region Entity Code"` // Region Entity Code
	DXCCEntityCode  string           `json:"DXCC Entity Code"`
	Region          string           `json:"Region"`
	Prefix          string           `json:"Prefix,omitempty"`
	Applicability   string           `json:"Applicability,omitempty"`
	StartDate       AdifSpecDateOnly `json:"Start Date,omitempty"`
	EndDate         AdifSpecDateOnly `json:"End Date,omitempty"`
}

func (r RegionSpec) Description() string {
	return r.Region
}

// RegionSpecMap contains all RegionSpec specifications.
type RegionSpecMap struct {
	Header  []string                     `json:"Header"`
	Records map[region.Region]RegionSpec `json:"Records"`
}
