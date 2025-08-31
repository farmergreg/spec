package qsocomplete

import (
	"fmt"

	"github.com/hamradiolog-net/adif-spec/v6/spectype"
)

// SpecMapContainer contains all QSOComplete specifications as defined by the ADIF Workgroup specification exports.
type SpecMapContainer struct {
	// Header  []string         `json:"Header"`
	Records map[QSOComplete]Spec `json:"Records"`
}

// Spec represents the specification for a single QSOComplete as defined by the ADIF Workgroup specification exports.
type Spec struct {
	// EnumerationName string           `json:"Enumeration Name"`
	IsImportOnly spectype.Boolean `json:"Import-only,omitempty"`
	// Comments     string           `json:"Comments,omitempty"`
	Key         QSOComplete `json:"Abbreviation"` // Abbreviation
	Description string      `json:"Meaning"`
}

func (s Spec) String() string {
	return fmt.Sprintf("%-4s = %s", s.Key, s.Description)
}
