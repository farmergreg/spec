package primaryadministrativesubdivision

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

// SpecMapContainer contains all PrimaryAdministrativeSubdivision specifications as defined by the ADIF Workgroup specification exports.
type SpecMapContainer struct {
	// Header  []string         `json:"Header"`
	Records map[PrimaryAdministrativeSubdivisionCompositeKey]Spec `json:"Records"`
}

// Spec represents the specification for a single PrimaryAdministrativeSubdivision as defined by the ADIF Workgroup specification exports.
type Spec struct {
	// EnumerationName string           `json:"Enumeration Name"`
	IsImportOnly    spectype.Boolean                     `json:"Import-only,omitempty"`
	Comments        string                               `json:"Comments,omitempty"`
	Code            PrimaryAdministrativeSubdivisionCode `json:"Code"` // Code
	PrimaryAdminSub string                               `json:"Primary Administrative Subdivision"`
	DXCCEntityCode  dxccentitycode.DXCCEntityCode        `json:"DXCC Entity Code"`
	ContainedWithin string                               `json:"Contained Within,omitempty"`
	Oblast          Oblast                               `json:"Oblast #,omitempty"`
	CQZone          CQZoneList                           `json:"CQ Zone,omitempty"`
	ITUZone         ITUZoneList                          `json:"ITU Zone,omitempty"`
	Prefix          string                               `json:"Prefix,omitempty"`
	IsDeleted       spectype.Boolean                     `json:"Deleted,omitempty"`
}

// PrimaryAdministrativeSubdivisionCode is the Code portion of the composite key.
type PrimaryAdministrativeSubdivisionCode string

func (s Spec) CodeGeneratorMetadata() codegen.CodeGeneratorMetadataForEnum {
	constName := string(s.Code) + "." + strconv.Itoa(int(s.DXCCEntityCode))
	if s.IsDeleted {
		constName += ".Deleted.0"
	}
	return codegen.CodeGeneratorMetadataForEnum{
		ConstName:     strconv.QuoteToASCII(constName),
		ConstValue:    strconv.QuoteToASCII(string(s.Code)),
		ConstComments: fmt.Sprintf("%5s.%-5s = %-5s ( %-5s ); IMPORTANT: This is NOT the Primary Administrative Subdivision Code. It is a lookup key for use with PrimaryAdministrativeSubdivisionCompositeKeyMap", s.Code, s.DXCCEntityCode, s.Code, s.PrimaryAdminSub),
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
		PackageName: "primaryadministrativesubdivision",
		DataType:    "PrimaryAdministrativeSubdivisionCompositeKey",
	}
}
