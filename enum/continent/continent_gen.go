// DO NOT EDIT; GENERATED CODE
// run `go generate ./...` from the project root to rebuild this file.

// Package continent provides code and constants as defined in ADIF 3.1.6 (Proposed)
package continent

import "sync"

const (
	AF Continent = "AF" // AF = Africa
	AN Continent = "AN" // AN = Antarctica
	AS Continent = "AS" // AS = Asia
	EU Continent = "EU" // EU = Europe
	NA Continent = "NA" // NA = North America
	OC Continent = "OC" // OC = Oceania
	SA Continent = "SA" // SA = South America
)

var (
	listActive     []Spec
	listActiveOnce sync.Once
)

// lookupList contains all known Continent specifications
var lookupList = []Spec{
	{IsImportOnly: false, Key: "AF", Continent: "Africa"},
	{IsImportOnly: false, Key: "AN", Continent: "Antarctica"},
	{IsImportOnly: false, Key: "AS", Continent: "Asia"},
	{IsImportOnly: false, Key: "EU", Continent: "Europe"},
	{IsImportOnly: false, Key: "NA", Continent: "North America"},
	{IsImportOnly: false, Key: "OC", Continent: "Oceania"},
	{IsImportOnly: false, Key: "SA", Continent: "South America"},
}

// lookupMap contains all known Continent specifications
var lookupMap = map[Continent]*Spec{
	AF: &lookupList[0],
	AN: &lookupList[1],
	AS: &lookupList[2],
	EU: &lookupList[3],
	NA: &lookupList[4],
	OC: &lookupList[5],
	SA: &lookupList[6],
}

// Lookup locates the ADIF 3.1.6 specification for the provided Continent
func Lookup(continent Continent) (Spec, bool) {
	spec, ok := lookupMap[continent]
	if !ok {
		return Spec{}, false
	}
	return *spec, true
}

// LookupByFilter returns all ADIF 3.1.6 Continent specifications that match the provided filter function.
func LookupByFilter(filter func(Spec) bool) []Spec {
	result := make([]Spec, 0, len(lookupList))
	for _, v := range lookupList {
		if filter(v) {
			result = append(result, v)
		}
	}
	return result
}

// ListActive returns a slice of ADIF 3.1.6 Continent specifications, but excludes those marked as import-only.
func ListActive() []Spec {
	listActiveOnce.Do(func() {
		listActive = LookupByFilter(func(spec Spec) bool { return !bool(spec.IsImportOnly) })
	})
	return listActive
}

// List returns a slice of all ADIF 3.1.6 Continent specifications. This includes those marked import-only.
func List() []Spec {
	list := make([]Spec, len(lookupList))
	copy(list, lookupList)
	return list
}
