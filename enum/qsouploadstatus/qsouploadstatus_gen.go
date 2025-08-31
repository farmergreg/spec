// DO NOT EDIT; GENERATED CODE
// run `go generate ./...` from the project root to rebuild this file.

// Package qsouploadstatus provides code and constants as defined in ADIF 3.1.6 (Proposed)
package qsouploadstatus

import "maps"

const (
	M QSOUploadStatus = "M" // M = the QSO has been modified since being uploaded to the online service
	N QSOUploadStatus = "N" // N = do not upload the QSO to the online service
	Y QSOUploadStatus = "Y" // Y = the QSO has been uploaded to, and accepted by, the online service
)

// All QSOUploadStatus specifications in ADIF 3.1.6 (Proposed) including depreciated and import only.
func QSOUploadStatusListAll() []Spec {
	return append([]Spec(nil), internalQSOUploadStatusListAll...)
}

// All QSOUploadStatus specifications values in ADIF 3.1.6 (Proposed) that are NOT marked import-only.
func QSOUploadStatusListCurrent() []Spec {
	return append([]Spec(nil), internalQSOUploadStatusListCurrent...)
}

// A map of all QSOUploadStatus from ADIF 3.1.6 (Proposed).
func QSOUploadStatusMap() map[QSOUploadStatus]Spec {
	cp := make(map[QSOUploadStatus]Spec, len(internalQSOUploadStatusMap))
	maps.Copy(cp, internalQSOUploadStatusMap)
	return cp
}

// A map of all QSOUploadStatus specifications.
var internalQSOUploadStatusMap = map[QSOUploadStatus]Spec{
	M: {IsImportOnly: false, Key: "M", Description: "the QSO has been modified since being uploaded to the online service"},
	N: {IsImportOnly: false, Key: "N", Description: "do not upload the QSO to the online service"},
	Y: {IsImportOnly: false, Key: "Y", Description: "the QSO has been uploaded to, and accepted by, the online service"},
}

var internalQSOUploadStatusListAll = []Spec{
	internalQSOUploadStatusMap[M],
	internalQSOUploadStatusMap[N],
	internalQSOUploadStatusMap[Y],
}

var internalQSOUploadStatusListCurrent = []Spec{
	internalQSOUploadStatusMap[M],
	internalQSOUploadStatusMap[N],
	internalQSOUploadStatusMap[Y],
}
