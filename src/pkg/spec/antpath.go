package spec

// AntPathRecord represents a single antenna path record
type AntPathRecord struct {
	BaseEnumerationRecord
	Id          string `json:"Abbreviation"` // Abbreviation
	Description string `json:"Meaning"`
}

// AntPathEnumeration represents the complete antenna path enumeration
type AntPathEnumeration struct {
	Header  []string                 `json:"Header"`
	Records map[string]AntPathRecord `json:"Records"`
}
