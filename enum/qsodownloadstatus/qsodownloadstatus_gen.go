// DO NOT EDIT; GENERATED CODE
// run `go generate ./...` from the project root to rebuild this file.

// Package qsodownloadstatus provides code and constants as defined in ADIF 3.1.6 (Proposed)
package qsodownloadstatus

import "maps"

const (
	I QSODownloadStatus = "I" // I = ignore or invalid
	N QSODownloadStatus = "N" // N = the QSO has not been downloaded from the online service
	Y QSODownloadStatus = "Y" // Y = the QSO has been downloaded from the online service
)

// All QSODownloadStatus specifications including depreciated and import only.
func QSODownloadStatusListAll() []Spec {
	return append([]Spec(nil), internalQSODownloadStatusListAll...)
}

// All QSODownloadStatus specifications values that are NOT marked import-only.
func QSODownloadStatusListCurrent() []Spec {
	return append([]Spec(nil), internalQSODownloadStatusListCurrent...)
}

// A map of all QSODownloadStatus specifications.
func QSODownloadStatusMap() map[QSODownloadStatus]Spec {
	cp := make(map[QSODownloadStatus]Spec, len(internalQSODownloadStatusMap))
	maps.Copy(cp, internalQSODownloadStatusMap)
	return cp
}

// A map of all QSODownloadStatus specifications.
var internalQSODownloadStatusMap = map[QSODownloadStatus]Spec{
	I: {IsImportOnly: false, Key: "I", Description: "ignore or invalid"},
	N: {IsImportOnly: false, Key: "N", Description: "the QSO has not been downloaded from the online service"},
	Y: {IsImportOnly: false, Key: "Y", Description: "the QSO has been downloaded from the online service"},
}

var internalQSODownloadStatusListAll = []Spec{
	internalQSODownloadStatusMap[I],
	internalQSODownloadStatusMap[N],
	internalQSODownloadStatusMap[Y],
}

var internalQSODownloadStatusListCurrent = []Spec{
	internalQSODownloadStatusMap[I],
	internalQSODownloadStatusMap[N],
	internalQSODownloadStatusMap[Y],
}
