// DO NOT EDIT; GENERATED CODE
// run `go generate ./...` from the project root to rebuild this file.

// Package qslsent provides code and constants as defined in ADIF 3.1.6 (Proposed)
package qslsent

import "sync"

const (
	I QSLSent = "I" // I =
	N QSLSent = "N" // N = do not send an outgoing QSL card do not upload the QSO to the online service
	Q QSLSent = "Q" // Q = an outgoing QSL card has been selected to be sent a QSO has been selected to be uploaded to the online service
	R QSLSent = "R" // R = the contacted station has requested a QSL card the contacted station has requested the QSO be uploaded to the online service
	Y QSLSent = "Y" // Y = an outgoing QSL card has been sent the QSO has been uploaded to, and accepted by, the online service
)

var (
	listActive     []Spec    // listActive is a cached copy of the active specifications (those not marked as import-only).
	listActiveOnce sync.Once // listActive is lazy loaded instead of utilizing an init() function. This allows the compiler to remove unused data / variables.
)

// lookupList contains all known QSLSent specifications.
var lookupList = []Spec{
	{IsImportOnly: false, Key: "I", Meaning: "ignore or invalid", Description: ""},
	{IsImportOnly: false, Key: "N", Meaning: "no", Description: "do not send an outgoing QSL card do not upload the QSO to the online service"},
	{IsImportOnly: false, Key: "Q", Meaning: "queued", Description: "an outgoing QSL card has been selected to be sent a QSO has been selected to be uploaded to the online service"},
	{IsImportOnly: false, Key: "R", Meaning: "requested", Description: "the contacted station has requested a QSL card the contacted station has requested the QSO be uploaded to the online service"},
	{IsImportOnly: false, Key: "Y", Meaning: "yes", Description: "an outgoing QSL card has been sent the QSO has been uploaded to, and accepted by, the online service"},
}

// lookupMap contains all known QSLSent specifications.
var lookupMap = map[QSLSent]*Spec{
	I: &lookupList[0],
	N: &lookupList[1],
	Q: &lookupList[2],
	R: &lookupList[3],
	Y: &lookupList[4],
}

// Lookup returns the specification for the provided QSLSent.
// ADIF 3.1.6
func Lookup(qslsent QSLSent) (Spec, bool) {
	spec, ok := lookupMap[qslsent]
	if !ok {
		return Spec{}, false
	}
	return *spec, true
}

// LookupByFilter returns all QSLSent specifications that match the provided filter function.
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

// List returns all QSLSent specifications.
// This list includes those marked import-only.
// ADIF 3.1.6
func List() []Spec {
	list := make([]Spec, len(lookupList))
	copy(list, lookupList)
	return list
}

// ListActive returns QSLSent specifications.
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
