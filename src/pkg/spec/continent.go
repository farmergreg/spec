package spec

import "github.com/hamradiolog-net/adif-spec/src/pkg/enum/continent"

// ContinentSpec represents the specification for a single Continent
type ContinentSpec struct {
	EnumerationName string          `json:"Enumeration Name"`
	IsImportOnly    AdifSpecBoolean `json:"Import-only,omitempty"`
	Comments        string          `json:"Comments,omitempty"`
	Id              string          `json:"Abbreviation"` // Abbreviation
	Continent       string          `json:"Continent"`
}

func (c ContinentSpec) Description() string {
	return c.Continent
}

// ContinentSpecMap contains all ContinentSpec specifications.
type ContinentSpecMap struct {
	Header  []string                              `json:"Header"`
	Records map[continent.Continent]ContinentSpec `json:"Records"`
}
