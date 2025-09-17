// DO NOT EDIT; GENERATED CODE
// run `go generate ./...` from the project root to rebuild this file.

// Package morsekeytype provides code and constants as defined in ADIF 3.1.6 (Released)
package morsekeytype

import "sync"

const (
	BUG MorseKeyType = "bug" // bug  = Mechanical semi-automatic keyer or Bug
	CPU MorseKeyType = "cpu" // cpu  = Computer Driven
	DP  MorseKeyType = "dp"  // dp   = Dual Paddle
	FAB MorseKeyType = "fab" // fab  = Mechanical fully-automatic keyer or Bug
	SK  MorseKeyType = "sk"  // sk   = Straight Key
	SP  MorseKeyType = "sp"  // sp   = Single Paddle
	SS  MorseKeyType = "ss"  // ss   = Sideswiper
)

var (
	listActive     []Spec    // listActive is a cached copy of the active specifications (those not marked as import-only).
	listActiveOnce sync.Once // listActive is lazy loaded instead of utilizing an init() function. This allows the compiler to remove unused data / variables.
)

// lookupList contains all known MorseKeyType specifications.
var lookupList = []Spec{
	{IsImportOnly: false, Key: "bug", Description: "Mechanical semi-automatic keyer or Bug", Characteristics: "a control which actuates a switch as well as a control which actuates a spring and pendulum mechanism which actuates a switch. Both switches are wired in parallel.", MorseComposition: "a machine makes the dits and a human makes the dahs and builds characters.", Examples: "Vibroplex Blue Racer Deluxe"},
	{IsImportOnly: false, Key: "cpu", Description: "Computer Driven", Characteristics: "an electronic device performs the actuation of the switch.", MorseComposition: "a machine makes the dits and dahs to build the characters.", Examples: "N1MM+ Logging Software"},
	{IsImportOnly: false, Key: "dp", Description: "Dual Paddle", Characteristics: "two controls which actuate independent switches.", MorseComposition: "a machine makes the dits and the dahs and a human builds the characters.", Examples: "Begali Sculpture, VK3IL pressure paddles, M0UKD capacitive touch paddles"},
	{IsImportOnly: false, Key: "fab", Description: "Mechanical fully-automatic keyer or Bug", Characteristics: "a control which actuates one of two separate spring and pendulum mechanisms at a time, each of which actuates a switch. Both switches are wired in parallel.", MorseComposition: "a machine makes the dits and the dahs and a human builds characters.", Examples: "GHD GN209FA fully-automatic bug"},
	{IsImportOnly: false, Key: "sk", Description: "Straight Key", Characteristics: "a single control which actuates a single switch.", MorseComposition: "a human makes the dits and dahs and builds characters", Examples: "Lionel J-38"},
	{IsImportOnly: false, Key: "sp", Description: "Single Paddle", Characteristics: "a single control which actuates two independent switches.", MorseComposition: "a machine makes the dits and the dahs and a human builds the characters.", Examples: "American Morse Mini-B"},
	{IsImportOnly: false, Key: "ss", Description: "Sideswiper", Characteristics: "a single control which actuates a SPDT (single poll, double throw) switch.", MorseComposition: "a human makes the dits and dahs and builds characters", Examples: "W1SFR Green Machine Torsion Bar Cootie"},
}

// lookupMap contains all known MorseKeyType specifications.
var lookupMap = map[MorseKeyType]*Spec{
	BUG: &lookupList[0],
	CPU: &lookupList[1],
	DP:  &lookupList[2],
	FAB: &lookupList[3],
	SK:  &lookupList[4],
	SP:  &lookupList[5],
	SS:  &lookupList[6],
}

// Lookup returns the specification for the provided MorseKeyType.
// ADIF 3.1.6
func Lookup(m MorseKeyType) (Spec, bool) {
	spec, ok := lookupMap[m]
	if !ok {
		return Spec{}, false
	}
	return *spec, true
}

// LookupByFilter returns all MorseKeyType specifications that match the provided filter function.
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

// List returns all MorseKeyType specifications.
// This list includes those marked import-only.
// ADIF 3.1.6
func List() []Spec {
	list := make([]Spec, len(lookupList))
	copy(list, lookupList)
	return list
}

// ListActive returns MorseKeyType specifications.
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
