package morsekeytype

import (
	"github.com/hamradiolog-net/adif-spec/src/pkg/spectype"
)

// MorseKeyTypeSpec represents the specification for a single MorseKeyType
type MorseKeyTypeSpec struct {
	EnumerationName  string                   `json:"Enumeration Name"`
	IsImportOnly     spectype.AdifSpecBoolean `json:"Import-only,omitempty"`
	Comments         string                   `json:"Comments,omitempty"`
	Id               string                   `json:"Abbreviation"` // Abbreviation
	Description      string                   `json:"Meaning"`
	Characteristics  string                   `json:"Characteristics,omitempty"`
	MorseComposition string                   `json:"Morse Composition,omitempty"`
	Examples         string                   `json:"Examples,omitempty"`
}

// SpecMap contains all MorseKeyTypeSpec specifications.
type SpecMap struct {
	Header  []string                          `json:"Header"`
	Records map[MorseKeyType]MorseKeyTypeSpec `json:"Records"`
}
