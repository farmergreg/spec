package morsekeytype

import (
	"fmt"

	"github.com/hamradiolog-net/adif-spec/v2/src/pkg/spectype"
)

// SpecMap contains all MorseKeyType specifications as defined by the ADIF Workgroup specification exports.
type SpecMapContainer struct {
	// Header  []string         `json:"Header"`
	Records map[MorseKeyType]Spec `json:"Records"`
}

// Spec represents the specification for a single MorseKeyType as defined by the ADIF Workgroup specification exports.
type Spec struct {
	// EnumerationName string           `json:"Enumeration Name"`
	IsImportOnly     spectype.Boolean `json:"Import-only,omitempty"`
	Comments         string           `json:"Comments,omitempty"`
	Key              MorseKeyType     `json:"Abbreviation"` // Abbreviation
	Description      string           `json:"Meaning"`
	Characteristics  string           `json:"Characteristics,omitempty"`
	MorseComposition string           `json:"Morse Composition,omitempty"`
	Examples         string           `json:"Examples,omitempty"`
}

func (s Spec) String() string {
	return fmt.Sprintf("%-4s = %s", s.Key, s.Description)
}
