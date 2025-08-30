package antpath

import (
	"fmt"

	"github.com/hamradiolog-net/adif-spec/v2/src/pkg/spectype"
)

// SpecMap contains all AntPath specifications as defined by the ADIF Workgroup specification exports.
type SpecMapContainer struct {
	// Header  []string         `json:"Header"`
	Records map[AntPath]Spec `json:"Records"`
}

// Spec represents the specification for a single AntPath as defined by the ADIF Workgroup specification exports.
type Spec struct {
	// EnumerationName string           `json:"Enumeration Name"`
	IsImportOnly spectype.Boolean `json:"Import-only,omitempty"`
	Comments     string           `json:"Comments,omitempty"`
	Key          AntPath          `json:"Abbreviation"` // Abbreviation
	Description  string           `json:"Meaning"`
}

func (s Spec) String() string {
	return fmt.Sprintf("%s = %s", s.Key, s.Description)
}
