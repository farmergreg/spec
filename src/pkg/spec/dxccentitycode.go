package spec

import (
	"strings"
)

// DXCCEntityCodeSpec represents the specification for a single DXCCEntityCode
type DXCCEntityCodeSpec struct {
	BaseEnumerationSpec
	Id         string          `json:"Entity Code"` // Entity Code
	EntityName string          `json:"Entity Name"`
	IsDeleted  AdifSpecBoolean `json:"Deleted,omitempty"`
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
	Header  []string                      `json:"Header"`
	Records map[string]DXCCEntityCodeSpec `json:"Records"`
}
