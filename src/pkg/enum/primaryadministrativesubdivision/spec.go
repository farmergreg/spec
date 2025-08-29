package primaryadministrativesubdivision

import (
	"github.com/hamradiolog-net/adif-spec/src/pkg/spectype"
)

// PrimaryAdministrativeSubdivisionSpec represents the specification for a single PrimaryAdministrativeSubdivision
type PrimaryAdministrativeSubdivisionSpec struct {
	EnumerationName string                   `json:"Enumeration Name"`
	IsImportOnly    spectype.AdifSpecBoolean `json:"Import-only,omitempty"`
	Comments        string                   `json:"Comments,omitempty"`
	Code            string                   `json:"Code"` // Code
	PrimaryAdminSub string                   `json:"Primary Administrative Subdivision"`
	DXCCEntityCode  string                   `json:"DXCC Entity Code"`
	ContainedWithin string                   `json:"Contained Within,omitempty"`
	Oblast          string                   `json:"Oblast #,omitempty"`
	CQZone          string                   `json:"CQ Zone,omitempty"`
	ITUZone         string                   `json:"ITU Zone,omitempty"`
	Prefix          string                   `json:"Prefix,omitempty"`
	IsDeleted       spectype.AdifSpecBoolean `json:"Deleted,omitempty"`
}

func (p PrimaryAdministrativeSubdivisionSpec) Description() string {
	return p.PrimaryAdminSub
}

// PrimaryAdministrativeSubdivisionSpecMap contains all PrimaryAdministrativeSubdivisionSpec specifications.
type PrimaryAdministrativeSubdivisionSpecMap struct {
	Header  []string                           `json:"Header"`
	Records map[PrimaryAdministrativeSubdivision]PrimaryAdministrativeSubdivisionSpec `json:"Records"`
}
