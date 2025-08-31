package qslrcvd

import (
	"fmt"

	"github.com/hamradiolog-net/adif-spec/v6/spectype"
)

// SpecMap contains all QSLRcvd specifications as defined by the ADIF Workgroup specification exports.
type SpecMapContainer struct {
	// Header  []string         `json:"Header"`
	Records map[QSLRcvd]Spec `json:"Records"`
}

// Spec represents the specification for a single QSLRcvd as defined by the ADIF Workgroup specification exports.
type Spec struct {
	// EnumerationName string           `json:"Enumeration Name"`
	IsImportOnly spectype.Boolean `json:"Import-only,omitempty"`
	// Comments     string           `json:"Comments,omitempty"`
	Key         QSLRcvd `json:"Status"` // Status
	Meaning     string  `json:"Meaning"`
	Description string  `json:"Description"`
}

func (s Spec) String() string {
	return fmt.Sprintf("%s = %s", s.Key, s.Description)
}
