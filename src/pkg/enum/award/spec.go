package award

import (
	"github.com/hamradiolog-net/adif-spec/v2/src/pkg/spectype"
)

// SpecMap contains all Award specifications as defined by the ADIF Workgroup specification exports.
type SpecMapContainer struct {
	// Header  []string         `json:"Header"`
	Records map[Award]Spec `json:"Records"`
}

// Spec represents the specification for a single Award as defined by the ADIF Workgroup specification exports.
type Spec struct {
	// EnumerationName string           `json:"Enumeration Name"`
	IsImportOnly spectype.Boolean `json:"Import-only,omitempty"`
	Comments     string           `json:"Comments,omitempty"`
	Key          Award            `json:"Award"` // Award
}

func (a Spec) String() string {
	return string(a.Key)
}
