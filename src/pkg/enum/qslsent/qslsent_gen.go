// DO NOT EDIT; GENERATED CODE
// ADIF: 3.1.6 Proposed

package qslsent

var (
	I QSLSent = "I" // I =
	N QSLSent = "N" // N = do not send an outgoing QSL card do not upload the QSO to the online service
	Q QSLSent = "Q" // Q = an outgoing QSL card has been selected to be sent a QSO has been selected to be uploaded to the online service
	R QSLSent = "R" // R = the contacted station has requested a QSL card the contacted station has requested the QSO be uploaded to the online service
	Y QSLSent = "Y" // Y = an outgoing QSL card has been sent the QSO has been uploaded to, and accepted by, the online service
)

// A map of all QSLSent specifications.
var QSLSentMap = map[QSLSent]Spec{
	I: {IsImportOnly: false, Key: "I", Meaning: "ignore or invalid", Description: ""},
	N: {IsImportOnly: false, Key: "N", Meaning: "no", Description: "do not send an outgoing QSL card do not upload the QSO to the online service"},
	Q: {IsImportOnly: false, Key: "Q", Meaning: "queued", Description: "an outgoing QSL card has been selected to be sent a QSO has been selected to be uploaded to the online service"},
	R: {IsImportOnly: false, Key: "R", Meaning: "requested", Description: "the contacted station has requested a QSL card the contacted station has requested the QSO be uploaded to the online service"},
	Y: {IsImportOnly: false, Key: "Y", Meaning: "yes", Description: "an outgoing QSL card has been sent the QSO has been uploaded to, and accepted by, the online service"},
}

// All QSLSent specifications including depreciated and import only.
var QSLSentListAll = []Spec{
	QSLSentMap[I],
	QSLSentMap[N],
	QSLSentMap[Q],
	QSLSentMap[R],
	QSLSentMap[Y],
}

// All QSLSent specifications values that are NOT marked import-only.
var QSLSentListCurrent = []Spec{
	QSLSentMap[I],
	QSLSentMap[N],
	QSLSentMap[Q],
	QSLSentMap[R],
	QSLSentMap[Y],
}
