package field

import "github.com/hamradiolog-net/adif-spec/src/pkg/spectype"

// Spec represents the specification for a single Field as defined by the ADIF Workgroup specification exports.
type Spec struct {
	Id            Field            `json:"Field Name"` // Field Name
	DataType      string           `json:"Data Type"`
	Enumeration   string           `json:"Enumeration,omitempty"`
	Description   string           `json:"Description"`
	IsHeaderField string           `json:"Header Field,omitempty"`
	MinimumValue  spectype.Integer `json:"Minimum Value,omitempty"`
	MaximumValue  spectype.Integer `json:"Maximum Value,omitempty"`
	IsImportOnly  spectype.Boolean `json:"Import-only,omitempty"`
	Comments      string           `json:"Comments,omitempty"`
}

// SpecMap contains all Field specifications as defined by the ADIF Workgroup specification exports.
type SpecMap struct {
	Header  []string       `json:"Header"`
	Records map[Field]Spec `json:"Records"`
}
