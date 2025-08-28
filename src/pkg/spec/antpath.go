package spec

import "github.com/hamradiolog-net/adif-spec/src/pkg/enum/antpath"

// AntPathSpec represents the specification for a single AntPath
type AntPathSpec struct {
	BaseEnumerationSpec
	Id          string `json:"Abbreviation"` // Abbreviation
	Description string `json:"Meaning"`
}

// AntPathEnumerationSpecMap contains all AntPathSpec specifications
type AntPathEnumerationSpecMap struct {
	Header  []string                        `json:"Header"`
	Records map[antpath.AntPath]AntPathSpec `json:"Records"`
}
