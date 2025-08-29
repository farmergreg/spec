package spec

import "github.com/hamradiolog-net/adif-spec/src/pkg/enum/secondaryadministrativesubdivision"

// SecondaryAdministrativeSubdivisionSpec represents the specification for a single SecondaryAdministrativeSubdivision
type SecondaryAdministrativeSubdivisionSpec struct {
	EnumerationName        string          `json:"Enumeration Name"`
	IsImportOnly           AdifSpecBoolean `json:"Import-only,omitempty"`
	Comments               string          `json:"Comments,omitempty"`
	Id                     string          `json:"Code"` // Code
	SecondaryAdminSub      string          `json:"Secondary Administrative Subdivision"`
	DXCCEntityCode         string          `json:"DXCC Entity Code"`
	AlaskaJudicialDistrict string          `json:"Alaska Judicial District,omitempty"`
	IsDeleted              AdifSpecBoolean `json:"Deleted,omitempty"`
}

func (s SecondaryAdministrativeSubdivisionSpec) Description() string {
	return s.SecondaryAdminSub
}

// SecondaryAdministrativeSubdivisionSpecMap contains all SecondaryAdministrativeSubdivisionSpec specifications.
type SecondaryAdministrativeSubdivisionSpecMap struct {
	Header  []string                                                                                                         `json:"Header"`
	Records map[secondaryadministrativesubdivision.SecondaryAdministrativeSubdivision]SecondaryAdministrativeSubdivisionSpec `json:"Records"`
}
