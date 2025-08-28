package spec

// SecondaryAdministrativeSubdivisionAltSpec represents the specification for a single SecondaryAdministrativeSubdivisionAlt
type SecondaryAdministrativeSubdivisionAltSpec struct {
	BaseEnumerationSpec
	Id             string          `json:"Code"` // Code
	DXCCEntityCode string          `json:"DXCC Entity Code"`
	Region         string          `json:"Region"`
	District       string          `json:"District"`
	IsDeleted      AdifSpecBoolean `json:"Deleted,omitempty"`
}

func (s SecondaryAdministrativeSubdivisionAltSpec) Description() string {
	return s.Region
}

// SecondaryAdministrativeSubdivisionAltSpecMap contains all SecondaryAdministrativeSubdivisionAltSpec specifications.
type SecondaryAdministrativeSubdivisionAltSpecMap struct {
	Header  []string                                             `json:"Header"`
	Records map[string]SecondaryAdministrativeSubdivisionAltSpec `json:"Records"`
}
