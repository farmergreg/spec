// DO NOT EDIT; GENERATED CODE
// run `go generate ./...` from the project root to rebuild this file.

// Package qslmedium provides code and constants as defined in ADIF 3.1.6 (Proposed)
package qslmedium

const (
	CARD QSLMedium = "CARD" // CARD = QSO confirmation via paper QSL card
	EQSL QSLMedium = "EQSL" // EQSL = QSO confirmation via eQSL.cc
	LOTW QSLMedium = "LOTW" // LOTW = QSO confirmation via ARRL Logbook of the World
)

// Lookup look up a specification for the given QSLMedium
func Lookup(qslmedium QSLMedium) (Spec, bool) {
	spec, ok := internalMap[qslmedium]
	return spec, ok
}

// LookupByFilter returns all QSLMedium specifications that match the provided filter function.
func LookupByFilter(filter func(Spec) bool) []Spec {
	result := make([]Spec, 0)
	for _, v := range List() {
		if filter(v) {
			result = append(result, v)
		}
	}
	return result
}

// Generate a list of QSLMedium specifications EXCLUDING those marked import only.
func ListActive() []Spec {
	return []Spec{
		internalMap[CARD],
		internalMap[EQSL],
		internalMap[LOTW],
	}
}

// Generate a list of all QSLMedium specifications INCLUDING those marked import only.
func List() []Spec {
	return []Spec{
		internalMap[CARD],
		internalMap[EQSL],
		internalMap[LOTW],
	}
}

// internalMap is a map of all known QSLMedium specifications
var internalMap = map[QSLMedium]Spec{
	CARD: {IsImportOnly: false, Key: "CARD", Description: "QSO confirmation via paper QSL card"},
	EQSL: {IsImportOnly: false, Key: "EQSL", Description: "QSO confirmation via eQSL.cc"},
	LOTW: {IsImportOnly: false, Key: "LOTW", Description: "QSO confirmation via ARRL Logbook of the World"},
}
