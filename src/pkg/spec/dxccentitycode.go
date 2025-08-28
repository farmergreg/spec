package spec

import (
	"strings"
)

// DXCCEntityCodeRecord represents a single DXCC entity code record
type DXCCEntityCodeRecord struct {
	BaseEnumerationSpec
	Id         string          `json:"Entity Code"` // Entity Code
	EntityName string          `json:"Entity Name"`
	IsDeleted  AdifSpecBoolean `json:"Deleted,omitempty"`
}

func (d DXCCEntityCodeRecord) Identifier() string {
	name := d.Description()

	if strings.HasPrefix(name, "None ") {
		name = "NONE"
	}

	if bool(d.IsDeleted) {
		name += "_DELETED"
	}

	return name
}

func (d DXCCEntityCodeRecord) Description() string {
	return d.EntityName
}

// DXCCEntityCodeEnumeration represents the complete DXCC entity code enumeration
type DXCCEntityCodeEnumeration struct {
	Header  []string                        `json:"Header"`
	Records map[string]DXCCEntityCodeRecord `json:"Records"`
}
