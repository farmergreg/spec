package spec

// AwardRecord represents a single award record
type AwardRecord struct {
	BaseEnumerationRecord
	Id string `json:"Award"` // Award
}

func (a AwardRecord) Description() string {
	return a.Id
}

// AwardEnumeration represents the complete award enumeration
type AwardEnumeration struct {
	Header  []string               `json:"Header"`
	Records map[string]AwardRecord `json:"Records"`
}
