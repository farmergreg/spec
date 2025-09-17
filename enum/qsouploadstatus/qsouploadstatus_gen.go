// DO NOT EDIT; GENERATED CODE
// run `go generate ./...` from the project root to rebuild this file.

// Package qsouploadstatus provides code and constants as defined in ADIF 3.1.6 (Released)
package qsouploadstatus

import "sync"

const (
	M QSOUploadStatus = "m" // m = the QSO has been modified since being uploaded to the online service
	N QSOUploadStatus = "n" // n = do not upload the QSO to the online service
	Y QSOUploadStatus = "y" // y = the QSO has been uploaded to, and accepted by, the online service
)

var (
	listActive     []Spec    // listActive is a cached copy of the active specifications (those not marked as import-only).
	listActiveOnce sync.Once // listActive is lazy loaded instead of utilizing an init() function. This allows the compiler to remove unused data / variables.
)

// lookupList contains all known QSOUploadStatus specifications.
var lookupList = []Spec{
	{IsImportOnly: false, Key: "m", Description: "the QSO has been modified since being uploaded to the online service"},
	{IsImportOnly: false, Key: "n", Description: "do not upload the QSO to the online service"},
	{IsImportOnly: false, Key: "y", Description: "the QSO has been uploaded to, and accepted by, the online service"},
}

// lookupMap contains all known QSOUploadStatus specifications.
var lookupMap = map[QSOUploadStatus]*Spec{
	M: &lookupList[0],
	N: &lookupList[1],
	Y: &lookupList[2],
}

// Lookup returns the specification for the provided QSOUploadStatus.
// ADIF 3.1.6
func Lookup(q QSOUploadStatus) (Spec, bool) {
	spec, ok := lookupMap[q]
	if !ok {
		return Spec{}, false
	}
	return *spec, true
}

// LookupByFilter returns all QSOUploadStatus specifications that match the provided filter function.
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

// List returns all QSOUploadStatus specifications.
// This list includes those marked import-only.
// ADIF 3.1.6
func List() []Spec {
	list := make([]Spec, len(lookupList))
	copy(list, lookupList)
	return list
}

// ListActive returns QSOUploadStatus specifications.
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
