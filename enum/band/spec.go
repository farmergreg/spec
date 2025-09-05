package band

import (
	"fmt"
	"strconv"

	"github.com/hamradiolog-net/adif-spec/v6/internal/codegen"
	"github.com/hamradiolog-net/adif-spec/v6/spectype"
)

var (
	_ codegen.CodeGenContainer = SpecMapContainer{}
	_ codegen.CodeGenSpec      = Spec{}
)

// SpecMapContainer contains all Band specifications as defined by the ADIF Workgroup specification exports.
type SpecMapContainer struct {
	// Header  []string         `json:"Header"`
	Records map[Band]Spec `json:"Records"`
}

// Spec represents the specification for a single Band as defined by the ADIF Workgroup specification exports.
type Spec struct {
	// EnumerationName string           `json:"Enumeration Name"`
	IsImportOnly spectype.Boolean `json:"Import-only,omitempty"`
	// Comments     string           `json:"Comments,omitempty"`
	Key          Band `json:"Band"` // Band
	LowerFreqMHz MHz  `json:"Lower Freq (MHz)"`
	UpperFreqMHz MHz  `json:"Upper Freq (MHz)"`
}

func (s Spec) CodeGeneratorMetadata() codegen.CodeGeneratorMetadataForEnum {
	return codegen.CodeGeneratorMetadataForEnum{
		ConstName:     codegen.ToGoIdentifier("Band" + string(s.Key)),
		ConstValue:    strconv.QuoteToASCII(string(s.Key)),
		ConstComments: fmt.Sprintf("%-6s = %12.4f MHz to %12.4f MHz", s.Key, s.LowerFreqMHz, s.UpperFreqMHz),
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
		PackageName: "band",
		DataType:    "Band",
	}
}

// IsInBand returns true if the specified frequency is within the band specification.
func (s *Spec) IsInBand(mhz float64) bool {
	return mhz >= float64(s.LowerFreqMHz) && mhz <= float64(s.UpperFreqMHz)
}

// FindBandByMHz returns the band specification that contains the given MHz value, if any.
func FindBandByMHz(mhz float64) (Spec, bool) {
	for _, item := range BandListAll {
		if item.IsInBand(mhz) {
			return item, true
		}
	}
	return Spec{}, false // Not found.
}

// Bandwidth returns the width of the frequency range in MHz
func (s *Spec) Bandwidth() float64 {
	return float64(s.UpperFreqMHz) - float64(s.LowerFreqMHz)
}
