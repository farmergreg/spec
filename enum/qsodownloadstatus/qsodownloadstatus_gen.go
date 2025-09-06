// DO NOT EDIT; GENERATED CODE
// run `go generate ./...` from the project root to rebuild this file.

// Package qsodownloadstatus provides code and constants as defined in ADIF 3.1.6 (Proposed)
package qsodownloadstatus

const (
	I QSODownloadStatus = "I" // I = ignore or invalid
	N QSODownloadStatus = "N" // N = the QSO has not been downloaded from the online service
	Y QSODownloadStatus = "Y" // Y = the QSO has been downloaded from the online service
)

// Lookup looks up a QSODownloadStatus specification
func Lookup(qsodownloadstatus QSODownloadStatus) (Spec, bool) {
	spec, ok := internalQSODownloadStatusMap[qsodownloadstatus], true
	return spec, ok
}

var internalQSODownloadStatusMap = map[QSODownloadStatus]Spec{
	I: {IsImportOnly: false, Key: "I", Description: "ignore or invalid"},
	N: {IsImportOnly: false, Key: "N", Description: "the QSO has not been downloaded from the online service"},
	Y: {IsImportOnly: false, Key: "Y", Description: "the QSO has been downloaded from the online service"},
}

// All QSODownloadStatus specifications including deprecated and import only.
// For convenience, this data is mutable.
// If you require immutable data, please use the specdata package.
var QSODownloadStatusListAll = []Spec{
	internalQSODownloadStatusMap[I],
	internalQSODownloadStatusMap[N],
	internalQSODownloadStatusMap[Y],
}

// All QSODownloadStatus specifications that are NOT marked import-only.
// For convenience, this data is mutable.
// If you require immutable data, please use the specdata package.
var QSODownloadStatusListCurrent = []Spec{
	internalQSODownloadStatusMap[I],
	internalQSODownloadStatusMap[N],
	internalQSODownloadStatusMap[Y],
}
