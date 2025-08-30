package aditype

import "github.com/hamradiolog-net/adif-spec/v2/src/pkg/spectype"

// SpecMapContainer contains all DataType specifications as defined by the ADIF Workgroup specification exports.
type SpecMapContainer struct {
	Header  []string         `json:"Header"`
	Records map[ADIType]Spec `json:"Records"`
}

// Spec represents the specification for a single DataType as defined by the ADIF Workgroup specification exports.
type Spec struct {
	Key               ADIType          `json:"Data Type Name"` // Data Type Name
	DataTypeIndicator string           `json:"Data Type Indicator,omitempty"`
	Description       string           `json:"Description"`
	MinimumValue      spectype.Integer `json:"Minimum Value,omitempty"`
	MaximumValue      spectype.Integer `json:"Maximum Value,omitempty"`
	IsImportOnly      spectype.Boolean `json:"Import-only,omitempty"`
	Comments          string           `json:"Comments,omitempty"`
}

func (s Spec) String() string {
	return s.Description
}
