package credit

import (
	"fmt"
	"strconv"

	"github.com/hamradiolog-net/adif-spec/v6/codegen"
	"github.com/hamradiolog-net/adif-spec/v6/spectype"
)

var (
	_ codegen.CodeGenContainer = SpecMapContainer{}
	_ codegen.CodeGenSpec      = Spec{}
)

// SpecMapContainer contains all Credit specifications as defined by the ADIF Workgroup specification exports.
type SpecMapContainer struct {
	// Header  []string         `json:"Header"`
	Records map[Credit]Spec `json:"Records"`
}

// Spec represents the specification for a single Credit as defined by the ADIF Workgroup specification exports.
type Spec struct {
	// EnumerationName string           `json:"Enumeration Name"`
	IsImportOnly spectype.Boolean `json:"Import-only,omitempty"`
	// Comments     string           `json:"Comments,omitempty"`
	Key     Credit `json:"Credit For"` // Credit For
	Sponsor string `json:"Sponsor"`
	Award   string `json:"Award"`
	Facet   string `json:"Facet"`
}

func (s Spec) CodeGeneratorMetadata() codegen.CodeGeneratorMetadataForEnum {
	return codegen.CodeGeneratorMetadataForEnum{
		ConstName:     codegen.ToGoIdentifier(string(s.Key)),
		ConstValue:    strconv.QuoteToASCII(string(s.Key)),
		ConstComments: fmt.Sprintf("%-20s = %-15s %-45s %-15s", s.Key, s.Sponsor, s.Award, s.Facet),
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
		PackageName: "credit",
		DataType:    "Credit",
	}
}
