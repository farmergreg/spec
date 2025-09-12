package contest

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/hamradiolog-net/spec/v6/internal/codegen"
	"github.com/hamradiolog-net/spec/v6/spectype"
)

var (
	_ codegen.CodeGenContainer = SpecMapContainer{}
	_ codegen.CodeGenSpec      = Spec{}
)

// SpecMapContainer contains all Contest specifications as defined by the ADIF Workgroup specification exports.
type SpecMapContainer struct {
	// Header  []string         `json:"Header"`
	Records map[Contest]Spec `json:"Records"`
}

// Spec represents the specification for a single Contest as defined by the ADIF Workgroup specification exports.
type Spec struct {
	// EnumerationName string           `json:"Enumeration Name"`
	IsImportOnly spectype.Boolean `json:"Import-only,omitempty"`
	// Comments     string           `json:"Comments,omitempty"`
	Key         Contest `json:"Contest-ID"` // Contest ID
	Description string  `json:"Description"`
}

func (s Spec) CodeGenMetadata() codegen.CodeGenEnumMetadata {
	return codegen.CodeGenEnumMetadata{
		ConstName:     codegen.ToGoIdentifier("CONTEST_" + string(s.Key)),
		ConstValue:    strconv.QuoteToASCII(string(s.Key)),
		ConstComments: fmt.Sprintf("%-20s = %s", s.Key, s.Description),
		IsDeprecated:  bool(s.IsImportOnly),
	}
}

func (c SpecMapContainer) CodeGenRecords() map[codegen.CodeGenKey]codegen.CodeGenSpec {
	result := make(map[codegen.CodeGenKey]codegen.CodeGenSpec, len(c.Records))
	for k, v := range c.Records {
		v.Key = Contest(strings.ToUpper(string(v.Key)))
		result[k] = v
	}
	return result
}

func (c SpecMapContainer) CodeGenMetadata() codegen.CodeGenContainerMetadata {
	return codegen.CodeGenContainerMetadata{
		PackageName: "contest",
		DataType:    "Contest",
	}
}
