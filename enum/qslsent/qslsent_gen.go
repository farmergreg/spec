// DO NOT EDIT; GENERATED CODE
// run `go generate ./...` from the project root to rebuild this file.

// Package qslsent provides code and constants as defined in ADIF 3.1.6 (Proposed)
package qslsent

import "maps"

const (
	I QSLSent = "I" // I =
	N QSLSent = "N" // N = do not send an outgoing QSL card do not upload the QSO to the online service
	Q QSLSent = "Q" // Q = an outgoing QSL card has been selected to be sent a QSO has been selected to be uploaded to the online service
	R QSLSent = "R" // R = the contacted station has requested a QSL card the contacted station has requested the QSO be uploaded to the online service
	Y QSLSent = "Y" // Y = an outgoing QSL card has been sent the QSO has been uploaded to, and accepted by, the online service
)

// All QSLSent specifications in ADIF 3.1.6 (Proposed) including depreciated and import only.
func QSLSentListAll() []Spec {
	return append([]Spec(nil), internalQSLSentListAll...)
}

// All QSLSent specifications values in ADIF 3.1.6 (Proposed) that are NOT marked import-only.
func QSLSentListCurrent() []Spec {
	return append([]Spec(nil), internalQSLSentListCurrent...)
}

// A map of all QSLSent from ADIF 3.1.6 (Proposed).
func QSLSentMap() map[QSLSent]Spec {
	cp := make(map[QSLSent]Spec, len(internalQSLSentMap))
	maps.Copy(cp, internalQSLSentMap)
	return cp
}

// A map of all QSLSent specifications.
var internalQSLSentMap = map[QSLSent]Spec{
	I: {IsImportOnly: false, Key: "I", Meaning: "ignore or invalid", Description: ""},
	N: {IsImportOnly: false, Key: "N", Meaning: "no", Description: "do not send an outgoing QSL card do not upload the QSO to the online service"},
	Q: {IsImportOnly: false, Key: "Q", Meaning: "queued", Description: "an outgoing QSL card has been selected to be sent a QSO has been selected to be uploaded to the online service"},
	R: {IsImportOnly: false, Key: "R", Meaning: "requested", Description: "the contacted station has requested a QSL card the contacted station has requested the QSO be uploaded to the online service"},
	Y: {IsImportOnly: false, Key: "Y", Meaning: "yes", Description: "an outgoing QSL card has been sent the QSO has been uploaded to, and accepted by, the online service"},
}

var internalQSLSentListAll = []Spec{
	internalQSLSentMap[I],
	internalQSLSentMap[N],
	internalQSLSentMap[Q],
	internalQSLSentMap[R],
	internalQSLSentMap[Y],
}

var internalQSLSentListCurrent = []Spec{
	internalQSLSentMap[I],
	internalQSLSentMap[N],
	internalQSLSentMap[Q],
	internalQSLSentMap[R],
	internalQSLSentMap[Y],
}
