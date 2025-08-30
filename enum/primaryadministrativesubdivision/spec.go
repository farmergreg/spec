package primaryadministrativesubdivision

import (
	"fmt"

	"github.com/hamradiolog-net/adif-spec/v3/enum/dxccentitycode"
	"github.com/hamradiolog-net/adif-spec/v3/spectype"
)

// SpecMap contains all PrimaryAdministrativeSubdivision specifications as defined by the ADIF Workgroup specification exports.
type SpecMapContainer struct {
	// Header  []string         `json:"Header"`
	Records map[PrimaryAdministrativeSubdivision]Spec `json:"Records"`
}

// Spec represents the specification for a single PrimaryAdministrativeSubdivision as defined by the ADIF Workgroup specification exports.
type Spec struct {
	// EnumerationName string           `json:"Enumeration Name"`
	IsImportOnly    spectype.Boolean              `json:"Import-only,omitempty"`
	Comments        string                        `json:"Comments,omitempty"`
	Code            string                        `json:"Code"` // Code
	PrimaryAdminSub string                        `json:"Primary Administrative Subdivision"`
	DXCCEntityCode  dxccentitycode.DXCCEntityCode `json:"DXCC Entity Code"`
	ContainedWithin string                        `json:"Contained Within,omitempty"`
	Oblast          string                        `json:"Oblast #,omitempty"`
	CQZone          string                        `json:"CQ Zone,omitempty"`
	ITUZone         string                        `json:"ITU Zone,omitempty"`
	Prefix          string                        `json:"Prefix,omitempty"`
	IsDeleted       spectype.Boolean              `json:"Deleted,omitempty"`
}

func (s Spec) String() string {
	return fmt.Sprintf("%5s.%-5s = %-5s ( %-5s )", s.Code, s.DXCCEntityCode, s.Code, s.PrimaryAdminSub)
}
