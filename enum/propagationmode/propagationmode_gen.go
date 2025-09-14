// DO NOT EDIT; GENERATED CODE
// run `go generate ./...` from the project root to rebuild this file.

// Package propagationmode provides code and constants as defined in ADIF 3.1.6 (Proposed)
package propagationmode

import "sync"

const (
	AS       PropagationMode = "AS"       // AS         = Aircraft Scatter
	AUE      PropagationMode = "AUE"      // AUE        = Aurora-E
	AUR      PropagationMode = "AUR"      // AUR        = Aurora
	BS       PropagationMode = "BS"       // BS         = Back scatter
	ECH      PropagationMode = "ECH"      // ECH        = EchoLink
	EME      PropagationMode = "EME"      // EME        = Earth-Moon-Earth
	ES       PropagationMode = "ES"       // ES         = Sporadic E
	F2       PropagationMode = "F2"       // F2         = F2 Reflection
	FAI      PropagationMode = "FAI"      // FAI        = Field Aligned Irregularities
	GWAVE    PropagationMode = "GWAVE"    // GWAVE      = Ground Wave
	INTERNET PropagationMode = "INTERNET" // INTERNET   = Internet-assisted
	ION      PropagationMode = "ION"      // ION        = Ionoscatter
	IRL      PropagationMode = "IRL"      // IRL        = IRLP
	LOS      PropagationMode = "LOS"      // LOS        = Line of Sight (includes transmission through obstacles such as walls)
	MS       PropagationMode = "MS"       // MS         = Meteor scatter
	RPT      PropagationMode = "RPT"      // RPT        = Terrestrial or atmospheric repeater or transponder
	RS       PropagationMode = "RS"       // RS         = Rain scatter
	SAT      PropagationMode = "SAT"      // SAT        = Satellite
	TEP      PropagationMode = "TEP"      // TEP        = Trans-equatorial
	TR       PropagationMode = "TR"       // TR         = Tropospheric ducting
)

var (
	listActive     []Spec    // listActive is a cached copy of the active specifications (those not marked as import-only).
	listActiveOnce sync.Once // listActive is lazy loaded instead of utilizing an init() function. This allows the compiler to remove unused data / variables.
)

// lookupList contains all known PropagationMode specifications.
var lookupList = []Spec{
	{IsImportOnly: false, Key: "AS", Description: "Aircraft Scatter"},
	{IsImportOnly: false, Key: "AUE", Description: "Aurora-E"},
	{IsImportOnly: false, Key: "AUR", Description: "Aurora"},
	{IsImportOnly: false, Key: "BS", Description: "Back scatter"},
	{IsImportOnly: false, Key: "ECH", Description: "EchoLink"},
	{IsImportOnly: false, Key: "EME", Description: "Earth-Moon-Earth"},
	{IsImportOnly: false, Key: "ES", Description: "Sporadic E"},
	{IsImportOnly: false, Key: "F2", Description: "F2 Reflection"},
	{IsImportOnly: false, Key: "FAI", Description: "Field Aligned Irregularities"},
	{IsImportOnly: false, Key: "GWAVE", Description: "Ground Wave"},
	{IsImportOnly: false, Key: "INTERNET", Description: "Internet-assisted"},
	{IsImportOnly: false, Key: "ION", Description: "Ionoscatter"},
	{IsImportOnly: false, Key: "IRL", Description: "IRLP"},
	{IsImportOnly: false, Key: "LOS", Description: "Line of Sight (includes transmission through obstacles such as walls)"},
	{IsImportOnly: false, Key: "MS", Description: "Meteor scatter"},
	{IsImportOnly: false, Key: "RPT", Description: "Terrestrial or atmospheric repeater or transponder"},
	{IsImportOnly: false, Key: "RS", Description: "Rain scatter"},
	{IsImportOnly: false, Key: "SAT", Description: "Satellite"},
	{IsImportOnly: false, Key: "TEP", Description: "Trans-equatorial"},
	{IsImportOnly: false, Key: "TR", Description: "Tropospheric ducting"},
}

// lookupMap contains all known PropagationMode specifications.
var lookupMap = map[PropagationMode]*Spec{
	AS:       &lookupList[0],
	AUE:      &lookupList[1],
	AUR:      &lookupList[2],
	BS:       &lookupList[3],
	ECH:      &lookupList[4],
	EME:      &lookupList[5],
	ES:       &lookupList[6],
	F2:       &lookupList[7],
	FAI:      &lookupList[8],
	GWAVE:    &lookupList[9],
	INTERNET: &lookupList[10],
	ION:      &lookupList[11],
	IRL:      &lookupList[12],
	LOS:      &lookupList[13],
	MS:       &lookupList[14],
	RPT:      &lookupList[15],
	RS:       &lookupList[16],
	SAT:      &lookupList[17],
	TEP:      &lookupList[18],
	TR:       &lookupList[19],
}

// Lookup returns the specification for the provided PropagationMode.
// ADIF 3.1.6
func Lookup(p PropagationMode) (Spec, bool) {
	spec, ok := lookupMap[p]
	if !ok {
		return Spec{}, false
	}
	return *spec, true
}

// LookupByFilter returns all PropagationMode specifications that match the provided filter function.
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

// List returns all PropagationMode specifications.
// This list includes those marked import-only.
// ADIF 3.1.6
func List() []Spec {
	list := make([]Spec, len(lookupList))
	copy(list, lookupList)
	return list
}

// ListActive returns PropagationMode specifications.
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
