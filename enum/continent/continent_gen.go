// DO NOT EDIT; GENERATED CODE
// ADIF: 3.1.6 Proposed

package continent

var (
	AF Continent = "AF" // AF = Africa
	AN Continent = "AN" // AN = Antarctica
	AS Continent = "AS" // AS = Asia
	EU Continent = "EU" // EU = Europe
	NA Continent = "NA" // NA = North America
	OC Continent = "OC" // OC = Oceania
	SA Continent = "SA" // SA = South America
)

// A map of all Continent specifications.
var ContinentMap = map[Continent]Spec{
	AF: {IsImportOnly: false, Key: "AF", Continent: "Africa"},
	AN: {IsImportOnly: false, Key: "AN", Continent: "Antarctica"},
	AS: {IsImportOnly: false, Key: "AS", Continent: "Asia"},
	EU: {IsImportOnly: false, Key: "EU", Continent: "Europe"},
	NA: {IsImportOnly: false, Key: "NA", Continent: "North America"},
	OC: {IsImportOnly: false, Key: "OC", Continent: "Oceania"},
	SA: {IsImportOnly: false, Key: "SA", Continent: "South America"},
}

// All Continent specifications including depreciated and import only.
var ContinentListAll = []Spec{
	ContinentMap[AF],
	ContinentMap[AN],
	ContinentMap[AS],
	ContinentMap[EU],
	ContinentMap[NA],
	ContinentMap[OC],
	ContinentMap[SA],
}

// All Continent specifications values that are NOT marked import-only.
var ContinentListCurrent = []Spec{
	ContinentMap[AF],
	ContinentMap[AN],
	ContinentMap[AS],
	ContinentMap[EU],
	ContinentMap[NA],
	ContinentMap[OC],
	ContinentMap[SA],
}
