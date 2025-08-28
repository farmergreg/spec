package spec

// QSOCompleteSpec represents the specification for a single QSOComplete
type QSOCompleteSpec struct {
	BaseEnumerationSpec
	Id          string `json:"Abbreviation"` // Abbreviation
	Description string `json:"Meaning"`
}

// QSOCompleteSpecMap contains all QSOCompleteSpec specifications.
type QSOCompleteSpecMap struct {
	Header  []string                   `json:"Header"`
	Records map[string]QSOCompleteSpec `json:"Records"`
}
