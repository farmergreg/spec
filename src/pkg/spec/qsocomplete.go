package spec

// QSOCompleteRecord represents a single QSO complete record
type QSOCompleteRecord struct {
	BaseEnumerationSpec
	Id          string `json:"Abbreviation"` // Abbreviation
	Description string `json:"Meaning"`
}

// QSOCompleteEnumeration represents the complete QSO complete enumeration
type QSOCompleteEnumeration struct {
	Header  []string                     `json:"Header"`
	Records map[string]QSOCompleteRecord `json:"Records"`
}
