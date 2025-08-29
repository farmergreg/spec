package dxccentitycode

import (
	"strings"

	"github.com/hamradiolog-net/adif-spec/src/pkg/spectype"
)

// Spec represents the specification for a single DXCCEntityCode as defined by the ADIF Workgroup specification exports.
type Spec struct {
	EnumerationName string                   `json:"Enumeration Name"`
	IsImportOnly    spectype.AdifSpecBoolean `json:"Import-only,omitempty"`
	Comments        string                   `json:"Comments,omitempty"`
	Id              string                   `json:"Entity Code"` // Entity Code
	EntityName      string                   `json:"Entity Name"`
	IsDeleted       spectype.AdifSpecBoolean `json:"Deleted,omitempty"`
}

func (d Spec) Identifier() string {
	name := d.Description()

	if strings.HasPrefix(name, "None ") {
		name = "NONE"
	}

	if bool(d.IsDeleted) {
		name += "_DELETED"
	}

	return name
}

func (d Spec) Description() string {
	return d.EntityName
}

// SpecMap contains all DXCCEntityCode specifications as defined by the ADIF Workgroup specification exports.
type SpecMap struct {
	Header  []string                `json:"Header"`
	Records map[DXCCEntityCode]Spec `json:"Records"`
}
