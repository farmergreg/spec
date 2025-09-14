// DO NOT EDIT; GENERATED CODE
// run `go generate ./...` from the project root to rebuild this file.

// Package qsocomplete provides code and constants as defined in ADIF 3.1.6 (Proposed)
package qsocomplete

import "sync"

const (
	Uncertain QSOComplete = "?"   // ?    = uncertain
	N         QSOComplete = "N"   // N    = no
	NIL       QSOComplete = "NIL" // NIL  = not heard
	Y         QSOComplete = "Y"   // Y    = yes
)

var (
	listActive     []Spec    // listActive is a cached copy of the active specifications (those not marked as import-only).
	listActiveOnce sync.Once // listActive is lazy loaded instead of utilizing an init() function. This allows the compiler to remove unused data / variables.
)

// lookupList contains all known QSOComplete specifications.
var lookupList = []Spec{
	{IsImportOnly: false, Key: "?", Description: "uncertain"},
	{IsImportOnly: false, Key: "N", Description: "no"},
	{IsImportOnly: false, Key: "NIL", Description: "not heard"},
	{IsImportOnly: false, Key: "Y", Description: "yes"},
}

// lookupMap contains all known QSOComplete specifications.
var lookupMap = map[QSOComplete]*Spec{
	Uncertain: &lookupList[0],
	N:         &lookupList[1],
	NIL:       &lookupList[2],
	Y:         &lookupList[3],
}

// Lookup returns the specification for the provided QSOComplete.
// ADIF 3.1.6
func Lookup(q QSOComplete) (Spec, bool) {
	spec, ok := lookupMap[q]
	if !ok {
		return Spec{}, false
	}
	return *spec, true
}

// LookupByFilter returns all QSOComplete specifications that match the provided filter function.
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

// List returns all QSOComplete specifications.
// This list includes those marked import-only.
// ADIF 3.1.6
func List() []Spec {
	list := make([]Spec, len(lookupList))
	copy(list, lookupList)
	return list
}

// ListActive returns QSOComplete specifications.
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
