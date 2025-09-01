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
	Key          Band `json:"Band"` // Band
	LowerFreqMHz MHz  `json:"Lower Freq (MHz)"`
	UpperFreqMHz MHz  `json:"Upper Freq (MHz)"`
}

func (b Spec) String() string {
	return fmt.Sprintf("%-6s = %12.4f MHz to %12.4f MHz", b.Key, b.LowerFreqMHz, b.UpperFreqMHz)
}

// IsInBand returns true if the specified frequency is within the band specification.
func (b *Spec) IsInBand(mhz MHz) bool {
	return mhz >= b.LowerFreqMHz && mhz <= b.UpperFreqMHz
}

// FindBandByMHz returns the band specification that contains the given MHz value, if any.
func FindBandByMHz(mhz MHz) (Spec, bool) {
	for _, item := range BandListAll {
		if item.IsInBand(mhz) {
			return item, true
		}
	}
	return Spec{}, false // Not found.
}

// Bandwidth returns the width of the frequency range in MHz
func (b *Spec) Bandwidth() MHz {
	return b.UpperFreqMHz - b.LowerFreqMHz
}
