package credit

import (
	"fmt"

	"github.com/hamradiolog-net/adif-spec/src/pkg/spectype"
)

// Spec represents the specification for a single Credit as defined by the ADIF Workgroup specification exports.
type Spec struct {
	EnumerationName string                   `json:"Enumeration Name"`
	IsImportOnly    spectype.AdifSpecBoolean `json:"Import-only,omitempty"`
	Comments        string                   `json:"Comments,omitempty"`
	Id              string                   `json:"Credit For"` // Credit For
	Sponsor         string                   `json:"Sponsor"`
	Award           string                   `json:"Award"`
	Facet           string                   `json:"Facet"`
}

func (c Spec) Description() string {
	return fmt.Sprintf("%s %s %s", c.Sponsor, c.Award, c.Facet)
}

// SpecMap contains all Credit specifications as defined by the ADIF Workgroup specification exports.
type SpecMap struct {
	Header  []string        `json:"Header"`
	Records map[Credit]Spec `json:"Records"`
}
