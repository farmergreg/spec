// DO NOT EDIT; GENERATED CODE
// run `go generate ./...` from the project root to rebuild this file.

// Package propagationmode provides code and constants as defined in ADIF 3.1.6 (Released)
package propagationmode

import "sync"

const (
	AS       PropagationMode = "as"       // as         = Aircraft Scatter
	AUE      PropagationMode = "aue"      // aue        = Aurora-E
	AUR      PropagationMode = "aur"      // aur        = Aurora
	BS       PropagationMode = "bs"       // bs         = Back scatter
	ECH      PropagationMode = "ech"      // ech        = EchoLink
	EME      PropagationMode = "eme"      // eme        = Earth-Moon-Earth
	ES       PropagationMode = "es"       // es         = Sporadic E
	F2       PropagationMode = "f2"       // f2         = F2 Reflection
	FAI      PropagationMode = "fai"      // fai        = Field Aligned Irregularities
	GWAVE    PropagationMode = "gwave"    // gwave      = Ground Wave
	INTERNET PropagationMode = "internet" // internet   = Internet-assisted
	ION      PropagationMode = "ion"      // ion        = Ionoscatter
	IRL      PropagationMode = "irl"      // irl        = IRLP
	LOS      PropagationMode = "los"      // los        = Line of Sight (includes transmission through obstacles such as walls)
	MS       PropagationMode = "ms"       // ms         = Meteor scatter
	RPT      PropagationMode = "rpt"      // rpt        = Terrestrial or atmospheric repeater or transponder
	RS       PropagationMode = "rs"       // rs         = Rain scatter
	SAT      PropagationMode = "sat"      // sat        = Satellite
	TEP      PropagationMode = "tep"      // tep        = Trans-equatorial
	TR       PropagationMode = "tr"       // tr         = Tropospheric ducting
)

var (
	listActive     []Spec    // listActive is a cached copy of the active specifications (those not marked as import-only).
	listActiveOnce sync.Once // listActive is lazy loaded instead of utilizing an init() function. This allows the compiler to remove unused data / variables.
)

// lookupList contains all known PropagationMode specifications.
var lookupList = []Spec{
	{IsImportOnly: false, Key: "as", Description: "Aircraft Scatter"},
	{IsImportOnly: false, Key: "aue", Description: "Aurora-E"},
	{IsImportOnly: false, Key: "aur", Description: "Aurora"},
	{IsImportOnly: false, Key: "bs", Description: "Back scatter"},
	{IsImportOnly: false, Key: "ech", Description: "EchoLink"},
	{IsImportOnly: false, Key: "eme", Description: "Earth-Moon-Earth"},
	{IsImportOnly: false, Key: "es", Description: "Sporadic E"},
	{IsImportOnly: false, Key: "f2", Description: "F2 Reflection"},
	{IsImportOnly: false, Key: "fai", Description: "Field Aligned Irregularities"},
	{IsImportOnly: false, Key: "gwave", Description: "Ground Wave"},
	{IsImportOnly: false, Key: "internet", Description: "Internet-assisted"},
	{IsImportOnly: false, Key: "ion", Description: "Ionoscatter"},
	{IsImportOnly: false, Key: "irl", Description: "IRLP"},
	{IsImportOnly: false, Key: "los", Description: "Line of Sight (includes transmission through obstacles such as walls)"},
	{IsImportOnly: false, Key: "ms", Description: "Meteor scatter"},
	{IsImportOnly: false, Key: "rpt", Description: "Terrestrial or atmospheric repeater or transponder"},
	{IsImportOnly: false, Key: "rs", Description: "Rain scatter"},
	{IsImportOnly: false, Key: "sat", Description: "Satellite"},
	{IsImportOnly: false, Key: "tep", Description: "Trans-equatorial"},
	{IsImportOnly: false, Key: "tr", Description: "Tropospheric ducting"},
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
