package spec

// QSLRcvdRecord represents a single QSL received record
type QSLRcvdRecord struct {
	BaseEnumerationSpec
	Id          string `json:"Status"` // Status
	Meaning     string `json:"Meaning"`
	Description string `json:"Description"`
}

// QSLRcvdEnumeration represents the complete QSL received enumeration
type QSLRcvdEnumeration struct {
	Header  []string                 `json:"Header"`
	Records map[string]QSLRcvdRecord `json:"Records"`
}
