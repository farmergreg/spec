// DO NOT EDIT; GENERATED CODE
// run `go generate ./...` from the project root to rebuild this file.

// Package qsodownloadstatus provides code and constants as defined in ADIF 3.1.6 (Proposed)
package qsodownloadstatus

const (
	I QSODownloadStatus = "I" // I = ignore or invalid
	N QSODownloadStatus = "N" // N = the QSO has not been downloaded from the online service
	Y QSODownloadStatus = "Y" // Y = the QSO has been downloaded from the online service
)

// Lookup look up a specification for QSODownloadStatus
func Lookup(qsodownloadstatus QSODownloadStatus) (Spec, bool) {
	spec, ok := internalMap[qsodownloadstatus]
	return spec, ok
}

// LookupByFilter returns all QSODownloadStatus specifications that match the provided filter function.
func LookupByFilter(filter func(Spec) bool) []Spec {
	result := make([]Spec, 0)
	for _, v := range List() {
		if filter(v) {
			result = append(result, v)
		}
	}
	return result
}

// All QSODownloadStatus specifications INCLUDING those marked import only.
func List() []Spec {
	return []Spec{
		internalMap[I],
		internalMap[N],
		internalMap[Y],
	}
}

// QSODownloadStatus specifications EXCLUDING those marked import only.
func ListActive() []Spec {
	return []Spec{
		internalMap[I],
		internalMap[N],
		internalMap[Y],
	}
}

var internalMap = map[QSODownloadStatus]Spec{
	I: {IsImportOnly: false, Key: "I", Description: "ignore or invalid"},
	N: {IsImportOnly: false, Key: "N", Description: "the QSO has not been downloaded from the online service"},
	Y: {IsImportOnly: false, Key: "Y", Description: "the QSO has been downloaded from the online service"},
}
