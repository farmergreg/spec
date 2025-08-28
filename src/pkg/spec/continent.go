package spec

// ContinentRecord represents a single continent record
type ContinentRecord struct {
	BaseEnumerationSpec
	Id        string `json:"Abbreviation"` // Abbreviation
	Continent string `json:"Continent"`
}

func (c ContinentRecord) Description() string {
	return c.Continent
}

// ContinentEnumeration represents the complete continent enumeration
type ContinentEnumeration struct {
	Header  []string                   `json:"Header"`
	Records map[string]ContinentRecord `json:"Records"`
}
