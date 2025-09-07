// DO NOT EDIT; GENERATED CODE
// run `go generate ./...` from the project root to rebuild this file.

// Package qslvia provides code and constants as defined in ADIF 3.1.6 (Proposed)
package qslvia

import "sync"

const (
	B QSLVia = "B" // B = bureau
	D QSLVia = "D" // D = direct
	E QSLVia = "E" // E = electronic
	M QSLVia = "M" // Deprecated: M = manager
)

var (
	listActive     []Spec
	listActiveOnce sync.Once
)

// lookupList contains all known QSLVia specifications
var lookupList = []Spec{
	{IsImportOnly: false, Key: "B", Description: "bureau"},
	{IsImportOnly: false, Key: "D", Description: "direct"},
	{IsImportOnly: false, Key: "E", Description: "electronic"},
	{IsImportOnly: true, Key: "M", Description: "manager"},
}

// lookupMap contains all known QSLVia specifications
var lookupMap = map[QSLVia]*Spec{
	B: &lookupList[0],
	D: &lookupList[1],
	E: &lookupList[2],
	M: &lookupList[3],
}

// Lookup locates the specification for the given QSLVia
func Lookup(qslvia QSLVia) (Spec, bool) {
	spec, ok := lookupMap[qslvia]
	if !ok {
		return Spec{}, false
	}
	return *spec, true
}

// LookupByFilter returns all QSLVia specifications that match the provided filter function.
func LookupByFilter(filter func(Spec) bool) []Spec {
	result := make([]Spec, 0, len(lookupList))
	for _, v := range lookupList {
		if filter(v) {
			result = append(result, v)
		}
	}
	return result
}

// ListActive returns a slice of QSLVia specifications excluding those marked as import-only.
func ListActive() []Spec {
	listActiveOnce.Do(func() {
		listActive = LookupByFilter(func(spec Spec) bool { return !bool(spec.IsImportOnly) })
	})
	return listActive
}

// List returns a slice of all QSLVia specifications including those marked as import-only.
func List() []Spec {
	list := make([]Spec, len(lookupList))
	copy(list, lookupList)
	return list
}
