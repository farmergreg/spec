// DO NOT EDIT; GENERATED CODE
// run `go generate ./...` from the project root to rebuild this file.

// Package qsodownloadstatus provides code and constants as defined in ADIF 3.1.6 (Released)
package qsodownloadstatus

import "sync"

const (
	I QSODownloadStatus = "I" // I = ignore or invalid
	N QSODownloadStatus = "N" // N = the QSO has not been downloaded from the online service
	Y QSODownloadStatus = "Y" // Y = the QSO has been downloaded from the online service
)

var (
	listActive     []Spec    // listActive is a cached copy of the active specifications (those not marked as import-only).
	listActiveOnce sync.Once // listActive is lazy loaded instead of utilizing an init() function. This allows the compiler to remove unused data / variables.
)

// lookupList contains all known QSODownloadStatus specifications.
var lookupList = []Spec{
	{IsImportOnly: false, Key: "I", Description: "ignore or invalid"},
	{IsImportOnly: false, Key: "N", Description: "the QSO has not been downloaded from the online service"},
	{IsImportOnly: false, Key: "Y", Description: "the QSO has been downloaded from the online service"},
}

// lookupMap contains all known QSODownloadStatus specifications.
var lookupMap = map[QSODownloadStatus]*Spec{
	I: &lookupList[0],
	N: &lookupList[1],
	Y: &lookupList[2],
}

// Lookup returns the specification for the provided QSODownloadStatus.
// ADIF 3.1.6
func Lookup(q QSODownloadStatus) (Spec, bool) {
	spec, ok := lookupMap[q]
	if !ok {
		return Spec{}, false
	}
	return *spec, true
}

// LookupByFilter returns all QSODownloadStatus specifications that match the provided filter function.
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

// List returns all QSODownloadStatus specifications.
// This list includes those marked import-only.
// ADIF 3.1.6
func List() []Spec {
	list := make([]Spec, len(lookupList))
	copy(list, lookupList)
	return list
}

// ListActive returns QSODownloadStatus specifications.
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
