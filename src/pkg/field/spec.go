package field

import "github.com/hamradiolog-net/adif-spec/src/pkg/spectype"

// FieldSpec represents the specification for a single Field
type FieldSpec struct {
	Id            string                   `json:"Field Name"` // Field Name
	DataType      string                   `json:"Data Type"`
	Enumeration   string                   `json:"Enumeration,omitempty"`
	Description   string                   `json:"Description"`
	IsHeaderField string                   `json:"Header Field,omitempty"`
	MinimumValue  spectype.AdifSpecInteger `json:"Minimum Value,omitempty"`
	MaximumValue  spectype.AdifSpecInteger `json:"Maximum Value,omitempty"`
	IsImportOnly  spectype.AdifSpecBoolean `json:"Import-only,omitempty"`
	Comments      string                   `json:"Comments,omitempty"`
}

// SpecMap contains all FieldSpec specifications.
type SpecMap struct {
	Header  []string             `json:"Header"`
	Records map[string]FieldSpec `json:"Records"`
}
