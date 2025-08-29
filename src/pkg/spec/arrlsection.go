package spec

import (
	"github.com/hamradiolog-net/adif-spec/src/pkg/enum/dxccentitycode"
)

// ARRLSectionSpec represents the specification for a single ARRLSection
type ARRLSectionSpec struct {
	EnumerationName string                            `json:"Enumeration Name"`
	IsImportOnly    AdifSpecBoolean                   `json:"Import-only,omitempty"`
	Comments        string                            `json:"Comments,omitempty"`
	Id              string                            `json:"Abbreviation"` // Abbreviation
	Description     string                            `json:"Section Name"` // Section Name
	DXCCEntityCode  dxccentitycode.DXCCEntityCodeList `json:"DXCC Entity Code"`
	FromDate        AdifSpecDateTime                  `json:"From Date,omitempty"`
	DeletedDate     AdifSpecDateTime                  `json:"Deleted Date,omitempty"`
}

// ARRLSectionSpecMap contains all ARRLSectionSpec specifications.
type ARRLSectionSpecMap struct {
	Header  []string                   `json:"Header"`
	Records map[string]ARRLSectionSpec `json:"Records"`
}
