// DO NOT EDIT; GENERATED CODE
// run `go generate ./...` from the project root to rebuild this file.

// Package morsekeytype provides code and constants as defined in ADIF 3.1.6 (Proposed)
package morsekeytype

import "maps"

const (
	BUG MorseKeyType = "BUG" // BUG  = Mechanical semi-automatic keyer or Bug
	CPU MorseKeyType = "CPU" // CPU  = Computer Driven
	DP  MorseKeyType = "DP"  // DP   = Dual Paddle
	FAB MorseKeyType = "FAB" // FAB  = Mechanical fully-automatic keyer or Bug
	SK  MorseKeyType = "SK"  // SK   = Straight Key
	SP  MorseKeyType = "SP"  // SP   = Single Paddle
	SS  MorseKeyType = "SS"  // SS   = Sideswiper
)

// All MorseKeyType specifications including depreciated and import only.
func MorseKeyTypeListAll() []Spec {
	return append([]Spec(nil), internalMorseKeyTypeListAll...)
}

// All MorseKeyType specifications values that are NOT marked import-only.
func MorseKeyTypeListCurrent() []Spec {
	return append([]Spec(nil), internalMorseKeyTypeListCurrent...)
}

// A map of all MorseKeyType specifications.
func MorseKeyTypeMap() map[MorseKeyType]Spec {
	cp := make(map[MorseKeyType]Spec, len(internalMorseKeyTypeMap))
	maps.Copy(cp, internalMorseKeyTypeMap)
	return cp
}

// A map of all MorseKeyType specifications.
var internalMorseKeyTypeMap = map[MorseKeyType]Spec{
	BUG: {IsImportOnly: false, Key: "BUG", Description: "Mechanical semi-automatic keyer or Bug", Characteristics: "a control which actuates a switch as well as a control which actuates a spring and pendulum mechanism which actuates a switch. Both switches are wired in parallel.", MorseComposition: "a machine makes the dits and a human makes the dahs and builds characters.", Examples: "Vibroplex Blue Racer Deluxe"},
	CPU: {IsImportOnly: false, Key: "CPU", Description: "Computer Driven", Characteristics: "an electronic device performs the actuation of the switch.", MorseComposition: "a machine makes the dits and dahs to build the characters.", Examples: "N1MM+ Logging Software"},
	DP:  {IsImportOnly: false, Key: "DP", Description: "Dual Paddle", Characteristics: "two controls which actuate independent switches.", MorseComposition: "a machine makes the dits and the dahs and a human builds the characters.", Examples: "Begali Sculpture, VK3IL pressure paddles, M0UKD capacitive touch paddles"},
	FAB: {IsImportOnly: false, Key: "FAB", Description: "Mechanical fully-automatic keyer or Bug", Characteristics: "a control which actuates one of two separate spring and pendulum mechanisms at a time, each of which actuates a switch. Both switches are wired in parallel.", MorseComposition: "a machine makes the dits and the dahs and a human builds characters.", Examples: "GHD GN209FA fully-automatic bug"},
	SK:  {IsImportOnly: false, Key: "SK", Description: "Straight Key", Characteristics: "a single control which actuates a single switch.", MorseComposition: "a human makes the dits and dahs and builds characters", Examples: "Lionel J-38"},
	SP:  {IsImportOnly: false, Key: "SP", Description: "Single Paddle", Characteristics: "a single control which actuates two independent switches.", MorseComposition: "a machine makes the dits and the dahs and a human builds the characters.", Examples: "American Morse Mini-B"},
	SS:  {IsImportOnly: false, Key: "SS", Description: "Sideswiper", Characteristics: "a single control which actuates a SPDT (single poll, double throw) switch.", MorseComposition: "a human makes the dits and dahs and builds characters", Examples: "W1SFR Green Machine Torsion Bar Cootie"},
}

var internalMorseKeyTypeListAll = []Spec{
	internalMorseKeyTypeMap[BUG],
	internalMorseKeyTypeMap[CPU],
	internalMorseKeyTypeMap[DP],
	internalMorseKeyTypeMap[FAB],
	internalMorseKeyTypeMap[SK],
	internalMorseKeyTypeMap[SP],
	internalMorseKeyTypeMap[SS],
}

var internalMorseKeyTypeListCurrent = []Spec{
	internalMorseKeyTypeMap[BUG],
	internalMorseKeyTypeMap[CPU],
	internalMorseKeyTypeMap[DP],
	internalMorseKeyTypeMap[FAB],
	internalMorseKeyTypeMap[SK],
	internalMorseKeyTypeMap[SP],
	internalMorseKeyTypeMap[SS],
}
