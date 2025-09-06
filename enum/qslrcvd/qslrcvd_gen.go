// DO NOT EDIT; GENERATED CODE
// run `go generate ./...` from the project root to rebuild this file.

// Package qslrcvd provides code and constants as defined in ADIF 3.1.6 (Proposed)
package qslrcvd

const (
	I QSLRcvd = "I" // I =
	N QSLRcvd = "N" // N = an incoming QSL card has not been received the QSO has not been confirmed by the online service
	R QSLRcvd = "R" // R = the logging station has requested a QSL card the logging station has requested the QSO be uploaded to the online service
	V QSLRcvd = "V" // Deprecated: V = DXCC award credit granted for the QSL card - instead use <CREDIT_GRANTED:39>DXCC:card,DXCC_BAND:card,DXCC_MODE:card DXCC credit granted for the LoTW confirmation - instead use <CREDIT_GRANTED:39>DXCC:lotw,DXCC_BAND:lotw,DXCC_MODE:lotw
	Y QSLRcvd = "Y" // Y = an incoming QSL card has been received the QSO has been confirmed by the online service
)

// Lookup look up a specification for QSLRcvd
func Lookup(qslrcvd QSLRcvd) (Spec, bool) {
	spec, ok := internalMap[qslrcvd], true
	return spec, ok
}

// All QSLRcvd specifications INCLUDING ones marked import only.
func AllQSLRcvd() []Spec {
	result := make([]Spec, 0, len(internalMap))
	for _, v := range internalMap {
		result = append(result, v)
	}
	return result
}

// AllActiveQSLRcvd specifications EXCLUDING ones marked import only.
func AllActiveQSLRcvd() []Spec {
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

var internalMap = map[QSLRcvd]Spec{
	I: {IsImportOnly: false, Key: "I", Meaning: "ignore or invalid", Description: ""},
	N: {IsImportOnly: false, Key: "N", Meaning: "no", Description: "an incoming QSL card has not been received the QSO has not been confirmed by the online service"},
	R: {IsImportOnly: false, Key: "R", Meaning: "requested", Description: "the logging station has requested a QSL card the logging station has requested the QSO be uploaded to the online service"},
	V: {IsImportOnly: true, Key: "V", Meaning: "verified", Description: "DXCC award credit granted for the QSL card - instead use <CREDIT_GRANTED:39>DXCC:card,DXCC_BAND:card,DXCC_MODE:card DXCC credit granted for the LoTW confirmation - instead use <CREDIT_GRANTED:39>DXCC:lotw,DXCC_BAND:lotw,DXCC_MODE:lotw"},
	Y: {IsImportOnly: false, Key: "Y", Meaning: "yes (confirmed)", Description: "an incoming QSL card has been received the QSO has been confirmed by the online service"},
}
