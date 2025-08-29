package secondaryadministrativesubdivision

import (
	"github.com/hamradiolog-net/adif-spec/src/pkg/spectype"
)

// SecondaryAdministrativeSubdivisionSpec represents the specification for a single SecondaryAdministrativeSubdivision
type SecondaryAdministrativeSubdivisionSpec struct {
	EnumerationName        string                   `json:"Enumeration Name"`
	IsImportOnly           spectype.AdifSpecBoolean `json:"Import-only,omitempty"`
	Comments               string                   `json:"Comments,omitempty"`
	Id                     string                   `json:"Code"` // Code
	SecondaryAdminSub      string                   `json:"Secondary Administrative Subdivision"`
	DXCCEntityCode         string                   `json:"DXCC Entity Code"`
	AlaskaJudicialDistrict string                   `json:"Alaska Judicial District,omitempty"`
	IsDeleted              spectype.AdifSpecBoolean `json:"Deleted,omitempty"`
}

func (s SecondaryAdministrativeSubdivisionSpec) Description() string {
	return s.SecondaryAdminSub
}

// SpecMap contains all SecondaryAdministrativeSubdivisionSpec specifications.
type SpecMap struct {
	Header  []string                                                                      `json:"Header"`
	Records map[SecondaryAdministrativeSubdivision]SecondaryAdministrativeSubdivisionSpec `json:"Records"`
}
