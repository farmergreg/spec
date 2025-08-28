package spec

// QSLSentRecord represents a single QSL sent record
type QSLSentRecord struct {
	BaseEnumerationRecord
	Id          string `json:"Status"` // Status
	Meaning     string `json:"Meaning"`
	Description string `json:"Description"`
}

// QSLSentEnumeration represents the complete QSL sent enumeration
type QSLSentEnumeration struct {
	Header  []string                 `json:"Header"`
	Records map[string]QSLSentRecord `json:"Records"`
}
