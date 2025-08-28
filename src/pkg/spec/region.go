package spec

// RegionRecord represents a single region record
type RegionRecord struct {
	BaseEnumerationRecord
	Id             string           `json:"Region Entity Code"` // Region Entity Code
	DXCCEntityCode string           `json:"DXCC Entity Code"`
	Region         string           `json:"Region"`
	Prefix         string           `json:"Prefix,omitempty"`
	Applicability  string           `json:"Applicability,omitempty"`
	StartDate      AdifSpecDateOnly `json:"Start Date,omitempty"`
	EndDate        AdifSpecDateOnly `json:"End Date,omitempty"`
}

func (r RegionRecord) Description() string {
	return r.Region
}

// RegionEnumeration represents the complete region enumeration
type RegionEnumeration struct {
	Header  []string                `json:"Header"`
	Records map[string]RegionRecord `json:"Records"`
}
