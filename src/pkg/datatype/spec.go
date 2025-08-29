package datatype

import "github.com/hamradiolog-net/adif-spec/src/pkg/spectype"

// Spec represents the specification for a single DataType as defined by the ADIF Workgroup specification exports.
type Spec struct {
	Id                string                   `json:"Data Type Name"` // Data Type Name
	DataTypeIndicator string                   `json:"Data Type Indicator,omitempty"`
	Description       string                   `json:"Description"`
	MinimumValue      spectype.AdifSpecInteger `json:"Minimum Value,omitempty"`
	MaximumValue      spectype.AdifSpecInteger `json:"Maximum Value,omitempty"`
	IsImportOnly      spectype.AdifSpecBoolean `json:"Import-only,omitempty"`
	Comments          string                   `json:"Comments,omitempty"`
}

// SpecMap contains all DataType specifications as defined by the ADIF Workgroup specification exports.
type SpecMap struct {
	Header  []string        `json:"Header"`
	Records map[string]Spec `json:"Records"`
}
