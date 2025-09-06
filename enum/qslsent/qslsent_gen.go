// DO NOT EDIT; GENERATED CODE
// run `go generate ./...` from the project root to rebuild this file.

// Package qslsent provides code and constants as defined in ADIF 3.1.6 (Proposed)
package qslsent

const (
	I QSLSent = "I" // I =
	N QSLSent = "N" // N = do not send an outgoing QSL card do not upload the QSO to the online service
	Q QSLSent = "Q" // Q = an outgoing QSL card has been selected to be sent a QSO has been selected to be uploaded to the online service
	R QSLSent = "R" // R = the contacted station has requested a QSL card the contacted station has requested the QSO be uploaded to the online service
	Y QSLSent = "Y" // Y = an outgoing QSL card has been sent the QSO has been uploaded to, and accepted by, the online service
)

// Lookup look up a specification for QSLSent
func Lookup(qslsent QSLSent) (Spec, bool) {
	spec, ok := internalMap[qslsent], true
	return spec, ok
}

// All QSLSent specifications INCLUDING ones marked import only.
func AllQSLSent() []Spec {
	result := make([]Spec, 0, len(internalMap))
	for _, v := range internalMap {
		result = append(result, v)
	}
	return result
}

// AllActiveQSLSent specifications EXCLUDING ones marked import only.
func AllActiveQSLSent() []Spec {
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

var internalMap = map[QSLSent]Spec{
	I: {IsImportOnly: false, Key: "I", Meaning: "ignore or invalid", Description: ""},
	N: {IsImportOnly: false, Key: "N", Meaning: "no", Description: "do not send an outgoing QSL card do not upload the QSO to the online service"},
	Q: {IsImportOnly: false, Key: "Q", Meaning: "queued", Description: "an outgoing QSL card has been selected to be sent a QSO has been selected to be uploaded to the online service"},
	R: {IsImportOnly: false, Key: "R", Meaning: "requested", Description: "the contacted station has requested a QSL card the contacted station has requested the QSO be uploaded to the online service"},
	Y: {IsImportOnly: false, Key: "Y", Meaning: "yes", Description: "an outgoing QSL card has been sent the QSO has been uploaded to, and accepted by, the online service"},
}
