package continent

import (
	"fmt"
	"strconv"

	"github.com/hamradiolog-net/adif-spec/v6/codegen"
	"github.com/hamradiolog-net/adif-spec/v6/spectype"
)

// SpecMapContainer contains all Continent specifications as defined by the ADIF Workgroup specification exports.
type SpecMapContainer struct {
	// Header  []string         `json:"Header"`
	Records map[Continent]Spec `json:"Records"`
}

// Spec represents the specification for a single Continent as defined by the ADIF Workgroup specification exports.
type Spec struct {
	// EnumerationName string           `json:"Enumeration Name"`
	IsImportOnly spectype.Boolean `json:"Import-only,omitempty"`
	// Comments     string           `json:"Comments,omitempty"`
	Key       Continent `json:"Abbreviation"` // Abbreviation
	Continent string    `json:"Continent"`
}

func (s Spec) CodeGeneratorMetadata() codegen.CodeGeneratorMetadataForEnum {
	return codegen.CodeGeneratorMetadataForEnum{
		ConstName:     codegen.ToGoIdentifier(string(s.Key)),
		ConstValue:    strconv.QuoteToASCII(string(s.Key)),
		ConstComments: fmt.Sprintf("%s = %s", s.Key, s.Continent),
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
		PackageName: "continent",
		DataType:    "Continent",
	}
}
