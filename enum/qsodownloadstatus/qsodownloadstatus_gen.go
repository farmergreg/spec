// DO NOT EDIT; GENERATED CODE
// ADIF: 3.1.6 Proposed

package qsodownloadstatus

const (
	I QSODownloadStatus = "I" // I = ignore or invalid
	N QSODownloadStatus = "N" // N = the QSO has not been downloaded from the online service
	Y QSODownloadStatus = "Y" // Y = the QSO has been downloaded from the online service
)

// A map of all QSODownloadStatus specifications.
var QSODownloadStatusMap = map[QSODownloadStatus]Spec{
	I: {IsImportOnly: false, Key: "I", Description: "ignore or invalid"},
	N: {IsImportOnly: false, Key: "N", Description: "the QSO has not been downloaded from the online service"},
	Y: {IsImportOnly: false, Key: "Y", Description: "the QSO has been downloaded from the online service"},
}

// All QSODownloadStatus specifications including depreciated and import only.
var QSODownloadStatusListAll = []Spec{
	QSODownloadStatusMap[I],
	QSODownloadStatusMap[N],
	QSODownloadStatusMap[Y],
}

// All QSODownloadStatus specifications values that are NOT marked import-only.
var QSODownloadStatusListCurrent = []Spec{
	QSODownloadStatusMap[I],
	QSODownloadStatusMap[N],
	QSODownloadStatusMap[Y],
}
