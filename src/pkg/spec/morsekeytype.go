package spec

// MorseKeyTypeRecord represents a single morse key type record
type MorseKeyTypeRecord struct {
	BaseEnumerationSpec
	Id               string `json:"Abbreviation"` // Abbreviation
	Description      string `json:"Meaning"`
	Characteristics  string `json:"Characteristics,omitempty"`
	MorseComposition string `json:"Morse Composition,omitempty"`
	Examples         string `json:"Examples,omitempty"`
}

// MorseKeyTypeEnumeration represents the complete morse key type enumeration
type MorseKeyTypeEnumeration struct {
	Header  []string                      `json:"Header"`
	Records map[string]MorseKeyTypeRecord `json:"Records"`
}
