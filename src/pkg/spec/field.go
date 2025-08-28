package spec

// FieldSpec represents the specification for a single Field
type FieldSpec struct {
	Id            string          `json:"Field Name"` // Field Name
	DataType      string          `json:"Data Type"`
	Enumeration   string          `json:"Enumeration,omitempty"`
	Description   string          `json:"Description"`
	IsHeaderField string          `json:"Header Field,omitempty"`
	MinimumValue  AdifSpecInteger `json:"Minimum Value,omitempty"`
	MaximumValue  AdifSpecInteger `json:"Maximum Value,omitempty"`
	IsImportOnly  AdifSpecBoolean `json:"Import-only,omitempty"`
	Comments      string          `json:"Comments,omitempty"`
}

// FieldSpecMap contains all FieldSpec specifications.
type FieldSpecMap struct {
	Header  []string             `json:"Header"`
	Records map[string]FieldSpec `json:"Records"`
}
