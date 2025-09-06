// DO NOT EDIT; GENERATED CODE
// run `go generate ./...` from the project root to rebuild this file.

// Package qsocomplete provides code and constants as defined in ADIF 3.1.6 (Proposed)
package qsocomplete

const (
	Uncertain QSOComplete = "?"   // ?    = uncertain
	N         QSOComplete = "N"   // N    = no
	NIL       QSOComplete = "NIL" // NIL  = not heard
	Y         QSOComplete = "Y"   // Y    = yes
)

// Lookup look up a specification for the given QSOComplete
func Lookup(qsocomplete QSOComplete) (Spec, bool) {
	spec, ok := internalMap[qsocomplete]
	return spec, ok
}

// LookupByFilter returns all QSOComplete specifications that match the provided filter function.
func LookupByFilter(filter func(Spec) bool) []Spec {
	result := make([]Spec, 0)
	for _, v := range List() {
		if filter(v) {
			result = append(result, v)
		}
	}
	return result
}

// Generate a list of QSOComplete specifications EXCLUDING those marked import only.
func ListActive() []Spec {
	return []Spec{
		internalMap[Uncertain],
		internalMap[N],
		internalMap[NIL],
		internalMap[Y],
	}
}

// Generate a list of all QSOComplete specifications INCLUDING those marked import only.
func List() []Spec {
	return []Spec{
		internalMap[Uncertain],
		internalMap[N],
		internalMap[NIL],
		internalMap[Y],
	}
}

// internalMap is a map of all known QSOComplete specifications
var internalMap = map[QSOComplete]Spec{
	Uncertain: {IsImportOnly: false, Key: "?", Description: "uncertain"},
	N:         {IsImportOnly: false, Key: "N", Description: "no"},
	NIL:       {IsImportOnly: false, Key: "NIL", Description: "not heard"},
	Y:         {IsImportOnly: false, Key: "Y", Description: "yes"},
}
