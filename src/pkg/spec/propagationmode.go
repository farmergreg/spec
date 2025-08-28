package spec

// PropagationModeSpec represents the specification for a single PropagationMode
type PropagationModeSpec struct {
	BaseEnumerationSpec
	Id          string `json:"Enumeration"` // Enumeration
	Description string `json:"Description"`
}

// PropagationModeSpecMap contains all PropagationModeSpec specifications.
type PropagationModeSpecMap struct {
	Header  []string                       `json:"Header"`
	Records map[string]PropagationModeSpec `json:"Records"`
}
