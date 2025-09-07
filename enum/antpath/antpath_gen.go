// DO NOT EDIT; GENERATED CODE
// run `go generate ./...` from the project root to rebuild this file.

// Package antpath provides code and constants as defined in ADIF 3.1.6 (Proposed)
package antpath

import "sync"

const (
	G AntPath = "G" // G = grayline
	L AntPath = "L" // L = long path
	O AntPath = "O" // O = other
	S AntPath = "S" // S = short path
)

var (
	listActive     []Spec
	listActiveOnce sync.Once
)

// lookupList contains all known AntPath specifications
var lookupList = []Spec{
	{IsImportOnly: false, Key: "G", Description: "grayline"},
	{IsImportOnly: false, Key: "L", Description: "long path"},
	{IsImportOnly: false, Key: "O", Description: "other"},
	{IsImportOnly: false, Key: "S", Description: "short path"},
}

// lookupMap contains all known AntPath specifications
var lookupMap = map[AntPath]*Spec{
	G: &lookupList[0],
	L: &lookupList[1],
	O: &lookupList[2],
	S: &lookupList[3],
}

// Lookup locates the ADIF 3.1.6 specification for the provided AntPath
func Lookup(antpath AntPath) (Spec, bool) {
	spec, ok := lookupMap[antpath]
	if !ok {
		return Spec{}, false
	}
	return *spec, true
}

// LookupByFilter returns all ADIF 3.1.6 AntPath specifications that match the provided filter function.
func LookupByFilter(filter func(Spec) bool) []Spec {
	result := make([]Spec, 0, len(lookupList))
	for _, v := range lookupList {
		if filter(v) {
			result = append(result, v)
		}
	}
	return result
}

// ListActive returns a slice of ADIF 3.1.6 AntPath specifications, but excludes those marked as import-only.
func ListActive() []Spec {
	listActiveOnce.Do(func() {
		listActive = LookupByFilter(func(spec Spec) bool { return !bool(spec.IsImportOnly) })
	})
	return listActive
}

// List returns a slice of all ADIF 3.1.6 AntPath specifications. This includes those marked import-only.
func List() []Spec {
	list := make([]Spec, len(lookupList))
	copy(list, lookupList)
	return list
}
