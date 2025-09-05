package secondaryadministrativesubdivisionalt

import (
	"fmt"
	"strconv"

	"github.com/hamradiolog-net/adif-spec/v6/enum/dxccentitycode"
	"github.com/hamradiolog-net/adif-spec/v6/internal/codegen"
	"github.com/hamradiolog-net/adif-spec/v6/spectype"
)

var (
	_ codegen.CodeGenContainer = SpecMapContainer{}
	_ codegen.CodeGenSpec      = Spec{}
)

// SpecMapContainer contains all SecondaryAdministrativeSubdivisionAlt specifications as defined by the ADIF Workgroup specification exports.
type SpecMapContainer struct {
	// Header  []string         `json:"Header"`
	Records map[SecondaryAdministrativeSubdivisionAlt]Spec `json:"Records"`
}

// Spec represents the specification for a single SecondaryAdministrativeSubdivisionAlt as defined by the ADIF Workgroup specification exports.
type Spec struct {
	// EnumerationName string           `json:"Enumeration Name"`
	IsImportOnly spectype.Boolean `json:"Import-only,omitempty"`
	// Comments       string                                `json:"Comments,omitempty"`
	Key            SecondaryAdministrativeSubdivisionAlt `json:"Code"` // Code
	DXCCEntityCode dxccentitycode.DXCCEntityCode         `json:"DXCC Entity Code"`
	Region         string                                `json:"Region"`
	District       string                                `json:"District"`
	IsDeleted      spectype.Boolean                      `json:"Deleted,omitempty"`
}

func (s Spec) CodeGeneratorMetadata() codegen.CodeGeneratorMetadataForEnum {
	return codegen.CodeGeneratorMetadataForEnum{
		ConstName:     codegen.ToGoIdentifier(string(s.Key)),
		ConstValue:    strconv.QuoteToASCII(string(s.Key)),
		ConstComments: fmt.Sprintf("%-50s = DXCC: %s %s", s.Key, s.DXCCEntityCode, s.Region),
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
		PackageName: "secondaryadministrativesubdivisionalt",
		DataType:    "SecondaryAdministrativeSubdivisionAlt",
	}
}
