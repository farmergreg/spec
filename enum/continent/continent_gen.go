// DO NOT EDIT; GENERATED CODE
// run `go generate ./...` from the project root to rebuild this file.

// Package continent provides code and constants as defined in ADIF 3.1.6 (Proposed)
package continent

import "maps"

const (
	AF Continent = "AF" // AF = Africa
	AN Continent = "AN" // AN = Antarctica
	AS Continent = "AS" // AS = Asia
	EU Continent = "EU" // EU = Europe
	NA Continent = "NA" // NA = North America
	OC Continent = "OC" // OC = Oceania
	SA Continent = "SA" // SA = South America
)

// All Continent specifications including depreciated and import only.
func ContinentListAll() []Spec {
	return append([]Spec(nil), internalContinentListAll...)
}

// All Continent specifications values that are NOT marked import-only.
func ContinentListCurrent() []Spec {
	return append([]Spec(nil), internalContinentListCurrent...)
}

// A map of all Continent specifications.
func ContinentMap() map[Continent]Spec {
	cp := make(map[Continent]Spec, len(internalContinentMap))
	maps.Copy(cp, internalContinentMap)
	return cp
}

// A map of all Continent specifications.
var internalContinentMap = map[Continent]Spec{
	AF: {IsImportOnly: false, Key: "AF", Continent: "Africa"},
	AN: {IsImportOnly: false, Key: "AN", Continent: "Antarctica"},
	AS: {IsImportOnly: false, Key: "AS", Continent: "Asia"},
	EU: {IsImportOnly: false, Key: "EU", Continent: "Europe"},
	NA: {IsImportOnly: false, Key: "NA", Continent: "North America"},
	OC: {IsImportOnly: false, Key: "OC", Continent: "Oceania"},
	SA: {IsImportOnly: false, Key: "SA", Continent: "South America"},
}

var internalContinentListAll = []Spec{
	internalContinentMap[AF],
	internalContinentMap[AN],
	internalContinentMap[AS],
	internalContinentMap[EU],
	internalContinentMap[NA],
	internalContinentMap[OC],
	internalContinentMap[SA],
}

var internalContinentListCurrent = []Spec{
	internalContinentMap[AF],
	internalContinentMap[AN],
	internalContinentMap[AS],
	internalContinentMap[EU],
	internalContinentMap[NA],
	internalContinentMap[OC],
	internalContinentMap[SA],
}
