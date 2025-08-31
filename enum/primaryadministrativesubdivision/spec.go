package primaryadministrativesubdivision

import (
	"fmt"

	"github.com/hamradiolog-net/adif-spec/v6/enum/dxccentitycode"
	"github.com/hamradiolog-net/adif-spec/v6/spectype"
)

// SpecMap contains all PrimaryAdministrativeSubdivision specifications as defined by the ADIF Workgroup specification exports.
type SpecMapContainer struct {
	// Header  []string         `json:"Header"`
	Records map[PrimaryAdministrativeSubdivisionCompositeKey]Spec `json:"Records"`
}

// Spec represents the specification for a single PrimaryAdministrativeSubdivision as defined by the ADIF Workgroup specification exports.
type Spec struct {
	// EnumerationName string           `json:"Enumeration Name"`
	IsImportOnly    spectype.Boolean                     `json:"Import-only,omitempty"`
	Comments        string                               `json:"Comments,omitempty"`
	Code            PrimaryAdministrativeSubdivisionCode `json:"Code"` // Code
	PrimaryAdminSub string                               `json:"Primary Administrative Subdivision"`
	DXCCEntityCode  dxccentitycode.DXCCEntityCode        `json:"DXCC Entity Code"`
	ContainedWithin string                               `json:"Contained Within,omitempty"`
	Oblast          spectype.Oblast                      `json:"Oblast #,omitempty"`
	CQZone          spectype.CQZoneList                  `json:"CQ Zone,omitempty"`
	ITUZone         spectype.ITUZoneList                 `json:"ITU Zone,omitempty"`
	Prefix          string                               `json:"Prefix,omitempty"`
	IsDeleted       spectype.Boolean                     `json:"Deleted,omitempty"`
}

// PrimaryAdministrativeSubdivisionCode is the Code portion of the composite key.
type PrimaryAdministrativeSubdivisionCode string

func (s Spec) String() string {
	return fmt.Sprintf("%5s.%-5s = %-5s ( %-5s ); IMPORTANT: This is NOT the Primary Administrative Subdivision Code. It is a lookup key for use with PrimaryAdministrativeSubdivisionCompositeKeyMap", s.Code, s.DXCCEntityCode, s.Code, s.PrimaryAdminSub)
}
