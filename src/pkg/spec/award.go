package spec

// AwardSpec represents the specification for a single Award
type AwardSpec struct {
	BaseEnumerationSpec
	Id string `json:"Award"` // Award
}

func (a AwardSpec) Description() string {
	return a.Id
}

// AwardSpecMap contains all AwardSpec specifications.
type AwardSpecMap struct {
	Header  []string             `json:"Header"`
	Records map[string]AwardSpec `json:"Records"`
}
