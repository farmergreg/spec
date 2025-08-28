package spec

// AntPathSpec represents the specification for AntPath
type AntPathSpec struct {
	BaseEnumerationSpec
	Id          string `json:"Abbreviation"` // Abbreviation
	Description string `json:"Meaning"`
}

// AntPathEnumeration represents list of all available AntPath specifications.
type AntPathEnumeration struct {
	Header  []string               `json:"Header"`
	Records map[string]AntPathSpec `json:"Records"`
}
