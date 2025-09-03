package eqslag

import (
	"github.com/hamradiolog-net/adif-spec/v6/spectype"
)

// SpecMapContainer contains all EQSLAG specifications as defined by the ADIF Workgroup specification exports.
type SpecMapContainer struct {
	// Header  []string         `json:"Header"`
	Records map[EQSLAG]Spec `json:"Records"`
}

// Spec represents the specification for a single EQSLAG as defined by the ADIF Workgroup specification exports.
type Spec struct {
	// EnumerationName string           `json:"Enumeration Name"`
	IsImportOnly spectype.Boolean `json:"Import-only,omitempty"`
	// Comments     string           `json:"Comments,omitempty"`
	Key         EQSLAG `json:"Status"` // Status
	Description string `json:"Description"`
}
