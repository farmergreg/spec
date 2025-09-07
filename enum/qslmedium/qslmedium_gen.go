// DO NOT EDIT; GENERATED CODE
// run `go generate ./...` from the project root to rebuild this file.

// Package qslmedium provides code and constants as defined in ADIF 3.1.6 (Proposed)
package qslmedium

import "sync"

const (
	CARD QSLMedium = "CARD" // CARD = QSO confirmation via paper QSL card
	EQSL QSLMedium = "EQSL" // EQSL = QSO confirmation via eQSL.cc
	LOTW QSLMedium = "LOTW" // LOTW = QSO confirmation via ARRL Logbook of the World
)

var (
	listActive     []Spec
	listActiveOnce sync.Once
)

// lookupList contains all known QSLMedium specifications
var lookupList = []Spec{
	{IsImportOnly: false, Key: "CARD", Description: "QSO confirmation via paper QSL card"},
	{IsImportOnly: false, Key: "EQSL", Description: "QSO confirmation via eQSL.cc"},
	{IsImportOnly: false, Key: "LOTW", Description: "QSO confirmation via ARRL Logbook of the World"},
}

// lookupMap contains all known QSLMedium specifications
var lookupMap = map[QSLMedium]*Spec{
	CARD: &lookupList[0],
	EQSL: &lookupList[1],
	LOTW: &lookupList[2],
}

// Lookup locates the ADIF 3.1.6 specification for the provided QSLMedium
func Lookup(qslmedium QSLMedium) (Spec, bool) {
	spec, ok := lookupMap[qslmedium]
	if !ok {
		return Spec{}, false
	}
	return *spec, true
}

// LookupByFilter returns all ADIF 3.1.6 QSLMedium specifications that match the provided filter function.
func LookupByFilter(filter func(Spec) bool) []Spec {
	result := make([]Spec, 0, len(lookupList))
	for _, v := range lookupList {
		if filter(v) {
			result = append(result, v)
		}
	}
	return result
}

// ListActive returns a slice of ADIF 3.1.6 QSLMedium specifications, but excludes those marked as import-only.
func ListActive() []Spec {
	listActiveOnce.Do(func() {
		listActive = LookupByFilter(func(spec Spec) bool { return !bool(spec.IsImportOnly) })
	})
	return listActive
}

// List returns a slice of all ADIF 3.1.6 QSLMedium specifications.
func List() []Spec {
	list := make([]Spec, len(lookupList))
	copy(list, lookupList)
	return list
}
