package spec

import (
	"fmt"
)

// CreditSpec represents the specification for a single Credit
type CreditSpec struct {
	BaseEnumerationSpec
	Id      string `json:"Credit For"` // Credit For
	Sponsor string `json:"Sponsor"`
	Award   string `json:"Award"`
	Facet   string `json:"Facet"`
}

func (c CreditSpec) Description() string {
	return fmt.Sprintf("%s %s %s", c.Sponsor, c.Award, c.Facet)
}

// CreditSpecMap contains all CreditSpec specifications.
type CreditSpecMap struct {
	Header  []string              `json:"Header"`
	Records map[string]CreditSpec `json:"Records"`
}
