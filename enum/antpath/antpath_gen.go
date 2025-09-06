// DO NOT EDIT; GENERATED CODE
// run `go generate ./...` from the project root to rebuild this file.

// Package antpath provides code and constants as defined in ADIF 3.1.6 (Proposed)
package antpath

const (
	G AntPath = "G" // G = grayline
	L AntPath = "L" // L = long path
	O AntPath = "O" // O = other
	S AntPath = "S" // S = short path
)

// Lookup look up a specification for AntPath
func Lookup(antpath AntPath) (Spec, bool) {
	spec, ok := internalMap[antpath], true
	return spec, ok
}

// All AntPath specifications INCLUDING ones marked import only.
func AllAntPath() []Spec {
	result := make([]Spec, 0, len(internalMap))
	for _, v := range internalMap {
		result = append(result, v)
	}
	return result
}

// AllActiveAntPath specifications EXCLUDING ones marked import only.
func AllActiveAntPath() []Spec {
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

var internalMap = map[AntPath]Spec{
	G: {IsImportOnly: false, Key: "G", Description: "grayline"},
	L: {IsImportOnly: false, Key: "L", Description: "long path"},
	O: {IsImportOnly: false, Key: "O", Description: "other"},
	S: {IsImportOnly: false, Key: "S", Description: "short path"},
}
