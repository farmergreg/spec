package spec

import (
	"strings"

	"github.com/hamradiolog-net/adif-spec/src/pkg/enum/dxccentitycode"
	"github.com/hamradiolog-net/adif-spec/src/pkg/spectype"
)

// DXCCEntityCodeSpec represents the specification for a single DXCCEntityCode
type DXCCEntityCodeSpec struct {
	EnumerationName string                   `json:"Enumeration Name"`
	IsImportOnly    spectype.AdifSpecBoolean `json:"Import-only,omitempty"`
	Comments        string                   `json:"Comments,omitempty"`
	Id              string                   `json:"Entity Code"` // Entity Code
	EntityName      string                   `json:"Entity Name"`
	IsDeleted       spectype.AdifSpecBoolean `json:"Deleted,omitempty"`
}

func (d DXCCEntityCodeSpec) Identifier() string {
	name := d.Description()

	if strings.HasPrefix(name, "None ") {
		name = "NONE"
	}

	if bool(d.IsDeleted) {
		name += "_DELETED"
	}

	return name
}

func (d DXCCEntityCodeSpec) Description() string {
	return d.EntityName
}

// DXCCEntityCodeSpecMap contains all DXCCEntityCodeSpec specifications.
type DXCCEntityCodeSpecMap struct {
	Header  []string                                             `json:"Header"`
	Records map[dxccentitycode.DXCCEntityCode]DXCCEntityCodeSpec `json:"Records"`
}
