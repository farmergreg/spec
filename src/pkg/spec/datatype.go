package spec

// DataType represents a single data type definition
type DataType struct {
	Id                string          `json:"Data Type Name"` // Data Type Name
	DataTypeIndicator string          `json:"Data Type Indicator,omitempty"`
	Description       string          `json:"Description"`
	MinimumValue      AdifSpecInteger `json:"Minimum Value,omitempty"`
	MaximumValue      AdifSpecInteger `json:"Maximum Value,omitempty"`
	IsImportOnly      AdifSpecBoolean `json:"Import-only,omitempty"`
	Comments          string          `json:"Comments,omitempty"`
}

// DataTypes represents the complete data types collection
type DataTypes struct {
	Header  []string            `json:"Header"`
	Records map[string]DataType `json:"Records"`
}
