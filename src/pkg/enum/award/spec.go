package award

import (
	"github.com/hamradiolog-net/adif-spec/src/pkg/spectype"
)

// AwardSpec represents the specification for a single Award
type AwardSpec struct {
	EnumerationName string                   `json:"Enumeration Name"`
	IsImportOnly    spectype.AdifSpecBoolean `json:"Import-only,omitempty"`
	Comments        string                   `json:"Comments,omitempty"`
	Id              string                   `json:"Award"` // Award
}

func (a AwardSpec) Description() string {
	return a.Id
}

// SpecMap contains all AwardSpec specifications.
type SpecMap struct {
	Header  []string            `json:"Header"`
	Records map[Award]AwardSpec `json:"Records"`
}
