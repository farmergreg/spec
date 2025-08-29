package continent

import (
	"github.com/hamradiolog-net/adif-spec/src/pkg/spectype"
)

// Spec represents the specification for a single Continent as defined by the ADIF Workgroup specification exports.
type Spec struct {
	EnumerationName string           `json:"Enumeration Name"`
	IsImportOnly    spectype.Boolean `json:"Import-only,omitempty"`
	Comments        string           `json:"Comments,omitempty"`
	Id              string           `json:"Abbreviation"` // Abbreviation
	Continent       string           `json:"Continent"`
}

func (c Spec) Description() string {
	return c.Continent
}

// SpecMap contains all Continent specifications as defined by the ADIF Workgroup specification exports.
type SpecMap struct {
	Header  []string           `json:"Header"`
	Records map[Continent]Spec `json:"Records"`
}
