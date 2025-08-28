// GENERATED CODE; DO NOT EDIT
// ADIF: 3.1.6 Proposed

package continent

var (
	AF Continent = "AF" // Africa
	AN Continent = "AN" // Antarctica
	AS Continent = "AS" // Asia
	EU Continent = "EU" // Europe
	NA Continent = "NA" // North America
	OC Continent = "OC" // Oceania
	SA Continent = "SA" // South America
)

// IsValid returns true if the Continent exists in the ADIF specification. This includes Import Only and Deleted values.
func (value Continent) IsValid() bool {
	switch value {
	case AF:
		return true
	case AN:
		return true
	case AS:
		return true
	case EU:
		return true
	case NA:
		return true
	case OC:
		return true
	case SA:
		return true
	}
	return false
}
