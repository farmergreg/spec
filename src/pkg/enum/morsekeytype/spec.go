package morsekeytype

import (
	"github.com/hamradiolog-net/adif-spec/src/pkg/spectype"
)

// Spec represents the specification for a single MorseKeyType as defined by the ADIF Workgroup specification exports.
type Spec struct {
	EnumerationName  string           `json:"Enumeration Name"`
	IsImportOnly     spectype.Boolean `json:"Import-only,omitempty"`
	Comments         string           `json:"Comments,omitempty"`
	Id               string           `json:"Abbreviation"` // Abbreviation
	Description      string           `json:"Meaning"`
	Characteristics  string           `json:"Characteristics,omitempty"`
	MorseComposition string           `json:"Morse Composition,omitempty"`
	Examples         string           `json:"Examples,omitempty"`
}

// SpecMap contains all MorseKeyType specifications as defined by the ADIF Workgroup specification exports.
type SpecMap struct {
	Header  []string              `json:"Header"`
	Records map[MorseKeyType]Spec `json:"Records"`
}
