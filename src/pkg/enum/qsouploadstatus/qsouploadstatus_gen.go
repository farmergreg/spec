// DO NOT EDIT; GENERATED CODE
// ADIF: 3.1.6 Proposed

package qsouploadstatus

var (
	M QSOUploadStatus = "M" // M = the QSO has been modified since being uploaded to the online service
	N QSOUploadStatus = "N" // N = do not upload the QSO to the online service
	Y QSOUploadStatus = "Y" // Y = the QSO has been uploaded to, and accepted by, the online service
)

// A map of all QSOUploadStatus specifications.
var QSOUploadStatusMap = map[QSOUploadStatus]Spec{
	M: {IsImportOnly: false, Key: "M", Description: "the QSO has been modified since being uploaded to the online service"},
	N: {IsImportOnly: false, Key: "N", Description: "do not upload the QSO to the online service"},
	Y: {IsImportOnly: false, Key: "Y", Description: "the QSO has been uploaded to, and accepted by, the online service"},
}

// All QSOUploadStatus specifications including depreciated and import only.
var QSOUploadStatusListAll = []Spec{
	QSOUploadStatusMap[M],
	QSOUploadStatusMap[N],
	QSOUploadStatusMap[Y],
}

// All QSOUploadStatus specifications values that are NOT marked import-only.
var QSOUploadStatusListCurrent = []Spec{
	QSOUploadStatusMap[M],
	QSOUploadStatusMap[N],
	QSOUploadStatusMap[Y],
}
