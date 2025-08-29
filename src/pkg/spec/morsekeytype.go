package spec

import (
	"github.com/hamradiolog-net/adif-spec/src/pkg/enum/morsekeytype"
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

// MorseKeyTypeSpecMap contains all MorseKeyTypeSpec specifications.
type MorseKeyTypeSpecMap struct {
	Header  []string                                       `json:"Header"`
	Records map[morsekeytype.MorseKeyType]MorseKeyTypeSpec `json:"Records"`
}
