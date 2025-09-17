// DO NOT EDIT; GENERATED CODE
// run `go generate ./...` from the project root to rebuild this file.

// Package qslvia provides code and constants as defined in ADIF 3.1.6 (Released)
package qslvia

import "sync"

const (
	B QSLVia = "b" // b = bureau
	D QSLVia = "d" // d = direct
	E QSLVia = "e" // e = electronic
	M QSLVia = "m" // Deprecated: m = manager
)

var (
	listActive     []Spec    // listActive is a cached copy of the active specifications (those not marked as import-only).
	listActiveOnce sync.Once // listActive is lazy loaded instead of utilizing an init() function. This allows the compiler to remove unused data / variables.
)

// lookupList contains all known QSLVia specifications.
var lookupList = []Spec{
	{IsImportOnly: false, Key: "b", Description: "bureau"},
	{IsImportOnly: false, Key: "d", Description: "direct"},
	{IsImportOnly: false, Key: "e", Description: "electronic"},
	{IsImportOnly: true, Key: "m", Description: "manager"},
}

// lookupMap contains all known QSLVia specifications.
var lookupMap = map[QSLVia]*Spec{
	B: &lookupList[0],
	D: &lookupList[1],
	E: &lookupList[2],
	M: &lookupList[3],
}

// Lookup returns the specification for the provided QSLVia.
// ADIF 3.1.6
func Lookup(q QSLVia) (Spec, bool) {
	spec, ok := lookupMap[q]
	if !ok {
		return Spec{}, false
	}
	return *spec, true
}

// LookupByFilter returns all QSLVia specifications that match the provided filter function.
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

// List returns all QSLVia specifications.
// This list includes those marked import-only.
// ADIF 3.1.6
func List() []Spec {
	list := make([]Spec, len(lookupList))
	copy(list, lookupList)
	return list
}

// ListActive returns QSLVia specifications.
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
