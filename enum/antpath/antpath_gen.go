// DO NOT EDIT; GENERATED CODE
// run `go generate ./...` from the project root to rebuild this file.

// Package antpath provides code and constants as defined in ADIF 3.1.6 (Released)
package antpath

import "sync"

const (
	G AntPath = "G" // G = grayline
	L AntPath = "L" // L = long path
	O AntPath = "O" // O = other
	S AntPath = "S" // S = short path
)

var (
	listActive     []Spec    // listActive is a cached copy of the active specifications (those not marked as import-only).
	listActiveOnce sync.Once // listActive is lazy loaded instead of utilizing an init() function. This allows the compiler to remove unused data / variables.
)

// lookupList contains all known AntPath specifications.
var lookupList = []Spec{
	{IsImportOnly: false, Key: "G", Description: "grayline"},
	{IsImportOnly: false, Key: "L", Description: "long path"},
	{IsImportOnly: false, Key: "O", Description: "other"},
	{IsImportOnly: false, Key: "S", Description: "short path"},
}

// lookupMap contains all known AntPath specifications.
var lookupMap = map[AntPath]*Spec{
	G: &lookupList[0],
	L: &lookupList[1],
	O: &lookupList[2],
	S: &lookupList[3],
}

// Lookup returns the specification for the provided AntPath.
// ADIF 3.1.6
func Lookup(a AntPath) (Spec, bool) {
	spec, ok := lookupMap[a]
	if !ok {
		return Spec{}, false
	}
	return *spec, true
}

// LookupByFilter returns all AntPath specifications that match the provided filter function.
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

// List returns all AntPath specifications.
// This list includes those marked import-only.
// ADIF 3.1.6
func List() []Spec {
	list := make([]Spec, len(lookupList))
	copy(list, lookupList)
	return list
}

// ListActive returns AntPath specifications.
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
