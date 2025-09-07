// DO NOT EDIT; GENERATED CODE
// run `go generate ./...` from the project root to rebuild this file.

// Package eqslag provides code and constants as defined in ADIF 3.1.6 (Proposed)
package eqslag

import "sync"

const (
	N EQSLAG = "N" // N = the QSO is confirmed but not "Authenticity Guaranteed" by eQSL
	U EQSLAG = "U" // U = unknown
	Y EQSLAG = "Y" // Y = the QSO is confirmed and "Authenticity Guaranteed" by eQSL
)

var (
	listActive     []Spec
	listActiveOnce sync.Once
)

// lookupList contains all known EQSLAG specifications
var lookupList = []Spec{
	{IsImportOnly: false, Key: "N", Description: "the QSO is confirmed but not \"Authenticity Guaranteed\" by eQSL"},
	{IsImportOnly: false, Key: "U", Description: "unknown"},
	{IsImportOnly: false, Key: "Y", Description: "the QSO is confirmed and \"Authenticity Guaranteed\" by eQSL"},
}

// lookupMap contains all known EQSLAG specifications
var lookupMap = map[EQSLAG]*Spec{
	N: &lookupList[0],
	U: &lookupList[1],
	Y: &lookupList[2],
}

// Lookup locates the ADIF 3.1.6 specification for the provided EQSLAG
func Lookup(eqslag EQSLAG) (Spec, bool) {
	spec, ok := lookupMap[eqslag]
	if !ok {
		return Spec{}, false
	}
	return *spec, true
}

// LookupByFilter returns all ADIF 3.1.6 EQSLAG specifications that match the provided filter function.
func LookupByFilter(filter func(Spec) bool) []Spec {
	result := make([]Spec, 0, len(lookupList))
	for _, v := range lookupList {
		if filter(v) {
			result = append(result, v)
		}
	}
	return result
}

// ListActive returns a slice of ADIF 3.1.6 EQSLAG specifications, but excludes those marked as import-only.
func ListActive() []Spec {
	listActiveOnce.Do(func() {
		listActive = LookupByFilter(func(spec Spec) bool { return !bool(spec.IsImportOnly) })
	})
	return listActive
}

// List returns a slice of all ADIF 3.1.6 EQSLAG specifications. This includes those marked import-only.
func List() []Spec {
	list := make([]Spec, len(lookupList))
	copy(list, lookupList)
	return list
}
