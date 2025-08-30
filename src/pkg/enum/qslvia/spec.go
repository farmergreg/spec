package qslvia

import (
	"fmt"

	"github.com/hamradiolog-net/adif-spec/src/pkg/spectype"
)

// SpecMap contains all QSLVia specifications as defined by the ADIF Workgroup specification exports.
type SpecMapContainer struct {
	// Header  []string         `json:"Header"`
	Records map[QSLVia]Spec `json:"Records"`
}

// Spec represents the specification for a single QSLVia as defined by the ADIF Workgroup specification exports.
type Spec struct {
	// EnumerationName string           `json:"Enumeration Name"`
	IsImportOnly spectype.Boolean `json:"Import-only,omitempty"`
	Comments     string           `json:"Comments,omitempty"`
	Key          QSLVia           `json:"Via"` // Via
	Description  string           `json:"Description"`
}

func (s Spec) String() string {
	return fmt.Sprintf("%s = %s", s.Key, s.Description)
}
