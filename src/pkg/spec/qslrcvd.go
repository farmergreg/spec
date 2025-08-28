package spec

// QSLRcvdSpec represents the specification for a single QSLRcvd
type QSLRcvdSpec struct {
	BaseEnumerationSpec
	Id          string `json:"Status"` // Status
	Meaning     string `json:"Meaning"`
	Description string `json:"Description"`
}

// QSLRcvdSpecMap contains all QSLRcvdSpec specifications.
type QSLRcvdSpecMap struct {
	Header  []string               `json:"Header"`
	Records map[string]QSLRcvdSpec `json:"Records"`
}
