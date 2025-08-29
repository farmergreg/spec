package arrlsection

import (
	"github.com/hamradiolog-net/adif-spec/src/pkg/enum/dxccentitycode"
	"github.com/hamradiolog-net/adif-spec/src/pkg/spectype"
)

// ARRLSectionSpec represents the specification for a single ARRLSection
type ARRLSectionSpec struct {
	EnumerationName string                            `json:"Enumeration Name"`
	IsImportOnly    spectype.AdifSpecBoolean          `json:"Import-only,omitempty"`
	Comments        string                            `json:"Comments,omitempty"`
	Id              string                            `json:"Abbreviation"` // Abbreviation
	Description     string                            `json:"Section Name"` // Section Name
	DXCCEntityCode  dxccentitycode.DXCCEntityCodeList `json:"DXCC Entity Code"`
	FromDate        spectype.AdifSpecDateTime         `json:"From Date,omitempty"`
	DeletedDate     spectype.AdifSpecDateTime         `json:"Deleted Date,omitempty"`
}

// SpecMap contains all ARRLSectionSpec specifications.
type SpecMap struct {
	Header  []string                        `json:"Header"`
	Records map[ARRLSection]ARRLSectionSpec `json:"Records"`
}
