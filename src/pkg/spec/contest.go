package spec

// ContestSpec represents the specification for a single Contest
type ContestSpec struct {
	BaseEnumerationSpec
	Id          string `json:"Contest-ID"` // Contest ID
	Description string `json:"Description"`
}

// ContestSpecMap contains all ContestSpec specifications.
type ContestSpecMap struct {
	Header  []string               `json:"Header"`
	Records map[string]ContestSpec `json:"Records"`
}
