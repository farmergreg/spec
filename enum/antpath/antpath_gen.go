// DO NOT EDIT; GENERATED CODE
// run `go generate ./...` from the project root to rebuild this file.

// Package antpath provides code and constants as defined in ADIF 3.1.6 (Proposed)
package antpath

import "maps"

const (
	G AntPath = "G" // G = grayline
	L AntPath = "L" // L = long path
	O AntPath = "O" // O = other
	S AntPath = "S" // S = short path
)

// All AntPath specifications in ADIF 3.1.6 (Proposed) including depreciated and import only.
func AntPathListAll() []Spec {
	return append([]Spec(nil), internalAntPathListAll...)
}

// All AntPath specifications values in ADIF 3.1.6 (Proposed) that are NOT marked import-only.
func AntPathListCurrent() []Spec {
	return append([]Spec(nil), internalAntPathListCurrent...)
}

// A map of all AntPath from ADIF 3.1.6 (Proposed).
func AntPathMap() map[AntPath]Spec {
	cp := make(map[AntPath]Spec, len(internalAntPathMap))
	maps.Copy(cp, internalAntPathMap)
	return cp
}

// A map of all AntPath specifications.
var internalAntPathMap = map[AntPath]Spec{
	G: {IsImportOnly: false, Key: "G", Description: "grayline"},
	L: {IsImportOnly: false, Key: "L", Description: "long path"},
	O: {IsImportOnly: false, Key: "O", Description: "other"},
	S: {IsImportOnly: false, Key: "S", Description: "short path"},
}

var internalAntPathListAll = []Spec{
	internalAntPathMap[G],
	internalAntPathMap[L],
	internalAntPathMap[O],
	internalAntPathMap[S],
}

var internalAntPathListCurrent = []Spec{
	internalAntPathMap[G],
	internalAntPathMap[L],
	internalAntPathMap[O],
	internalAntPathMap[S],
}
