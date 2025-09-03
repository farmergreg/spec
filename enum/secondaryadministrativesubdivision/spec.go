package secondaryadministrativesubdivision

import (
	"github.com/hamradiolog-net/adif-spec/v6/enum/dxccentitycode"
	"github.com/hamradiolog-net/adif-spec/v6/spectype"
)

// SpecMapContainer contains all SecondaryAdministrativeSubdivision specifications as defined by the ADIF Workgroup specification exports.
type SpecMapContainer struct {
	// Header  []string         `json:"Header"`
	Records map[SecondaryAdministrativeSubdivision]Spec `json:"Records"`
}

// Spec represents the specification for a single SecondaryAdministrativeSubdivision as defined by the ADIF Workgroup specification exports.
type Spec struct {
	// EnumerationName string           `json:"Enumeration Name"`
	IsImportOnly spectype.Boolean `json:"Import-only,omitempty"`
	// Comments               string                             `json:"Comments,omitempty"`
	Key                    SecondaryAdministrativeSubdivision `json:"Code"` // Code
	SecondaryAdminSub      string                             `json:"Secondary Administrative Subdivision"`
	DXCCEntityCode         dxccentitycode.DXCCEntityCode      `json:"DXCC Entity Code"`
	AlaskaJudicialDistrict string                             `json:"Alaska Judicial District,omitempty"`
	IsDeleted              spectype.Boolean                   `json:"Deleted,omitempty"`
}
