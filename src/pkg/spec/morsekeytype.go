package spec

// MorseKeyTypeSpec represents the specification for a single MorseKeyType
type MorseKeyTypeSpec struct {
	BaseEnumerationSpec
	Id               string `json:"Abbreviation"` // Abbreviation
	Description      string `json:"Meaning"`
	Characteristics  string `json:"Characteristics,omitempty"`
	MorseComposition string `json:"Morse Composition,omitempty"`
	Examples         string `json:"Examples,omitempty"`
}

// MorseKeyTypeSpecMap contains all MorseKeyTypeSpec specifications.
type MorseKeyTypeSpecMap struct {
	Header  []string                    `json:"Header"`
	Records map[string]MorseKeyTypeSpec `json:"Records"`
}
