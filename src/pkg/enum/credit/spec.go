package credit

import (
	"fmt"

	"github.com/hamradiolog-net/adif-spec/src/pkg/spectype"
)

// CreditSpec represents the specification for a single Credit
type CreditSpec struct {
	EnumerationName string                   `json:"Enumeration Name"`
	IsImportOnly    spectype.AdifSpecBoolean `json:"Import-only,omitempty"`
	Comments        string                   `json:"Comments,omitempty"`
	Id              string                   `json:"Credit For"` // Credit For
	Sponsor         string                   `json:"Sponsor"`
	Award           string                   `json:"Award"`
	Facet           string                   `json:"Facet"`
}

func (c CreditSpec) Description() string {
	return fmt.Sprintf("%s %s %s", c.Sponsor, c.Award, c.Facet)
}

// CreditSpecMap contains all CreditSpec specifications.
type CreditSpecMap struct {
	Header  []string     `json:"Header"`
	Records map[Credit]CreditSpec `json:"Records"`
}
