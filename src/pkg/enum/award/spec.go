package award

import (
	"github.com/hamradiolog-net/adif-spec/src/pkg/spectype"
)

// Spec represents the specification for a single Award as defined by the ADIF Workgroup specification exports.
type Spec struct {
	EnumerationName string           `json:"Enumeration Name"`
	IsImportOnly    spectype.Boolean `json:"Import-only,omitempty"`
	Comments        string           `json:"Comments,omitempty"`
	Id              string           `json:"Award"` // Award
}

func (a Spec) Description() string {
	return a.Id
}

// SpecMap contains all Award specifications as defined by the ADIF Workgroup specification exports.
type SpecMap struct {
	Header  []string       `json:"Header"`
	Records map[Award]Spec `json:"Records"`
}
