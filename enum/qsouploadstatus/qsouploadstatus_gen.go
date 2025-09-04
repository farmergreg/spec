// DO NOT EDIT; GENERATED CODE
// run `go generate ./...` from the project root to rebuild this file.

// Package qsouploadstatus provides code and constants as defined in ADIF 3.1.6 (Proposed)
package qsouploadstatus

const (
	M QSOUploadStatus = "M" // M = the QSO has been modified since being uploaded to the online service
	N QSOUploadStatus = "N" // N = do not upload the QSO to the online service
	Y QSOUploadStatus = "Y" // Y = the QSO has been uploaded to, and accepted by, the online service
)

// A map of all QSOUploadStatus specifications.
// For convenience, this data is mutable.
// If you require immutable data, please use the specdata package.
var QSOUploadStatusMap = map[QSOUploadStatus]Spec{
	M: {IsImportOnly: false, Key: "M", Description: "the QSO has been modified since being uploaded to the online service"},
	N: {IsImportOnly: false, Key: "N", Description: "do not upload the QSO to the online service"},
	Y: {IsImportOnly: false, Key: "Y", Description: "the QSO has been uploaded to, and accepted by, the online service"},
}

// All QSOUploadStatus specifications including deprecated and import only.
// For convenience, this data is mutable.
// If you require immutable data, please use the specdata package.
var QSOUploadStatusListAll = []Spec{
	QSOUploadStatusMap[M],
	QSOUploadStatusMap[N],
	QSOUploadStatusMap[Y],
}

// All QSOUploadStatus specifications that are NOT marked import-only.
// For convenience, this data is mutable.
// If you require immutable data, please use the specdata package.
var QSOUploadStatusListCurrent = []Spec{
	QSOUploadStatusMap[M],
	QSOUploadStatusMap[N],
	QSOUploadStatusMap[Y],
}
