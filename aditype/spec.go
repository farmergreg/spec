package aditype

import (
	"strconv"

	"github.com/hamradiolog-net/spec/v6/internal/codegen"
	"github.com/hamradiolog-net/spec/v6/spectype"
)

var (
	_ codegen.CodeGenContainer = SpecMapContainer{}
	_ codegen.CodeGenSpec      = Spec{}
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

func (s Spec) CodeGenMetadata() codegen.CodeGenEnumMetadata {
	return codegen.CodeGenEnumMetadata{
		ConstName:     codegen.ToGoIdentifier(string(s.Key)),
		ConstValue:    strconv.QuoteToASCII(string(s.Key)),
		ConstComments: s.Description,
		IsDeprecated:  bool(s.IsImportOnly),
	}
}

func (c SpecMapContainer) CodeGenRecords() map[codegen.CodeGenKey]codegen.CodeGenSpec {
	result := make(map[codegen.CodeGenKey]codegen.CodeGenSpec, len(c.Records))
	for k, v := range c.Records {
		result[k] = v
	}
	return result
}

func (c SpecMapContainer) CodeGenMetadata() codegen.CodeGenContainerMetadata {
	return codegen.CodeGenContainerMetadata{
		PackageName: "aditype",
		DataType:    "ADIType",
	}
}
