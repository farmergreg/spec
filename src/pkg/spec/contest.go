package spec

import (
	"github.com/hamradiolog-net/adif-spec/src/pkg/enum/contest"
	"github.com/hamradiolog-net/adif-spec/src/pkg/spectype"
)

// ContestSpec represents the specification for a single Contest
type ContestSpec struct {
	EnumerationName string                   `json:"Enumeration Name"`
	IsImportOnly    spectype.AdifSpecBoolean `json:"Import-only,omitempty"`
	Comments        string                   `json:"Comments,omitempty"`
	Id              string                   `json:"Contest-ID"` // Contest ID
	Description     string                   `json:"Description"`
}

// ContestSpecMap contains all ContestSpec specifications.
type ContestSpecMap struct {
	Header  []string                        `json:"Header"`
	Records map[contest.Contest]ContestSpec `json:"Records"`
}
