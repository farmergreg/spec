package spec

import (
	"github.com/hamradiolog-net/adif-spec/src/pkg/enum/dxccentitycode"
)

// ARRLSectionRecord represents a single ARRL section record
type ARRLSectionRecord struct {
	BaseEnumerationRecord
	Id             string                            `json:"Abbreviation"` // Abbreviation
	Description    string                            `json:"Section Name"` // Section Name
	DXCCEntityCode dxccentitycode.DXCCEntityCodeList `json:"DXCC Entity Code"`
	FromDate       AdifSpecDateTime                  `json:"From Date,omitempty"`
	DeletedDate    AdifSpecDateTime                  `json:"Deleted Date,omitempty"`
}

// ARRLSectionEnumeration represents the complete ARRL section enumeration
type ARRLSectionEnumeration struct {
	Header  []string                     `json:"Header"`
	Records map[string]ARRLSectionRecord `json:"Records"`
}
