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

// Lookup look up a specification for the given AntPath
func Lookup(antpath AntPath) (Spec, bool) {
	spec, ok := internalMap[antpath]
	return spec, ok
}

// LookupByFilter returns all AntPath specifications that match the provided filter function.
func LookupByFilter(filter func(Spec) bool) []Spec {
	result := make([]Spec, 0)
	for _, v := range List() {
		if filter(v) {
			result = append(result, v)
		}
	}
	return result
}

// Generate a list of AntPath specifications EXCLUDING those marked import only.
func ListActive() []Spec {
	return []Spec{
		internalMap[G],
		internalMap[L],
		internalMap[O],
		internalMap[S],
	}
}

// Generate a list of all AntPath specifications INCLUDING those marked import only.
func List() []Spec {
	return []Spec{
		internalMap[G],
		internalMap[L],
		internalMap[O],
		internalMap[S],
	}
}

// internalMap is a map of all known AntPath specifications
var internalMap = map[AntPath]Spec{
	G: {IsImportOnly: false, Key: "G", Description: "grayline"},
	L: {IsImportOnly: false, Key: "L", Description: "long path"},
	O: {IsImportOnly: false, Key: "O", Description: "other"},
	S: {IsImportOnly: false, Key: "S", Description: "short path"},
}
