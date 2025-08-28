package spec

import (
	"fmt"
)

// CreditRecord represents a single credit record
type CreditRecord struct {
	BaseEnumerationRecord
	Id      string `json:"Credit For"` // Credit For
	Sponsor string `json:"Sponsor"`
	Award   string `json:"Award"`
	Facet   string `json:"Facet"`
}

func (c CreditRecord) Description() string {
	return fmt.Sprintf("%s %s %s", c.Sponsor, c.Award, c.Facet)
}

// CreditEnumeration represents the complete credit enumeration
type CreditEnumeration struct {
	Header  []string                `json:"Header"`
	Records map[string]CreditRecord `json:"Records"`
}
