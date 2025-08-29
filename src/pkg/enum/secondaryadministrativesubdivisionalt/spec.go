package secondaryadministrativesubdivisionalt

import (
	"github.com/hamradiolog-net/adif-spec/src/pkg/spectype"
)

// Spec represents the specification for a single SecondaryAdministrativeSubdivisionAlt as defined by the ADIF Workgroup specification exports.
type Spec struct {
	EnumerationName string                   `json:"Enumeration Name"`
	IsImportOnly    spectype.AdifSpecBoolean `json:"Import-only,omitempty"`
	Comments        string                   `json:"Comments,omitempty"`
	Id              string                   `json:"Code"`             // Code
	DXCCEntityCode  string                   `json:"DXCC Entity Code"` // TODO: get this to deserialize into dxccentitycode.DXCCEntityCode...
	Region          string                   `json:"Region"`
	District        string                   `json:"District"`
	IsDeleted       spectype.AdifSpecBoolean `json:"Deleted,omitempty"`
}

func (s Spec) Description() string {
	return s.Region
}

// SpecMap contains all SecondaryAdministrativeSubdivisionAlt specifications as defined by the ADIF Workgroup specification exports.
type SpecMap struct {
	Header  []string                                       `json:"Header"`
	Records map[SecondaryAdministrativeSubdivisionAlt]Spec `json:"Records"`
}
