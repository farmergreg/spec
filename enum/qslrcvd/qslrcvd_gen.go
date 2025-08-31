// DO NOT EDIT; GENERATED CODE
// run `go generate ./...` from the project root to rebuild this file.

// Package qslrcvd provides code and constants as defined in ADIF 3.1.6 (Proposed)
package qslrcvd

import "maps"

const (
	I QSLRcvd = "I" // I =
	N QSLRcvd = "N" // N = an incoming QSL card has not been received the QSO has not been confirmed by the online service
	R QSLRcvd = "R" // R = the logging station has requested a QSL card the logging station has requested the QSO be uploaded to the online service
	V QSLRcvd = "V" // Deprecated: V = DXCC award credit granted for the QSL card - instead use <CREDIT_GRANTED:39>DXCC:card,DXCC_BAND:card,DXCC_MODE:card DXCC credit granted for the LoTW confirmation - instead use <CREDIT_GRANTED:39>DXCC:lotw,DXCC_BAND:lotw,DXCC_MODE:lotw
	Y QSLRcvd = "Y" // Y = an incoming QSL card has been received the QSO has been confirmed by the online service
)

// All QSLRcvd specifications including depreciated and import only.
func QSLRcvdListAll() []Spec {
	return append([]Spec(nil), internalQSLRcvdListAll...)
}

// All QSLRcvd specifications values that are NOT marked import-only.
func QSLRcvdListCurrent() []Spec {
	return append([]Spec(nil), internalQSLRcvdListCurrent...)
}

// A map of all QSLRcvd specifications.
func QSLRcvdMap() map[QSLRcvd]Spec {
	cp := make(map[QSLRcvd]Spec, len(internalQSLRcvdMap))
	maps.Copy(cp, internalQSLRcvdMap)
	return cp
}

// A map of all QSLRcvd specifications.
var internalQSLRcvdMap = map[QSLRcvd]Spec{
	I: {IsImportOnly: false, Key: "I", Meaning: "ignore or invalid", Description: ""},
	N: {IsImportOnly: false, Key: "N", Meaning: "no", Description: "an incoming QSL card has not been received the QSO has not been confirmed by the online service"},
	R: {IsImportOnly: false, Key: "R", Meaning: "requested", Description: "the logging station has requested a QSL card the logging station has requested the QSO be uploaded to the online service"},
	V: {IsImportOnly: true, Key: "V", Meaning: "verified", Description: "DXCC award credit granted for the QSL card - instead use <CREDIT_GRANTED:39>DXCC:card,DXCC_BAND:card,DXCC_MODE:card DXCC credit granted for the LoTW confirmation - instead use <CREDIT_GRANTED:39>DXCC:lotw,DXCC_BAND:lotw,DXCC_MODE:lotw"},
	Y: {IsImportOnly: false, Key: "Y", Meaning: "yes (confirmed)", Description: "an incoming QSL card has been received the QSO has been confirmed by the online service"},
}

var internalQSLRcvdListAll = []Spec{
	internalQSLRcvdMap[I],
	internalQSLRcvdMap[N],
	internalQSLRcvdMap[R],
	internalQSLRcvdMap[V],
	internalQSLRcvdMap[Y],
}

var internalQSLRcvdListCurrent = []Spec{
	internalQSLRcvdMap[I],
	internalQSLRcvdMap[N],
	internalQSLRcvdMap[R],
	internalQSLRcvdMap[Y],
}
