package credit

import (
	"fmt"

	"github.com/hamradiolog-net/adif-spec/v6/spectype"
)

// SpecMapContainer contains all Credit specifications as defined by the ADIF Workgroup specification exports.
type SpecMapContainer struct {
	// Header  []string         `json:"Header"`
	Records map[Credit]Spec `json:"Records"`
}

// Spec represents the specification for a single Credit as defined by the ADIF Workgroup specification exports.
type Spec struct {
	// EnumerationName string           `json:"Enumeration Name"`
	IsImportOnly spectype.Boolean `json:"Import-only,omitempty"`
	// Comments     string           `json:"Comments,omitempty"`
	Key     Credit `json:"Credit For"` // Credit For
	Sponsor string `json:"Sponsor"`
	Award   string `json:"Award"`
	Facet   string `json:"Facet"`
}

func (c Spec) String() string {
	return fmt.Sprintf("%-20s = %-15s %-45s %-15s", c.Key, c.Sponsor, c.Award, c.Facet)
}
