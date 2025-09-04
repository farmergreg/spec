package aditype

import (
	"strconv"

	"github.com/hamradiolog-net/adif-spec/v6/codegen"
	"github.com/hamradiolog-net/adif-spec/v6/spectype"
)

// SpecMapContainer contains all DataType specifications as defined by the ADIF Workgroup specification exports.
type SpecMapContainer struct {
	Header  []string         `json:"Header"`
	Records map[ADIType]Spec `json:"Records"`
}

// Spec represents the specification for a single DataType as defined by the ADIF Workgroup specification exports.
type Spec struct {
	Key               ADIType          `json:"Data Type Name"` // Data Type Name
	DataTypeIndicator string           `json:"Data Type Indicator,omitempty"`
	Description       string           `json:"Description"`
	MinimumValue      spectype.Integer `json:"Minimum Value,omitempty"`
	MaximumValue      spectype.Integer `json:"Maximum Value,omitempty"`
	IsImportOnly      spectype.Boolean `json:"Import-only,omitempty"`
	Comments          string           `json:"Comments,omitempty"`
}

func (s Spec) CodeGeneratorMetadata() codegen.CodeGeneratorMetadataForEnum {
	return codegen.CodeGeneratorMetadataForEnum{
		ConstName:     codegen.ToGoIdentifier(string(s.Key)),
		ConstValue:    strconv.QuoteToASCII(string(s.Key)),
		ConstComments: s.Description,
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
		PackageName: "aditype",
		DataType:    "ADIType",
	}
}
