package morsekeytype

import (
	"fmt"
	"strconv"

	"github.com/hamradiolog-net/adif-spec/v6/codegen"
	"github.com/hamradiolog-net/adif-spec/v6/spectype"
)

// SpecMapContainer contains all MorseKeyType specifications as defined by the ADIF Workgroup specification exports.
type SpecMapContainer struct {
	// Header  []string         `json:"Header"`
	Records map[MorseKeyType]Spec `json:"Records"`
}

// Spec represents the specification for a single MorseKeyType as defined by the ADIF Workgroup specification exports.
type Spec struct {
	// EnumerationName string           `json:"Enumeration Name"`
	IsImportOnly spectype.Boolean `json:"Import-only,omitempty"`
	// Comments         string           `json:"Comments,omitempty"`
	Key              MorseKeyType `json:"Abbreviation"` // Abbreviation
	Description      string       `json:"Meaning"`
	Characteristics  string       `json:"Characteristics,omitempty"`
	MorseComposition string       `json:"Morse Composition,omitempty"`
	Examples         string       `json:"Examples,omitempty"`
}

func (s Spec) CodeGeneratorMetadata() codegen.CodeGeneratorMetadataForEnum {
	return codegen.CodeGeneratorMetadataForEnum{
		ConstName:     codegen.ToGoIdentifier(string(s.Key)),
		ConstValue:    strconv.QuoteToASCII(string(s.Key)),
		ConstComments: fmt.Sprintf("%-4s = %s", s.Key, s.Description),
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
		PackageName: "morsekeytype",
		DataType:    "MorseKeyType",
	}
}
