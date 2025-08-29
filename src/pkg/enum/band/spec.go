package band

import (
	"fmt"

	"github.com/hamradiolog-net/adif-spec/src/pkg/spectype"
)

// Spec represents the specification for a single Band as defined by the ADIF Workgroup specification exports.
type Spec struct {
	EnumerationName string           `json:"Enumeration Name"`
	IsImportOnly    spectype.Boolean `json:"Import-only,omitempty"`
	Comments        string           `json:"Comments,omitempty"`
	Id              string           `json:"Band"` // Band
	LowerFreqMHz    MHz              `json:"Lower Freq (MHz)"`
	UpperFreqMHz    MHz              `json:"Upper Freq (MHz)"`
}

func (b Spec) Description() string {
	return fmt.Sprintf("%-6s %12.4f MHz - %12.4f MHz", b.Id, b.LowerFreqMHz, b.UpperFreqMHz)
}

// SpecMap contains all Band specifications as defined by the ADIF Workgroup specification exports.
type SpecMap struct {
	Header  []string      `json:"Header"`
	Records map[Band]Spec `json:"Records"`
}
