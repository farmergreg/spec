package spec

// EQSLAGRecord represents a single EQSL AG record
type EQSLAGRecord struct {
	BaseEnumerationSpec
	Id          string `json:"Status"` // Status
	Description string `json:"Description"`
}

// EQSLAGEnumeration represents the complete EQSL AG enumeration
type EQSLAGEnumeration struct {
	Header  []string                `json:"Header"`
	Records map[string]EQSLAGRecord `json:"Records"`
}
