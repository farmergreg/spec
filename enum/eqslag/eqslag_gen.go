// DO NOT EDIT; GENERATED CODE
// run `go generate ./...` from the project root to rebuild this file.

// Package eqslag provides code and constants as defined in ADIF 3.1.6 (Proposed)
package eqslag

const (
	N EQSLAG = "N" // N = the QSO is confirmed but not "Authenticity Guaranteed" by eQSL
	U EQSLAG = "U" // U = unknown
	Y EQSLAG = "Y" // Y = the QSO is confirmed and "Authenticity Guaranteed" by eQSL
)

// Lookup look up a specification for EQSLAG
func Lookup(eqslag EQSLAG) (Spec, bool) {
	spec, ok := internalMap[eqslag]
	return spec, ok
}

// LookupByFilter returns all EQSLAG specifications that match the provided filter function.
func LookupByFilter(filter func(Spec) bool) []Spec {
	result := make([]Spec, 0)
	for _, v := range List() {
		if filter(v) {
			result = append(result, v)
		}
	}
	return result
}

// All EQSLAG specifications INCLUDING those marked import only.
func List() []Spec {
	return []Spec{
		internalMap[N],
		internalMap[U],
		internalMap[Y],
	}
}

// EQSLAG specifications EXCLUDING those marked import only.
func ListActive() []Spec {
	return []Spec{
		internalMap[N],
		internalMap[U],
		internalMap[Y],
	}
}

var internalMap = map[EQSLAG]Spec{
	N: {IsImportOnly: false, Key: "N", Description: "the QSO is confirmed but not \"Authenticity Guaranteed\" by eQSL"},
	U: {IsImportOnly: false, Key: "U", Description: "unknown"},
	Y: {IsImportOnly: false, Key: "Y", Description: "the QSO is confirmed and \"Authenticity Guaranteed\" by eQSL"},
}
