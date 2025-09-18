package dxccentitycode

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/farmergreg/spec/v6/internal/codegen"
	"github.com/farmergreg/spec/v6/spectype"
)

var (
	_ codegen.CodeGenContainer = SpecMapContainer{}
	_ codegen.CodeGenSpec      = Spec{}
)

const nonePrefix = "None "

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

	if strings.HasPrefix(name, nonePrefix) {
		name = "NONE"
	}

	if bool(d.IsDeleted) {
		name += "_DELETED"
	}

	return name
}

func (s Spec) CodeGenMetadata() codegen.CodeGenEnumMetadata {
	deleted := ""
	if s.IsDeleted {
		deleted = " (DELETED) "
	}
	return codegen.CodeGenEnumMetadata{
		ConstName:     codegen.ToGoIdentifier(s.Identifier()),
		ConstValue:    strconv.Itoa(int(s.Key)),
		ConstComments: fmt.Sprintf("%s = %s%s", s.Key, s.EntityName, deleted),
		IsDeprecated:  bool(s.IsImportOnly),
	}
}

func (c SpecMapContainer) CodeGenRecords() map[codegen.CodeGenKey]codegen.CodeGenSpec {
	result := make(map[codegen.CodeGenKey]codegen.CodeGenSpec, len(c.Records))
	for k, v := range c.Records {
		if strings.HasPrefix(v.EntityName, nonePrefix) {
			v.EntityName = "NONE"
		}
		result[k] = v
	}
	return result
}

func (c SpecMapContainer) CodeGenMetadata() codegen.CodeGenContainerMetadata {
	return codegen.CodeGenContainerMetadata{
		PackageName: "dxccentitycode",
		DataType:    "DXCCEntityCode",
	}
}
