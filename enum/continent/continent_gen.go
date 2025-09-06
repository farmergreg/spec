// DO NOT EDIT; GENERATED CODE
// run `go generate ./...` from the project root to rebuild this file.

// Package continent provides code and constants as defined in ADIF 3.1.6 (Proposed)
package continent

const (
	AF Continent = "AF" // AF = Africa
	AN Continent = "AN" // AN = Antarctica
	AS Continent = "AS" // AS = Asia
	EU Continent = "EU" // EU = Europe
	NA Continent = "NA" // NA = North America
	OC Continent = "OC" // OC = Oceania
	SA Continent = "SA" // SA = South America
)

// Lookup look up a specification for Continent
func Lookup(continent Continent) (Spec, bool) {
	spec, ok := internalMap[continent]
	return spec, ok
}

// LookupByFilter returns all Continent specifications that match the provided filter function.
func LookupByFilter(filter func(Spec) bool) []Spec {
	result := make([]Spec, 0)
	for _, v := range List() {
		if filter(v) {
			result = append(result, v)
		}
	}
	return result
}

// All Continent specifications INCLUDING those marked import only.
func List() []Spec {
	return []Spec{
		internalMap[AF],
		internalMap[AN],
		internalMap[AS],
		internalMap[EU],
		internalMap[NA],
		internalMap[OC],
		internalMap[SA],
	}
}

// Continent specifications EXCLUDING those marked import only.
func ListActive() []Spec {
	return []Spec{
		internalMap[AF],
		internalMap[AN],
		internalMap[AS],
		internalMap[EU],
		internalMap[NA],
		internalMap[OC],
		internalMap[SA],
	}
}

var internalMap = map[Continent]Spec{
	AF: {IsImportOnly: false, Key: "AF", Continent: "Africa"},
	AN: {IsImportOnly: false, Key: "AN", Continent: "Antarctica"},
	AS: {IsImportOnly: false, Key: "AS", Continent: "Asia"},
	EU: {IsImportOnly: false, Key: "EU", Continent: "Europe"},
	NA: {IsImportOnly: false, Key: "NA", Continent: "North America"},
	OC: {IsImportOnly: false, Key: "OC", Continent: "Oceania"},
	SA: {IsImportOnly: false, Key: "SA", Continent: "South America"},
}
