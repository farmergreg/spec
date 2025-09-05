package dxccentitycode

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/hamradiolog-net/adif-spec/v6/internal/codegen"
	"github.com/hamradiolog-net/adif-spec/v6/spectype"
)

var (
	_ codegen.CodeGenContainer = SpecMapContainer{}
	_ codegen.CodeGenSpec      = Spec{}
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

func (s Spec) CodeGeneratorMetadata() codegen.CodeGeneratorMetadataForEnum {
	deleted := ""
	if s.IsDeleted {
		deleted = " (DELETED) "
	}
	return codegen.CodeGeneratorMetadataForEnum{
		ConstName:     codegen.ToGoIdentifier(s.Identifier()),
		ConstValue:    strconv.Itoa(int(s.Key)),
		ConstComments: fmt.Sprintf("%s = %s%s", s.Key, s.EntityName, deleted),
		IsDeprecated:  bool(s.IsImportOnly),
	}
}

func (c SpecMapContainer) CodeGeneratorRecords() map[codegen.CodeGeneratorEnumValue]codegen.CodeGenSpec {
	result := make(map[codegen.CodeGeneratorEnumValue]codegen.CodeGenSpec, len(c.Records))
	for k, v := range c.Records {
		result[k] = v
	}
	return result
}

func (c SpecMapContainer) CodeGeneratorMetadata() codegen.CodeGeneratorMetadataForContainer {
	return codegen.CodeGeneratorMetadataForContainer{
		PackageName: "dxccentitycode",
		DataType:    "DXCCEntityCode",
	}
}
