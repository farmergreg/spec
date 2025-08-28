package spec

// Field represents a single field definition
type Field struct {
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

// Fields represents the complete fields collection
type Fields struct {
	Header  []string         `json:"Header"`
	Records map[string]Field `json:"Records"`
}
