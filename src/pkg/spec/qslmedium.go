package spec

// QSLMediumRecord represents a single QSL medium record
type QSLMediumRecord struct {
	BaseEnumerationSpec
	Id          string `json:"Medium"` // Medium
	Description string `json:"Description"`
}

// QSLMediumEnumeration represents the complete QSL medium enumeration
type QSLMediumEnumeration struct {
	Header  []string                   `json:"Header"`
	Records map[string]QSLMediumRecord `json:"Records"`
}
