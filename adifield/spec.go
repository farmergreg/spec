package adifield

import (
	"fmt"
	"path"
	"reflect"
	"strconv"

	"github.com/hamradiolog-net/adif-spec/v6/codegen"
	"github.com/hamradiolog-net/adif-spec/v6/spectype"
)

// SpecMapContainer contains all Field specifications as defined by the ADIF Workgroup specification exports.
type SpecMapContainer struct {
	// Header  []string         `json:"Header"`
	Records map[ADIField]Spec `json:"Records"`
}

// Spec represents the specification for a single Field as defined by the ADIF Workgroup specification exports.
type Spec struct {
	Key      ADIField `json:"Field Name"` // Field Name
	DataType string   `json:"Data Type"`
	// Enumeration   string           `json:"Enumeration,omitempty"`
	Description   string           `json:"Description"`
	IsHeaderField spectype.Boolean `json:"Header Field,omitempty"`
	MinimumValue  spectype.Integer `json:"Minimum Value,omitempty"`
	MaximumValue  spectype.Integer `json:"Maximum Value,omitempty"`
	IsImportOnly  spectype.Boolean `json:"Import-only,omitempty"`
	Comments      string           `json:"Comments,omitempty"`
}

func (s Spec) CodeGeneratorMetadata() codegen.CodeGeneratorMetadataForEnum {
	var headerOrRecord = "Record"
	if s.IsHeaderField {
		headerOrRecord = "Header"
	}

	return codegen.CodeGeneratorMetadataForEnum{
		ConstName:     codegen.ToGoIdentifier(string(s.Key)),
		ConstDataType: reflect.TypeOf(s.Key).Name(),
		ConstValue:    strconv.QuoteToASCII(string(s.Key)),
		Comments:      fmt.Sprintf("%s: %s", headerOrRecord, s.Description),
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
		PackageName: path.Base(reflect.TypeOf(c).PkgPath()),
	}
}
