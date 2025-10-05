// DO NOT EDIT; GENERATED CODE
// run `go generate ./...` from the project root to rebuild this file.

// Package eqslag provides code and constants as defined in ADIF 3.1.6 (Released)
package eqslag

import "sync"

const (
	N EQSLAG = "N" // N = the QSO is confirmed but not "Authenticity Guaranteed" by eQSL
	U EQSLAG = "U" // U = unknown
	Y EQSLAG = "Y" // Y = the QSO is confirmed and "Authenticity Guaranteed" by eQSL
)

var (
	listActive     []Spec    // listActive is a cached copy of the active specifications (those not marked as import-only).
	listActiveOnce sync.Once // listActive is lazy loaded instead of utilizing an init() function. This allows the compiler to remove unused data / variables.
)

// lookupList contains all known EQSLAG specifications.
var lookupList = []Spec{
	{IsImportOnly: false, Key: "N", Description: "the QSO is confirmed but not \"Authenticity Guaranteed\" by eQSL"},
	{IsImportOnly: false, Key: "U", Description: "unknown"},
	{IsImportOnly: false, Key: "Y", Description: "the QSO is confirmed and \"Authenticity Guaranteed\" by eQSL"},
}

// lookupMap contains all known EQSLAG specifications.
var lookupMap = map[EQSLAG]*Spec{
	N: &lookupList[0],
	U: &lookupList[1],
	Y: &lookupList[2],
}

// Lookup returns the specification for the provided EQSLAG.
// ADIF 3.1.6
func Lookup(e EQSLAG) (Spec, bool) {
	spec, ok := lookupMap[e]
	if !ok {
		return Spec{}, false
	}
	return *spec, true
}

// LookupByFilter returns all EQSLAG specifications that match the provided filter function.
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

// List returns all EQSLAG specifications.
// This list includes those marked import-only.
// ADIF 3.1.6
func List() []Spec {
	list := make([]Spec, len(lookupList))
	copy(list, lookupList)
	return list
}

// ListActive returns EQSLAG specifications.
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
