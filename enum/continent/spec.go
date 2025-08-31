package continent

import (
	"fmt"

	"github.com/hamradiolog-net/adif-spec/v6/spectype"
)

// SpecMapContainer contains all Continent specifications as defined by the ADIF Workgroup specification exports.
type SpecMapContainer struct {
	// Header  []string         `json:"Header"`
	Records map[Continent]Spec `json:"Records"`
}

// Spec represents the specification for a single Continent as defined by the ADIF Workgroup specification exports.
type Spec struct {
	// EnumerationName string           `json:"Enumeration Name"`
	IsImportOnly spectype.Boolean `json:"Import-only,omitempty"`
	// Comments     string           `json:"Comments,omitempty"`
	Key       Continent `json:"Abbreviation"` // Abbreviation
	Continent string    `json:"Continent"`
}

func (s Spec) String() string {
	return fmt.Sprintf("%s = %s", s.Key, s.Continent)
}
