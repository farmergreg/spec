package primaryadministrativesubdivision

import (
	"fmt"
	"strconv"

	"github.com/hamradiolog-net/spec/v6/enum/dxccentitycode"
	"github.com/hamradiolog-net/spec/v6/internal/codegen"
	"github.com/hamradiolog-net/spec/v6/spectype"
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

// String returns the string representation of the PrimaryAdministrativeSubdivisionCode.
func (c PrimaryAdministrativeSubdivisionCode) String() string {
	return string(c)
}

func (s Spec) CodeGenMetadata() codegen.CodeGenEnumMetadata {
	constName := string(s.Code) + "." + strconv.Itoa(int(s.DXCCEntityCode))
	if s.IsDeleted {
		constName += ".Deleted.0"
	}
	return codegen.CodeGenEnumMetadata{
		ConstName:     strconv.QuoteToASCII(constName),
		ConstValue:    strconv.QuoteToASCII(string(s.Code)),
		ConstComments: fmt.Sprintf("%5s.%-5s = %-5s ( %-5s ); IMPORTANT: This is NOT the Primary Administrative Subdivision Code. It is a lookup key for use with PrimaryAdministrativeSubdivisionCompositeKeyMap", s.Code, s.DXCCEntityCode, s.Code, s.PrimaryAdminSub),
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
	keyMap := make(map[string]string)
	for _, v := range c.Records {
		keyMap["PrimaryAdministrativeSubdivision_"+string(v.Code)] = string(v.Code)
	}

	return codegen.CodeGenContainerMetadata{
		PackageName:      "primaryadministrativesubdivision",
		DataType:         "PrimaryAdministrativeSubdivisionCompositeKey",
		CompositeKeyType: "PrimaryAdministrativeSubdivisionCode",
		CompositeKeyMap:  keyMap,
	}
}
