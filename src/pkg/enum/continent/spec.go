package continent

import (
	"github.com/hamradiolog-net/adif-spec/src/pkg/spectype"
)

// ContinentSpec represents the specification for a single Continent
type ContinentSpec struct {
	EnumerationName string                   `json:"Enumeration Name"`
	IsImportOnly    spectype.AdifSpecBoolean `json:"Import-only,omitempty"`
	Comments        string                   `json:"Comments,omitempty"`
	Id              string                   `json:"Abbreviation"` // Abbreviation
	Continent       string                   `json:"Continent"`
}

func (c ContinentSpec) Description() string {
	return c.Continent
}

// SpecMap contains all ContinentSpec specifications.
type SpecMap struct {
	Header  []string                    `json:"Header"`
	Records map[Continent]ContinentSpec `json:"Records"`
}
