package spec

// EQSLAGSpec represents the specification for a single EQSLAG
type EQSLAGSpec struct {
	BaseEnumerationSpec
	Id          string `json:"Status"` // Status
	Description string `json:"Description"`
}

// EQSLAGSpecMap contains all EQSLAGSpec specifications.
type EQSLAGSpecMap struct {
	Header  []string              `json:"Header"`
	Records map[string]EQSLAGSpec `json:"Records"`
}
