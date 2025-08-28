package spec

// SecondaryAdministrativeSubdivisionRecord represents a single secondary administrative subdivision record
type SecondaryAdministrativeSubdivisionRecord struct {
	BaseEnumerationRecord
	Id                     string          `json:"Code"` // Code
	SecondaryAdminSub      string          `json:"Secondary Administrative Subdivision"`
	DXCCEntityCode         string          `json:"DXCC Entity Code"`
	AlaskaJudicialDistrict string          `json:"Alaska Judicial District,omitempty"`
	IsDeleted              AdifSpecBoolean `json:"Deleted,omitempty"`
}

func (s SecondaryAdministrativeSubdivisionRecord) Description() string {
	return s.SecondaryAdminSub
}

// SecondaryAdministrativeSubdivisionEnumeration represents the complete secondary administrative subdivision enumeration
type SecondaryAdministrativeSubdivisionEnumeration struct {
	Header  []string                                            `json:"Header"`
	Records map[string]SecondaryAdministrativeSubdivisionRecord `json:"Records"`
}
