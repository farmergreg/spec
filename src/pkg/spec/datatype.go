package spec

// DataTypeSpec represents the specification for a single DataType
type DataTypeSpec struct {
	Id                string          `json:"Data Type Name"` // Data Type Name
	DataTypeIndicator string          `json:"Data Type Indicator,omitempty"`
	Description       string          `json:"Description"`
	MinimumValue      AdifSpecInteger `json:"Minimum Value,omitempty"`
	MaximumValue      AdifSpecInteger `json:"Maximum Value,omitempty"`
	IsImportOnly      AdifSpecBoolean `json:"Import-only,omitempty"`
	Comments          string          `json:"Comments,omitempty"`
}

// DataTypeSpecMap contains all DataTypeSpec specifications.
type DataTypeSpecMap struct {
	Header  []string                `json:"Header"`
	Records map[string]DataTypeSpec `json:"Records"`
}
