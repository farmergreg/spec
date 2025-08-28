package spec

// AntPathSpec represents the specification for a single AntPath
type AntPathSpec struct {
	BaseEnumerationSpec
	Id          string `json:"Abbreviation"` // Abbreviation
	Description string `json:"Meaning"`
}

// AntPathEnumerationSpecMap contains all AntPathSpec specifications
type AntPathEnumerationSpecMap struct {
	Header  []string               `json:"Header"`
	Records map[string]AntPathSpec `json:"Records"`
}
