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

// A map of all Continent specifications.
// For convenience, this data is mutable.
// If you require immutable data, please use the specdata package.
var ContinentMap = map[Continent]Spec{
	AF: {IsImportOnly: false, Key: "AF", Continent: "Africa"},
	AN: {IsImportOnly: false, Key: "AN", Continent: "Antarctica"},
	AS: {IsImportOnly: false, Key: "AS", Continent: "Asia"},
	EU: {IsImportOnly: false, Key: "EU", Continent: "Europe"},
	NA: {IsImportOnly: false, Key: "NA", Continent: "North America"},
	OC: {IsImportOnly: false, Key: "OC", Continent: "Oceania"},
	SA: {IsImportOnly: false, Key: "SA", Continent: "South America"},
}

// All Continent specifications including deprecated and import only.
// For convenience, this data is mutable.
// If you require immutable data, please use the specdata package.
var ContinentListAll = []Spec{
	ContinentMap[AF],
	ContinentMap[AN],
	ContinentMap[AS],
	ContinentMap[EU],
	ContinentMap[NA],
	ContinentMap[OC],
	ContinentMap[SA],
}

// All Continent specifications that are NOT marked import-only.
// For convenience, this data is mutable.
// If you require immutable data, please use the specdata package.
var ContinentListCurrent = []Spec{
	ContinentMap[AF],
	ContinentMap[AN],
	ContinentMap[AS],
	ContinentMap[EU],
	ContinentMap[NA],
	ContinentMap[OC],
	ContinentMap[SA],
}
