package spec

// PropagationModeRecord represents a single propagation mode record
type PropagationModeRecord struct {
	BaseEnumerationSpec
	Id          string `json:"Enumeration"` // Enumeration
	Description string `json:"Description"`
}

// PropagationModeEnumeration represents the complete propagation mode enumeration
type PropagationModeEnumeration struct {
	Header  []string                         `json:"Header"`
	Records map[string]PropagationModeRecord `json:"Records"`
}
