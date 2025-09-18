package submode

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/farmergreg/spec/v6/internal/codegen"
	"github.com/farmergreg/spec/v6/spectype"
)

var (
	_ codegen.CodeGenContainer = SpecMapContainer{}
	_ codegen.CodeGenSpec      = Spec{}
)

// SpecMapContainer contains all Submode specifications as defined by the ADIF Workgroup specification exports.
type SpecMapContainer struct {
	// Header  []string         `json:"Header"`
	Records map[SubMode]Spec `json:"Records"`
}

// Spec represents the specification for a single Submode as defined by the ADIF Workgroup specification exports.
type Spec struct {
	// EnumerationName string           `json:"Enumeration Name"`
	IsImportOnly spectype.Boolean `json:"Import-only,omitempty"`
	// Comments     string           `json:"Comments,omitempty"`
	Key         SubMode `json:"Submode"` // Submode
	Mode        string  `json:"Mode"`
	Description string  `json:"Description,omitempty"`
}

func (s Spec) CodeGenMetadata() codegen.CodeGenEnumMetadata {
	return codegen.CodeGenEnumMetadata{
		ConstName:     codegen.ToGoIdentifier("SUBMODE_" + string(s.Key)),
		ConstValue:    strconv.QuoteToASCII(string(s.Key)),
		ConstComments: fmt.Sprintf("%-15s = %-15s %s", s.Key, s.Mode, s.Description),
		IsDeprecated:  bool(s.IsImportOnly),
	}
}

func (c SpecMapContainer) CodeGenRecords() map[codegen.CodeGenKey]codegen.CodeGenSpec {
	result := make(map[codegen.CodeGenKey]codegen.CodeGenSpec, len(c.Records))
	for k, v := range c.Records {
		v.Key = SubMode(strings.ToLower(string(v.Key)))
		v.Mode = strings.ToLower(v.Mode)
		result[k] = v
	}
	return result
}

func (c SpecMapContainer) CodeGenMetadata() codegen.CodeGenContainerMetadata {
	return codegen.CodeGenContainerMetadata{
		PackageName: "submode",
		DataType:    "SubMode",
	}
}
