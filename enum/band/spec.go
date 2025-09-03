package band

import (
	"fmt"

	"github.com/hamradiolog-net/adif-spec/v6/spectype"
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
	Key          Band         `json:"Band"` // Band
	LowerFreqMHz spectype.MHz `json:"Lower Freq (MHz)"`
	UpperFreqMHz spectype.MHz `json:"Upper Freq (MHz)"`
}

// Depreciated: CodeGeneratorMetadata is not part of the stable API and may change without warning in the future even for minor version numbers.
func (s Spec) CodeGeneratorMetadata() string {
	return fmt.Sprintf("%-6s = %12.4f MHz to %12.4f MHz", s.Key, s.LowerFreqMHz, s.UpperFreqMHz)
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
