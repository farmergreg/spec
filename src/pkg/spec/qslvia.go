package spec

// QSLViaRecord represents a single QSL via record
type QSLViaRecord struct {
	BaseEnumerationRecord
	Id          string `json:"Via"` // Via
	Description string `json:"Description"`
}

// QSLViaEnumeration represents the complete QSL via enumeration
type QSLViaEnumeration struct {
	Header  []string                `json:"Header"`
	Records map[string]QSLViaRecord `json:"Records"`
}
