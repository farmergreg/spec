package spec

// SecondaryAdministrativeSubdivisionAltRecord represents a single secondary administrative subdivision alt record
type SecondaryAdministrativeSubdivisionAltRecord struct {
	BaseEnumerationSpec
	Id             string          `json:"Code"` // Code
	DXCCEntityCode string          `json:"DXCC Entity Code"`
	Region         string          `json:"Region"`
	District       string          `json:"District"`
	IsDeleted      AdifSpecBoolean `json:"Deleted,omitempty"`
}

func (s SecondaryAdministrativeSubdivisionAltRecord) Description() string {
	return s.Region
}

// SecondaryAdministrativeSubdivisionAltEnumeration represents the complete secondary administrative subdivision alt enumeration
type SecondaryAdministrativeSubdivisionAltEnumeration struct {
	Header  []string                                               `json:"Header"`
	Records map[string]SecondaryAdministrativeSubdivisionAltRecord `json:"Records"`
}
