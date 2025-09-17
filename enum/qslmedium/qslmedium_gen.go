// DO NOT EDIT; GENERATED CODE
// run `go generate ./...` from the project root to rebuild this file.

// Package qslmedium provides code and constants as defined in ADIF 3.1.6 (Released)
package qslmedium

import "sync"

const (
	CARD QSLMedium = "card" // card = QSO confirmation via paper QSL card
	EQSL QSLMedium = "eqsl" // eqsl = QSO confirmation via eQSL.cc
	LOTW QSLMedium = "lotw" // lotw = QSO confirmation via ARRL Logbook of the World
)

var (
	listActive     []Spec    // listActive is a cached copy of the active specifications (those not marked as import-only).
	listActiveOnce sync.Once // listActive is lazy loaded instead of utilizing an init() function. This allows the compiler to remove unused data / variables.
)

// lookupList contains all known QSLMedium specifications.
var lookupList = []Spec{
	{IsImportOnly: false, Key: "card", Description: "QSO confirmation via paper QSL card"},
	{IsImportOnly: false, Key: "eqsl", Description: "QSO confirmation via eQSL.cc"},
	{IsImportOnly: false, Key: "lotw", Description: "QSO confirmation via ARRL Logbook of the World"},
}

// lookupMap contains all known QSLMedium specifications.
var lookupMap = map[QSLMedium]*Spec{
	CARD: &lookupList[0],
	EQSL: &lookupList[1],
	LOTW: &lookupList[2],
}

// Lookup returns the specification for the provided QSLMedium.
// ADIF 3.1.6
func Lookup(q QSLMedium) (Spec, bool) {
	spec, ok := lookupMap[q]
	if !ok {
		return Spec{}, false
	}
	return *spec, true
}

// LookupByFilter returns all QSLMedium specifications that match the provided filter function.
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

// List returns all QSLMedium specifications.
// This list includes those marked import-only.
// ADIF 3.1.6
func List() []Spec {
	list := make([]Spec, len(lookupList))
	copy(list, lookupList)
	return list
}

// ListActive returns QSLMedium specifications.
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
