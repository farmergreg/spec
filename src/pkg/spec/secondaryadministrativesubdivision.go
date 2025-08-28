package spec

// SecondaryAdministrativeSubdivisionSpec represents the specification for a single SecondaryAdministrativeSubdivision
type SecondaryAdministrativeSubdivisionSpec struct {
	BaseEnumerationSpec
	Id                     string          `json:"Code"` // Code
	SecondaryAdminSub      string          `json:"Secondary Administrative Subdivision"`
	DXCCEntityCode         string          `json:"DXCC Entity Code"`
	AlaskaJudicialDistrict string          `json:"Alaska Judicial District,omitempty"`
	IsDeleted              AdifSpecBoolean `json:"Deleted,omitempty"`
}

func (s SecondaryAdministrativeSubdivisionSpec) Description() string {
	return s.SecondaryAdminSub
}

// SecondaryAdministrativeSubdivisionSpecMap contains all SecondaryAdministrativeSubdivisionSpec specifications.
type SecondaryAdministrativeSubdivisionSpecMap struct {
	Header  []string                                          `json:"Header"`
	Records map[string]SecondaryAdministrativeSubdivisionSpec `json:"Records"`
}
