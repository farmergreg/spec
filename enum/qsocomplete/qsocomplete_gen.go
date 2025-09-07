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
	listActive     []Spec
	listActiveOnce sync.Once
)

// lookupList contains all known QSOComplete specifications
var lookupList = []Spec{
	{IsImportOnly: false, Key: "?", Description: "uncertain"},
	{IsImportOnly: false, Key: "N", Description: "no"},
	{IsImportOnly: false, Key: "NIL", Description: "not heard"},
	{IsImportOnly: false, Key: "Y", Description: "yes"},
}

// lookupMap contains all known QSOComplete specifications
var lookupMap = map[QSOComplete]*Spec{
	Uncertain: &lookupList[0],
	N:         &lookupList[1],
	NIL:       &lookupList[2],
	Y:         &lookupList[3],
}

// Lookup locates the specification for the given QSOComplete
func Lookup(qsocomplete QSOComplete) (Spec, bool) {
	spec, ok := lookupMap[qsocomplete]
	if !ok {
		return Spec{}, false
	}
	return *spec, true
}

// LookupByFilter returns all QSOComplete specifications that match the provided filter function.
func LookupByFilter(filter func(Spec) bool) []Spec {
	result := make([]Spec, 0, len(lookupList))
	for _, v := range lookupList {
		if filter(v) {
			result = append(result, v)
		}
	}
	return result
}

// ListActive returns a slice of QSOComplete specifications excluding those marked as import-only.
func ListActive() []Spec {
	listActiveOnce.Do(func() {
		listActive = LookupByFilter(func(spec Spec) bool { return !bool(spec.IsImportOnly) })
	})
	return listActive
}

// List returns a slice of all QSOComplete specifications including those marked as import-only.
func List() []Spec {
	list := make([]Spec, len(lookupList))
	copy(list, lookupList)
	return list
}
