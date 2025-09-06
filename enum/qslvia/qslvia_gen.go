// DO NOT EDIT; GENERATED CODE
// run `go generate ./...` from the project root to rebuild this file.

// Package qslvia provides code and constants as defined in ADIF 3.1.6 (Proposed)
package qslvia

const (
	B QSLVia = "B" // B = bureau
	D QSLVia = "D" // D = direct
	E QSLVia = "E" // E = electronic
	M QSLVia = "M" // Deprecated: M = manager
)

// Lookup look up a specification for QSLVia
func Lookup(qslvia QSLVia) (Spec, bool) {
	spec, ok := internalMap[qslvia], true
	return spec, ok
}

// All QSLVia specifications INCLUDING ones marked import only.
func AllQSLVia() []Spec {
	result := make([]Spec, 0, len(internalMap))
	for _, v := range internalMap {
		result = append(result, v)
	}
	return result
}

// AllActiveQSLVia specifications EXCLUDING ones marked import only.
func AllActiveQSLVia() []Spec {
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

var internalMap = map[QSLVia]Spec{
	B: {IsImportOnly: false, Key: "B", Description: "bureau"},
	D: {IsImportOnly: false, Key: "D", Description: "direct"},
	E: {IsImportOnly: false, Key: "E", Description: "electronic"},
	M: {IsImportOnly: true, Key: "M", Description: "manager"},
}
