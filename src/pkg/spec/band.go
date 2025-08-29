package spec

import (
	"fmt"

	"github.com/hamradiolog-net/adif-spec/src/pkg/enum/band"
	"github.com/hamradiolog-net/adif-spec/src/pkg/spectype"
)

// BandSpec represents the specification for a single Band
type BandSpec struct {
	EnumerationName string                   `json:"Enumeration Name"`
	IsImportOnly    spectype.AdifSpecBoolean `json:"Import-only,omitempty"`
	Comments        string                   `json:"Comments,omitempty"`
	Id              string                   `json:"Band"` // Band
	LowerFreqMHz    band.MHz                 `json:"Lower Freq (MHz)"`
	UpperFreqMHz    band.MHz                 `json:"Upper Freq (MHz)"`
}

func (b BandSpec) Description() string {
	return fmt.Sprintf("%-6s %12.4f MHz - %12.4f MHz", b.Id, b.LowerFreqMHz, b.UpperFreqMHz)
}

// BandSpecMap contains all BandSpec specifications.
type BandSpecMap struct {
	Header  []string               `json:"Header"`
	Records map[band.Band]BandSpec `json:"Records"`
}
