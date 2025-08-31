package adifield

import (
	"fmt"

	"github.com/hamradiolog-net/adif-spec/v6/spectype"
)

// SpecMapContainer contains all Field specifications as defined by the ADIF Workgroup specification exports.
type SpecMapContainer struct {
	// Header  []string         `json:"Header"`
	Records map[ADIField]Spec `json:"Records"`
}

// Spec represents the specification for a single Field as defined by the ADIF Workgroup specification exports.
type Spec struct {
	Key      ADIField `json:"Field Name"` // Field Name
	DataType string   `json:"Data Type"`
	// Enumeration   string           `json:"Enumeration,omitempty"`
	Description   string           `json:"Description"`
	IsHeaderField spectype.Boolean `json:"Header Field,omitempty"`
	MinimumValue  spectype.Integer `json:"Minimum Value,omitempty"`
	MaximumValue  spectype.Integer `json:"Maximum Value,omitempty"`
	IsImportOnly  spectype.Boolean `json:"Import-only,omitempty"`
	Comments      string           `json:"Comments,omitempty"`
}

func (s Spec) String() string {
	if s.IsHeaderField {
		return fmt.Sprintf("Header: %s", s.Description)
	} else {
		return fmt.Sprintf("Record: %s", s.Description)
	}
}
