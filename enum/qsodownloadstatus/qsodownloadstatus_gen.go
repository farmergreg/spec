// DO NOT EDIT; GENERATED CODE
// run `go generate ./...` from the project root to rebuild this file.

// Package qsodownloadstatus provides code and constants as defined in ADIF 3.1.6 (Proposed)
package qsodownloadstatus

import "sync"

const (
	I QSODownloadStatus = "I" // I = ignore or invalid
	N QSODownloadStatus = "N" // N = the QSO has not been downloaded from the online service
	Y QSODownloadStatus = "Y" // Y = the QSO has been downloaded from the online service
)

var (
	listActive     []Spec
	listActiveOnce sync.Once
)

// lookupList contains all known QSODownloadStatus specifications
var lookupList = []Spec{
	{IsImportOnly: false, Key: "I", Description: "ignore or invalid"},
	{IsImportOnly: false, Key: "N", Description: "the QSO has not been downloaded from the online service"},
	{IsImportOnly: false, Key: "Y", Description: "the QSO has been downloaded from the online service"},
}

// lookupMap contains all known QSODownloadStatus specifications
var lookupMap = map[QSODownloadStatus]*Spec{
	I: &lookupList[0],
	N: &lookupList[1],
	Y: &lookupList[2],
}

// Lookup locates the ADIF 3.1.6 specification for the provided QSODownloadStatus
func Lookup(qsodownloadstatus QSODownloadStatus) (Spec, bool) {
	spec, ok := lookupMap[qsodownloadstatus]
	if !ok {
		return Spec{}, false
	}
	return *spec, true
}

// LookupByFilter returns all ADIF 3.1.6 QSODownloadStatus specifications that match the provided filter function.
func LookupByFilter(filter func(Spec) bool) []Spec {
	result := make([]Spec, 0, len(lookupList))
	for _, v := range lookupList {
		if filter(v) {
			result = append(result, v)
		}
	}
	return result
}

// ListActive returns a slice of ADIF 3.1.6 QSODownloadStatus specifications, but excludes those marked as import-only.
func ListActive() []Spec {
	listActiveOnce.Do(func() {
		listActive = LookupByFilter(func(spec Spec) bool { return !bool(spec.IsImportOnly) })
	})
	return listActive
}

// List returns a slice of all ADIF 3.1.6 QSODownloadStatus specifications.
func List() []Spec {
	list := make([]Spec, len(lookupList))
	copy(list, lookupList)
	return list
}
