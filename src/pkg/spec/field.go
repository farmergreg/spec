package spec

import "github.com/hamradiolog-net/adif-spec/v2/src/pkg/spectype"

// FieldSpec represents the specification for a single Field
type FieldSpec struct {
	Id            string           `json:"Field Name"` // Field Name
	DataType      string           `json:"Data Type"`
	Enumeration   string           `json:"Enumeration,omitempty"`
	Description   string           `json:"Description"`
	IsHeaderField string           `json:"Header Field,omitempty"`
	MinimumValue  spectype.Integer `json:"Minimum Value,omitempty"`
	MaximumValue  spectype.Integer `json:"Maximum Value,omitempty"`
	IsImportOnly  spectype.Boolean `json:"Import-only,omitempty"`
	Comments      string           `json:"Comments,omitempty"`
}

// FieldSpecMap contains all FieldSpec specifications.
type FieldSpecMap struct {
	// Header  []string         `json:"Header"`
	Records map[string]FieldSpec `json:"Records"`
}
