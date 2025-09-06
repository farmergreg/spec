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

// Lookup look up a specification for QSOComplete
func Lookup(qsocomplete QSOComplete) (Spec, bool) {
	spec, ok := internalMap[qsocomplete]
	return spec, ok
}

// IsValid returns true if the specification for QSOComplete exists and is not import only.
func IsValid(qsocomplete QSOComplete) bool {
	spec, ok := internalMap[qsocomplete]
	if ok && bool(spec.IsImportOnly) {
		return false
	}
	return ok
}

// All QSOComplete specifications INCLUDING ones marked import only.
func AllQSOComplete() []Spec {
	result := make([]Spec, 0, len(internalMap))
	for _, v := range internalMap {
		result = append(result, v)
	}
	return result
}

// AllActiveQSOComplete specifications EXCLUDING ones marked import only.
func AllActiveQSOComplete() []Spec {
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

var internalMap = map[QSOComplete]Spec{
	Uncertain: {IsImportOnly: false, Key: "?", Description: "uncertain"},
	N:         {IsImportOnly: false, Key: "N", Description: "no"},
	NIL:       {IsImportOnly: false, Key: "NIL", Description: "not heard"},
	Y:         {IsImportOnly: false, Key: "Y", Description: "yes"},
}
