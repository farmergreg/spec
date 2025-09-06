// DO NOT EDIT; GENERATED CODE
// run `go generate ./...` from the project root to rebuild this file.

// Package qslmedium provides code and constants as defined in ADIF 3.1.6 (Proposed)
package qslmedium

const (
	CARD QSLMedium = "CARD" // CARD = QSO confirmation via paper QSL card
	EQSL QSLMedium = "EQSL" // EQSL = QSO confirmation via eQSL.cc
	LOTW QSLMedium = "LOTW" // LOTW = QSO confirmation via ARRL Logbook of the World
)

// Lookup look up a specification for QSLMedium
func Lookup(qslmedium QSLMedium) (Spec, bool) {
	spec, ok := internalMap[qslmedium]
	return spec, ok
}

// IsValid returns true if the specification for QSLMedium exists and is not import only.
func IsValid(qslmedium QSLMedium) bool {
	spec, ok := internalMap[qslmedium]
	if ok && bool(spec.IsImportOnly) {
		return false
	}
	return ok
}

// All QSLMedium specifications INCLUDING ones marked import only.
func AllQSLMedium() []Spec {
	result := make([]Spec, 0, len(internalMap))
	for _, v := range internalMap {
		result = append(result, v)
	}
	return result
}

// AllActiveQSLMedium specifications EXCLUDING ones marked import only.
func AllActiveQSLMedium() []Spec {
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

var internalMap = map[QSLMedium]Spec{
	CARD: {IsImportOnly: false, Key: "CARD", Description: "QSO confirmation via paper QSL card"},
	EQSL: {IsImportOnly: false, Key: "EQSL", Description: "QSO confirmation via eQSL.cc"},
	LOTW: {IsImportOnly: false, Key: "LOTW", Description: "QSO confirmation via ARRL Logbook of the World"},
}
