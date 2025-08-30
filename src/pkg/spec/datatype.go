package spec

import "github.com/hamradiolog-net/adif-spec/v2/src/pkg/spectype"

// DataTypeSpec represents the specification for a single DataType
type DataTypeSpec struct {
	Id                string           `json:"Data Type Name"` // Data Type Name
	DataTypeIndicator string           `json:"Data Type Indicator,omitempty"`
	Description       string           `json:"Description"`
	MinimumValue      spectype.Integer `json:"Minimum Value,omitempty"`
	MaximumValue      spectype.Integer `json:"Maximum Value,omitempty"`
	IsImportOnly      spectype.Boolean `json:"Import-only,omitempty"`
	Comments          string           `json:"Comments,omitempty"`
}

// DataTypeSpecMap contains all DataTypeSpec specifications.
type DataTypeSpecMap struct {
	// Header  []string         `json:"Header"`
	Records map[string]DataTypeSpec `json:"Records"`
}
