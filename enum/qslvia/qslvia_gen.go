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
	spec, ok := internalMap[qslvia]
	return spec, ok
}

// LookupByFilter returns all QSLVia specifications that match the provided filter function.
func LookupByFilter(filter func(Spec) bool) []Spec {
	result := make([]Spec, 0)
	for _, v := range List() {
		if filter(v) {
			result = append(result, v)
		}
	}
	return result
}

// All QSLVia specifications INCLUDING those marked import only.
func List() []Spec {
	return []Spec{
		internalMap[B],
		internalMap[D],
		internalMap[E],
		internalMap[M],
	}
}

// QSLVia specifications EXCLUDING those marked import only.
func ListActive() []Spec {
	return []Spec{
		internalMap[B],
		internalMap[D],
		internalMap[E],
	}
}

var internalMap = map[QSLVia]Spec{
	B: {IsImportOnly: false, Key: "B", Description: "bureau"},
	D: {IsImportOnly: false, Key: "D", Description: "direct"},
	E: {IsImportOnly: false, Key: "E", Description: "electronic"},
	M: {IsImportOnly: true, Key: "M", Description: "manager"},
}
