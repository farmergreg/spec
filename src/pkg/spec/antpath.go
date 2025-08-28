package spec

// AntPathSpec represents the specification for AntPath
type AntPathSpec struct {
	BaseEnumerationSpec
	Id          string `json:"Abbreviation"` // Abbreviation
	Description string `json:"Meaning"`
}

// AntPathEnumerationSpec represents list of all available AntPath specifications.
type AntPathEnumerationSpec struct {
	Header  []string               `json:"Header"`
	Records map[string]AntPathSpec `json:"Records"`
}
