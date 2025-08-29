package secondaryadministrativesubdivision

import (
	"github.com/hamradiolog-net/adif-spec/src/pkg/spectype"
)

// Spec represents the specification for a single SecondaryAdministrativeSubdivision as defined by the ADIF Workgroup specification exports.
type Spec struct {
	EnumerationName        string           `json:"Enumeration Name"`
	IsImportOnly           spectype.Boolean `json:"Import-only,omitempty"`
	Comments               string           `json:"Comments,omitempty"`
	Id                     string           `json:"Code"` // Code
	SecondaryAdminSub      string           `json:"Secondary Administrative Subdivision"`
	DXCCEntityCode         string           `json:"DXCC Entity Code"` // TODO: get this to deserialize into dxccentitycode.DXCCEntityCode...
	AlaskaJudicialDistrict string           `json:"Alaska Judicial District,omitempty"`
	IsDeleted              spectype.Boolean `json:"Deleted,omitempty"`
}

func (s Spec) Description() string {
	return s.SecondaryAdminSub
}

// SpecMap contains all SecondaryAdministrativeSubdivision specifications as defined by the ADIF Workgroup specification exports.
type SpecMap struct {
	Header  []string                                    `json:"Header"`
	Records map[SecondaryAdministrativeSubdivision]Spec `json:"Records"`
}
