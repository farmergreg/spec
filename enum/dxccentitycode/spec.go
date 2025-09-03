package dxccentitycode

import (
	"strings"

	"github.com/hamradiolog-net/adif-spec/v6/spectype"
)

// SpecMapContainer contains all DXCCEntityCode specifications as defined by the ADIF Workgroup specification exports.
type SpecMapContainer struct {
	// Header  []string         `json:"Header"`
	Records map[DXCCEntityCode]Spec `json:"Records"`
}

// Spec represents the specification for a single DXCCEntityCode as defined by the ADIF Workgroup specification exports.
type Spec struct {
	// EnumerationName string           `json:"Enumeration Name"`
	IsImportOnly spectype.Boolean `json:"Import-only,omitempty"`
	// Comments     string           `json:"Comments,omitempty"`
	Key        DXCCEntityCode   `json:"Entity Code"` // Entity Code
	EntityName string           `json:"Entity Name"`
	IsDeleted  spectype.Boolean `json:"Deleted,omitempty"`
}

func (d Spec) Identifier() string {
	name := d.EntityName

	if strings.HasPrefix(name, "None ") {
		name = "NONE"
	}

	if bool(d.IsDeleted) {
		name += "_DELETED"
	}

	return name
}
