package spec

// PrimaryAdministrativeSubdivisionRecord represents a single primary administrative subdivision record
type PrimaryAdministrativeSubdivisionRecord struct {
	BaseEnumerationRecord
	Code            string          `json:"Code"` // Code
	PrimaryAdminSub string          `json:"Primary Administrative Subdivision"`
	DXCCEntityCode  string          `json:"DXCC Entity Code"`
	ContainedWithin string          `json:"Contained Within,omitempty"`
	Oblast          string          `json:"Oblast #,omitempty"`
	CQZone          string          `json:"CQ Zone,omitempty"`
	ITUZone         string          `json:"ITU Zone,omitempty"`
	Prefix          string          `json:"Prefix,omitempty"`
	IsDeleted       AdifSpecBoolean `json:"Deleted,omitempty"`
}

func (p PrimaryAdministrativeSubdivisionRecord) Description() string {
	return p.PrimaryAdminSub
}

// PrimaryAdministrativeSubdivisionEnumeration represents the complete primary administrative subdivision enumeration
type PrimaryAdministrativeSubdivisionEnumeration struct {
	Header  []string                                          `json:"Header"`
	Records map[string]PrimaryAdministrativeSubdivisionRecord `json:"Records"`
}
