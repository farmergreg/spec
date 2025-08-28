package spec

import (
	"fmt"

	"github.com/hamradiolog-net/adif-spec/src/pkg/enum/band"
)

// BandRecord represents a single band record
type BandRecord struct {
	BaseEnumerationRecord
	Id           string   `json:"Band"` // Band
	LowerFreqMHz band.MHz `json:"Lower Freq (MHz)"`
	UpperFreqMHz band.MHz `json:"Upper Freq (MHz)"`
}

func (b BandRecord) Description() string {
	return fmt.Sprintf("%-6s %12.4f MHz - %12.4f MHz", b.Id, b.LowerFreqMHz, b.UpperFreqMHz)
}

// BandEnumeration represents the complete band enumeration
type BandEnumeration struct {
	Header  []string              `json:"Header"`
	Records map[string]BandRecord `json:"Records"`
}
