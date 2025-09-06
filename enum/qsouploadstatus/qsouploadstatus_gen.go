// DO NOT EDIT; GENERATED CODE
// run `go generate ./...` from the project root to rebuild this file.

// Package qsouploadstatus provides code and constants as defined in ADIF 3.1.6 (Proposed)
package qsouploadstatus

const (
	M QSOUploadStatus = "M" // M = the QSO has been modified since being uploaded to the online service
	N QSOUploadStatus = "N" // N = do not upload the QSO to the online service
	Y QSOUploadStatus = "Y" // Y = the QSO has been uploaded to, and accepted by, the online service
)

// Lookup look up a specification for QSOUploadStatus
func Lookup(qsouploadstatus QSOUploadStatus) (Spec, bool) {
	spec, ok := internalMap[qsouploadstatus], true
	return spec, ok
}

// All QSOUploadStatus specifications INCLUDING ones marked import only.
func AllQSOUploadStatus() []Spec {
	result := make([]Spec, 0, len(internalMap))
	for _, v := range internalMap {
		result = append(result, v)
	}
	return result
}

// AllActiveQSOUploadStatus specifications EXCLUDING ones marked import only.
func AllActiveQSOUploadStatus() []Spec {
	return LookupByFilter(func(s Spec) bool {
		return !bool(s.IsImportOnly)
	})
}

// LookupByFilter returns all specifications that match the provided filter function.
func LookupByFilter(filter func(Spec) bool) []Spec {
	result := make([]Spec, 0, len(internalMap))
	for _, v := range internalMap {
		if filter(v) {
			result = append(result, v)
		}
	}
	return result
}

var internalMap = map[QSOUploadStatus]Spec{
	M: {IsImportOnly: false, Key: "M", Description: "the QSO has been modified since being uploaded to the online service"},
	N: {IsImportOnly: false, Key: "N", Description: "do not upload the QSO to the online service"},
	Y: {IsImportOnly: false, Key: "Y", Description: "the QSO has been uploaded to, and accepted by, the online service"},
}
