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
	listActive     []Spec    // listActive is a cached copy of the active specifications (those not marked as import-only).
	listActiveOnce sync.Once // listActive is lazy loaded instead of utilizing an init() function. This allows the compiler to remove unused data / variables.
)

// lookupList contains all known Continent specifications.
var lookupList = []Spec{
	{IsImportOnly: false, Key: "AF", Continent: "Africa"},
	{IsImportOnly: false, Key: "AN", Continent: "Antarctica"},
	{IsImportOnly: false, Key: "AS", Continent: "Asia"},
	{IsImportOnly: false, Key: "EU", Continent: "Europe"},
	{IsImportOnly: false, Key: "NA", Continent: "North America"},
	{IsImportOnly: false, Key: "OC", Continent: "Oceania"},
	{IsImportOnly: false, Key: "SA", Continent: "South America"},
}

// lookupMap contains all known Continent specifications.
var lookupMap = map[Continent]*Spec{
	AF: &lookupList[0],
	AN: &lookupList[1],
	AS: &lookupList[2],
	EU: &lookupList[3],
	NA: &lookupList[4],
	OC: &lookupList[5],
	SA: &lookupList[6],
}

// Lookup returns the specification for the provided Continent.
// ADIF 3.1.6
func Lookup(continent Continent) (Spec, bool) {
	spec, ok := lookupMap[continent]
	if !ok {
		return Spec{}, false
	}
	return *spec, true
}

// LookupByFilter returns all Continent specifications that match the provided filter function.
// ADIF 3.1.6
func LookupByFilter(filter func(Spec) bool) []Spec {
	result := make([]Spec, 0, len(lookupList))
	for _, v := range lookupList {
		if filter(v) {
			result = append(result, v)
		}
	}
	return result
}

// List returns all Continent specifications.
// This list includes those marked import-only.
// ADIF 3.1.6
func List() []Spec {
	list := make([]Spec, len(lookupList))
	copy(list, lookupList)
	return list
}

// ListActive returns Continent specifications.
// This list excludes those marked as import-only.
// ADIF 3.1.6
func ListActive() []Spec {
	listActiveOnce.Do(func() {
		listActive = LookupByFilter(func(spec Spec) bool { return !bool(spec.IsImportOnly) })
	})

	result := make([]Spec, len(listActive))
	copy(result, listActive)
	return result
}
