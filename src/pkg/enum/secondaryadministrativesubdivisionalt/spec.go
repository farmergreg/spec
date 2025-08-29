package secondaryadministrativesubdivisionalt

import (
	"github.com/hamradiolog-net/adif-spec/src/pkg/spectype"
)

// SecondaryAdministrativeSubdivisionAltSpec represents the specification for a single SecondaryAdministrativeSubdivisionAlt
type SecondaryAdministrativeSubdivisionAltSpec struct {
	EnumerationName string                   `json:"Enumeration Name"`
	IsImportOnly    spectype.AdifSpecBoolean `json:"Import-only,omitempty"`
	Comments        string                   `json:"Comments,omitempty"`
	Id              string                   `json:"Code"` // Code
	DXCCEntityCode  string                   `json:"DXCC Entity Code"`
	Region          string                   `json:"Region"`
	District        string                   `json:"District"`
	IsDeleted       spectype.AdifSpecBoolean `json:"Deleted,omitempty"`
}

func (s SecondaryAdministrativeSubdivisionAltSpec) Description() string {
	return s.Region
}

// SecondaryAdministrativeSubdivisionAltSpecMap contains all SecondaryAdministrativeSubdivisionAltSpec specifications.
type SecondaryAdministrativeSubdivisionAltSpecMap struct {
	Header  []string                           `json:"Header"`
	Records map[SecondaryAdministrativeSubdivisionAlt]SecondaryAdministrativeSubdivisionAltSpec `json:"Records"`
}
